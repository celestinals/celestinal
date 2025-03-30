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
	"github.com/jackc/pgx/v5"
	"github.com/tickexvn/tickex/pkg/driver/db"
)

var _ db.Rows[int] = (*dbRows[int])(nil)

type dbRows[T any] struct {
	rows pgx.Rows
}

func (r *dbRows[T]) CollectOne() (T, error) {
	return pgx.CollectOneRow[T](r.rows, pgx.RowToStructByName[T])
}

func (r *dbRows[T]) CollectAll() ([]T, error) {
	return pgx.CollectRows[T](r.rows, pgx.RowToStructByName[T])
}
