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

	"github.com/tickexvn/tickex/pkg/errors"
	"github.com/tickexvn/tickex/pkg/txlog"
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
func (db *Database[T, ID]) Create(ctx context.Context, entity T) (T, error) {
	var empty T
	if db.storage == nil {
		txlog.Debug("[database.Create] storage layer is nil")
		return empty, errors.F("database.Create err: storage is nil")
	}

	// insert into PostgreSQL first
	createdEntity, err := db.storage.Create(ctx, entity)
	if err != nil {
		txlog.Debugf("[database.Create] storage err: %v", err)
		return empty, errors.F("database.Create err: %v", err)
	}

	// if search layer is nil, return created entity
	if db.search == nil {
		txlog.Infof("[database.Create] search => pass")
		return createdEntity, nil
	}

	// insert into Elasticsearch
	// ignore error if exist
	_, err = db.search.Create(ctx, entity)
	if err != nil {
		txlog.Warnf("[database.Create] search err: %v", err)
	}

	return createdEntity, nil
}

// Update modifies an existing entity in both storage (PostgreSQL) and
// search (Elasticsearch).If Elasticsearch fails after PostgreSQL succeeds,
func (db *Database[T, ID]) Update(ctx context.Context, id ID, entity T) (T, error) {
	var empty T
	if db.storage == nil {
		txlog.Debug("[database.Update] storage layer is nil")
		return empty, errors.F("database.Update err: storage is nil")
	}

	// update in PostgreSQL
	updatedEntity, err := db.storage.Update(ctx, id, entity)
	if err != nil {
		txlog.Debugf("[database.Update] storage err: %v", err)
		return empty, errors.F("database.Update err: %v", err)
	}

	// if search layer is nill, return updated entity
	if db.search == nil {
		txlog.Info("[database.Update] search => pass")
		return updatedEntity, nil
	}

	// update in Elasticsearch if they are exist
	if existed, err := db.search.Exists(ctx, id); err == nil && existed {
		_, err = db.search.Update(ctx, id, updatedEntity)
		if err != nil {
			txlog.Warnf("[database.Update] search err: %v", err)
		}
	}

	return updatedEntity, nil
}

// Delete removes an entity from both storage (PostgreSQL) and search
// (Elasticsearch). If Elasticsearch fails after PostgreSQL succeeds,
func (db *Database[T, ID]) Delete(ctx context.Context, id ID) error {
	if db.storage == nil {
		txlog.Debug("[database.Delete] err: storage is nil")
		return errors.F("database.Delete err: storage is nil")
	}

	// if search layer is nil, jump to storage layer
	if db.search == nil {
		txlog.Info("[database.Delete] search => pass")
		if err := db.storage.Delete(ctx, id); err != nil {
			txlog.Debugf("[database.Delete][search] err: %v", err)
			return errors.F("database.Delete storage err: %v", err)
		}

		// end of function
		return nil
	}

	// delete from Elasticsearch, if error, return error
	if esErr := db.search.Delete(ctx, id); esErr != nil {
		txlog.Debugf("[database.Delete][search] err: %v", esErr)
		return errors.F("database.Delete search err: %v", esErr)
	}

	// after delete from search, delete from storage
	if err := db.storage.Delete(ctx, id); err != nil {
		txlog.Debugf("[database.Delete][storage] err: %v", err)
		return errors.F("database.Delete storage err: %v", err)
	}

	return nil
}

// Get retrieves an entity by ID from the search layer (Elasticsearch).
func (db *Database[T, ID]) Get(ctx context.Context, id ID) (T, error) {
	var empty T

	// if storage layer is nil, return error
	if db.storage == nil {
		txlog.Debug("[database.Get] storage layer is nil")
		return empty, errors.F("database.Get err: storage is nil")
	}

	// if search layer is nil, jump to storage layer
	if db.search == nil {
		txlog.Info("[database.Get] search => pass")

		data, err := db.storage.Get(ctx, id)
		if err != nil {
			txlog.Debugf("[database.Get][search] err: %v", err)
			return empty, errors.F("database.Get search err: %v", err)
		}

		return data, nil
	}

	// search data in search layer
	result, err := db.search.Get(ctx, id)
	if err != nil {
		txlog.Debugf("[database.Get][search] err: ", err)

		// if search layer return error, get data from storage layer
		data, err := db.storage.Get(ctx, id)
		if err != nil {
			txlog.Debugf("[database.Get][storage] err: %v", err)
			return empty, errors.F("database.Get storage err: %v", err)
		}

		// not existed in search layer, create new one
		existed, err := db.search.Exists(ctx, id)
		if err != nil {
			txlog.Warnf("[database.Get][search] err: %v", err)
		}

		if !existed {
			if _, err := db.search.Create(ctx, data); err != nil {
				txlog.Warnf("[database.Get][search] err: %v", err)
			}
		}

		return data, nil
	}

	return result, nil
}

// GetAll retrieves all entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) GetAll(ctx context.Context) ([]T, error) {
	if db.storage == nil {
		txlog.Debug("[database.GetAll] err: storage is nil")
		return nil, errors.F("database.GetAll err: storage is nil")
	}

	if db.search == nil {
		txlog.Info("[database.GetAll] search => pass")
		resp, err := db.storage.GetAll(ctx)
		if err != nil {
			txlog.Debugf("[database.GetAll][search] err: %v", err)
			return nil, errors.F("database.GetAll search err: %v", err)
		}

		return resp, nil
	}

	ts, err := db.search.GetAll(ctx)
	if err != nil || len(ts) == 0 {
		if err != nil {
			txlog.Debugf("[database.GetAll][search] err: %v", err)
		}
		// if search layer return error or empty, get data from storage layer
		result, errStorageGetAll := db.storage.GetAll(ctx)
		if errStorageGetAll != nil {
			txlog.Debugf("[database.GetAll][storage] err: %v", errStorageGetAll)
			return nil, errors.F("database,GetAll storage err: %v", errStorageGetAll)
		}

		return result, nil
	}

	return ts, nil
}

// Exists checks if an entity exists in the search layer (Elasticsearch).
func (db *Database[T, ID]) Exists(ctx context.Context, id ID) (bool, error) {
	if db.storage == nil {
		txlog.Debug("[database.Exists] storage layer is nil")
		return false, errors.F("database.Exists storage layer is nil")
	}

	if db.search != nil {
		existed, err := db.search.Exists(ctx, id)
		if err == nil {
			return existed, nil
		}

		txlog.Warnf("[database.Exists] search err: %v", err)
	}

	txlog.Info("[database.Exists] search => pass")
	existed, err := db.storage.Exists(ctx, id)
	if err != nil {
		txlog.Debugf("[database.Exists] storage err: ", err)
		return false, errors.F("database.Exists storage err: %v", err)
	}

	return existed, nil
}

// Count returns the total number of entities from the search layer (Elasticsearch).
func (db *Database[T, ID]) Count(ctx context.Context) (int64, error) {
	if db.storage == nil {
		txlog.Debug("[database.Count][storage] layer is nil")
		return -1, errors.F("database.Count storage layer is nil")
	}

	if db.search != nil {
		count, err := db.search.Count(ctx)
		if err == nil {
			return count, nil
		}

		txlog.Warnf("[database.Count][search] err: %v", err)
	}

	txlog.Info("[database.Exists] search => pass")
	count, err := db.storage.Count(ctx)
	if err != nil {
		txlog.Debugf("[database.Count][storage] err: %v", err)
		return -1, errors.F("database.Count storage err: %v", err)
	}

	return count, nil
}
