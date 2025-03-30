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

package pg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/pkg/driver/db"
	"github.com/tickexvn/tickex/pkg/utils"
)

var _ db.Driver[int] = (*Driver[int])(nil)

func New[T any](conf *tickex.Config) (*Driver[T], error) {
	pool, err := utils.NewPgxPool(conf)
	if err != nil {
		return nil, err
	}

	return &Driver[T]{pool: pool}, nil
}

type Driver[T any] struct {
	pool *pgxpool.Pool
}

func (d *Driver[T]) Query(ctx context.Context, sql string, args ...any) (db.Rows[T], error) {
	pgRows, err := d.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dbRows[T]{rows: pgRows}, nil
}
