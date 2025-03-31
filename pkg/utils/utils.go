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

// Package utils provides utility functions for the service.
package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
)

// CallBack calls the function and ignores the error.
func CallBack(fn func() error) {
	_ = fn()
}

// NewPgxPool create new pool connection for multiple query
func NewPgxPool(conf *tickex.Config) (*pgxpool.Pool, error) {
	_ = conf
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		"config.DbUser", "config.DbPassword", "config.DbHost", "config.DbPort", "config.DbName")

	return pgxpool.New(context.Background(), dsn)
}

// NewPgxConn create new connection for single query
func NewPgxConn(ctx context.Context, conf *tickex.Config) (*pgx.Conn, error) {
	_ = conf
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		"config.DbUser", "config.DbPassword", "config.DbHost", "config.DbPort", "config.DbName")

	return pgx.Connect(ctx, dsn)
}
