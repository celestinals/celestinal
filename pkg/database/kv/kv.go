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

// Package kv implement key-value storage base on postgresql
package kv

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"

	google "google.golang.org/protobuf/proto"
)

var _ KeyValue = (*KV)(nil)

// New creates a new KeyValue instance.
func New(tableName string, conf *tickex.Config) KeyValue {
	return &KV{
		tableName: tableName,
	}
}

// KeyValue is an interface for key-value storage.
type KeyValue interface {
	Create(ctx context.Context, entity google.Message) (google.Message, error)
	Get(ctx context.Context, id string) (google.Message, error)
	GetAll(ctx context.Context) ([]google.Message, error)
	Update(ctx context.Context, id string, entity google.Message) (google.Message, error)
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
	Count(ctx context.Context) (int64, error)
}

// KV is a key-value storage.
type KV struct {
	tableName  string
	connection *pgx.Conn
}

// Count implements KeyValue.
func (k *KV) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// Create implements KeyValue.
func (k *KV) Create(ctx context.Context, entity google.Message) (google.Message, error) {
	panic("unimplemented")
}

// Delete implements KeyValue.
func (k *KV) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Exists implements KeyValue.
func (k *KV) Exists(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// Get implements KeyValue.
func (k *KV) Get(ctx context.Context, id string) (google.Message, error) {
	panic("unimplemented")
}

// GetAll implements KeyValue.
func (k *KV) GetAll(ctx context.Context) ([]google.Message, error) {
	panic("unimplemented")
}

// Update implements KeyValue.
func (k *KV) Update(ctx context.Context, id string, entity google.Message) (google.Message, error) {
	panic("unimplemented")
}
