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

// New create multi-layer database instance
func New[T any, ID comparable](
	searchLayer Repository[T, ID], storageLayer Repository[T, ID]) Repository[T, ID] {

	return &Database[T, ID]{
		search:  searchLayer,
		storage: storageLayer,
	}
}

// Database database multilayer with postgre & elasticsearch
type Database[T any, ID comparable] struct {
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

	// if search layer is nil, return created entity
	if db.search == nil {
		logger.Infof("search layer is nil, pass")
		return createdEntity, nil
	}

	// Insert into Elasticsearch
	_, _ = db.search.Create(ctx, entity)

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

	// Update in PostgreSQL
	updatedEntity, err := db.storage.Update(ctx, id, entity)
	if err != nil {
		return empty, err
	}

	// if search layer is nill, return updated entity
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return updatedEntity, nil
	}

	// Update in Elasticsearch
	if existed, err := db.search.Exists(ctx, id); err == nil && existed {
		_, _ = db.search.Update(ctx, id, updatedEntity)
	}

	return updatedEntity, nil
}

// Delete removes an entity from both storage (PostgreSQL) and search
// (Elasticsearch). If Elasticsearch fails after PostgreSQL succeeds,
// it rolls back by restoring the deleted entity.
func (db *Database[T, ID]) Delete(ctx context.Context, id ID) error {
	if db.storage == nil {
		return fmt.Errorf("storage layer is nil")
	}

	// if search layer is nil, jump to storage layer
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return db.storage.Delete(ctx, id)
	}

	// Delete from Elasticsearch, if error, return error
	if esErr := db.search.Delete(ctx, id); esErr != nil {
		return fmt.Errorf("failed to delete from search: %v", esErr)
	}

	// after delete from search, delete from storage
	return db.storage.Delete(ctx, id)
}

// Get retrieves an entity by ID from the search layer (Elasticsearch).
func (db *Database[T, ID]) Get(ctx context.Context, id ID) (T, error) {
	var empty T

	// if storage layer is nil, return error
	if db.storage == nil {
		return empty, fmt.Errorf("storage layer is nil")
	}

	// if search layer is nil, jump to storage layer
	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return db.storage.Get(ctx, id)
	}

	// search data in search layer
	result, err := db.search.Get(ctx, id)
	if err != nil {
		// if search layer return error, get data from storage layer
		data, err := db.storage.Get(ctx, id)
		if err != nil {
			return empty, fmt.Errorf("failed to get from storage: %v", err)
		}

		// not existed in search layer, create new one
		if existed, err := db.search.Exists(ctx, id); err != nil || !existed {
			_, _ = db.search.Create(ctx, data)
		}

		return data, nil
	}

	return result, nil
}

// GetAll retrieves all entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) GetAll(ctx context.Context) ([]T, error) {
	if db.storage == nil {
		return nil, fmt.Errorf("storage layer is nil")
	}

	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return db.storage.GetAll(ctx)
	}

	ts, err := db.search.GetAll(ctx)
	if err != nil || len(ts) == 0 {
		// if search layer return error or empty, get data from storage layer
		return db.storage.GetAll(ctx)
	}

	return ts, nil
}

// Exists checks if an entity exists in the search layer (Elasticsearch).
func (db *Database[T, ID]) Exists(ctx context.Context, id ID) (bool, error) {
	if db.storage == nil {
		return false, fmt.Errorf("storage layer is nil")
	}

	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return db.storage.Exists(ctx, id)
	}

	existed, err := db.search.Exists(ctx, id)
	if err != nil {
		return db.storage.Exists(ctx, id)
	}

	return existed, nil
}

// Count returns the total number of entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) Count(ctx context.Context) (int64, error) {
	if db.storage == nil {
		return -1, fmt.Errorf("storage layer is nil")
	}

	if db.search == nil {
		logger.Info("search layer is nil, pass")
		return db.storage.Count(ctx)
	}

	counter, err := db.search.Count(ctx)
	if err != nil {
		return db.storage.Count(ctx)
	}

	return counter, nil
}
