/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package database provides the database interface.
package database

import (
	"context"
	"fmt"

	"github.com/tickexvn/tickex/pkg/logger"
)

// Repository provides the interface for the database.
type Repository[T any, ID comparable] interface {
	Create(ctx context.Context, entity T) (T, error)
	Get(ctx context.Context, id ID) (T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, id ID, entity T) (T, error)
	Delete(ctx context.Context, id ID) error
	Exists(ctx context.Context, id ID) (bool, error)
	Count(ctx context.Context) (int64, error)
}

// Entity is the interface that wraps the basic GetID method.
type Entity[ID comparable] interface {
	GetID() ID
}

// New create multi-layer database instance
func New[T Entity[ID], ID comparable](
	searchLayer Repository[T, ID], storageLayer Repository[T, ID]) Repository[T, ID] {

	return &Database[T, ID]{
		search:  searchLayer,
		storage: storageLayer,
	}
}

// Database database multilayer with postgre & elasticsearch
type Database[T Entity[ID], ID comparable] struct {
	search  Repository[T, ID]
	storage Repository[T, ID]
}

// Create inserts a new entity into both storage (PostgreSQL) and
// search (Elasticsearch). If Elasticsearch fails after PostgreSQL succeeds,
// it rolls back by deleting the entity from PostgreSQL.
func (db *Database[T, ID]) Create(ctx context.Context, entity T) (T, error) {
	var empty T
	if db.storage == nil {
		return empty, fmt.Errorf("storage layer is nil")
	}

	// Insert into PostgreSQL first
	createdEntity, err := db.storage.Create(ctx, entity)
	if err != nil {
		return empty, err
	}

	if db.search == nil {
		logger.Infof("search layer is nil, pass")
		return createdEntity, nil
	}

	// Insert into Elasticsearch
	_, esErr := db.search.Create(ctx, entity)
	if esErr != nil {
		// Rollback: delete from PostgreSQL if Elasticsearch insertion fails
		_ = db.storage.Delete(ctx, getEntityID(entity))
		return empty,
			fmt.Errorf("failed to insert into first layer, rollback storage: %v", esErr)
	}

	return createdEntity, nil
}

// Update modifies an existing entity in both storage (PostgreSQL) and
// search (Elasticsearch).If Elasticsearch fails after PostgreSQL succeeds,
// it rolls back by restoring the old value in PostgreSQL.
func (db *Database[T, ID]) Update(ctx context.Context, id ID, entity T) (T, error) {
	var empty T
	if db.storage == nil {
		return empty, fmt.Errorf("storage layer is nil")
	}

	// Get current entity from storage for rollback purposes
	oldEntity, err := db.storage.Get(ctx, id)
	if err != nil {
		return empty, fmt.Errorf("failed to retrieve old entity before update: %v", err)
	}

	// Update in PostgreSQL
	updatedEntity, err := db.storage.Update(ctx, id, entity)
	if err != nil {
		return empty, err
	}

	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return updatedEntity, nil
	}

	// Update in Elasticsearch
	_, esErr := db.search.Update(ctx, id, entity)
	if esErr != nil {
		// Rollback: Restore old entity in PostgreSQL if Elasticsearch update fails
		_, _ = db.storage.Update(ctx, id, oldEntity)
		return empty,
			fmt.Errorf("failed to update in first layer, rollback storage: %v", esErr)
	}

	return updatedEntity, nil
}

// Delete removes an entity from both storage (PostgreSQL) and search (Elasticsearch).
// If Elasticsearch fails after PostgreSQL succeeds, it rolls back by restoring the deleted entity.
func (db *Database[T, ID]) Delete(ctx context.Context, id ID) error {
	if db.storage == nil {
		return fmt.Errorf("storage layer is nil")
	}

	// Get the entity before deletion for rollback
	oldEntity, err := db.storage.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to retrieve entity before deletion: %v", err)
	}

	// Delete from PostgreSQL
	if err := db.storage.Delete(ctx, id); err != nil {
		return err
	}

	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return nil
	}

	// Delete from Elasticsearch
	if esErr := db.search.Delete(ctx, id); esErr != nil {
		// Rollback: Restore deleted entity in PostgreSQL if Elasticsearch deletion fails
		_, _ = db.storage.Create(ctx, oldEntity)
		return fmt.Errorf("failed to delete from first layer, rollback storage: %v", esErr)
	}

	return nil
}

// Get retrieves an entity by ID from the search layer (Elasticsearch).
func (db *Database[T, ID]) Get(ctx context.Context, id ID) (T, error) {
	var empty T
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		if db.storage == nil {
			return empty, fmt.Errorf("storage layer is nil")
		}

		return db.storage.Get(ctx, id)
	}

	return db.search.Get(ctx, id)
}

// GetAll retrieves all entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) GetAll(ctx context.Context) ([]T, error) {
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		if db.storage == nil {
			return nil, fmt.Errorf("storage layer is nil")
		}

		return db.storage.GetAll(ctx)
	}

	return db.search.GetAll(ctx)
}

// Exists checks if an entity exists in the search layer (Elasticsearch).
func (db *Database[T, ID]) Exists(ctx context.Context, id ID) (bool, error) {
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		if db.storage == nil {
			return false, fmt.Errorf("storage layer is nil")
		}

		return db.storage.Exists(ctx, id)
	}

	return db.search.Exists(ctx, id)
}

// Count returns the total number of entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) Count(ctx context.Context) (int64, error) {
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		if db.storage == nil {
			return -1, fmt.Errorf("storage layer is nil")
		}

		return db.storage.Count(ctx)
	}

	return db.search.Count(ctx)
}

// getEntityID extracts the ID from an entity (you need to customize this function)
func getEntityID[T Entity[ID], ID comparable](entity T) ID {
	// This function should extract the ID field from the entity.
	// It depends on your entity structure, you may need reflection or a defined method.
	return entity.GetID()
}
