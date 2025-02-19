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

package repos

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/tickexvn/tickex/pkg/database"
	"github.com/tickexvn/tickex/pkg/database/pg"
	"github.com/tickexvn/tickex/x/greeter/v1/internal/models/gen/authors"
)

var _ database.Repository[authors.Author, int64] = (*Authors)(nil)

// NewAuthor creates a new database repository.
func NewAuthor(pgCon *pgx.Conn) *Authors {
	storage := pg.NewStorageLayer[authors.Author, int64](pgCon, "authors")

	return &Authors{
		BaseStorageLayer: storage,
		query:            authors.New(pgCon),
	}
}

// Authors repository
type Authors struct {
	*pg.BaseStorageLayer[authors.Author, int64]
	query *authors.Queries
}

// Create method implement BaseStorageLayer.Create
func (a *Authors) Create(ctx context.Context, entity authors.Author) (authors.Author, error) {
	resp, err := a.query.Create(ctx, authors.CreateParams{
		Name: entity.Name,
		Bio:  entity.Bio,
	})
	return resp, err
}
