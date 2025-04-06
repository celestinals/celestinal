// Copyright 2025 The Celestinal Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos

import (
	"context"

	"github.com/jackc/pgx/v5"

	cestdb "github.com/celestinals/celestinal/pkg/database"
	"github.com/celestinals/celestinal/pkg/database/sql"
	"github.com/celestinals/celestinal/staging/mod/greeter/models/gen/authors"
)

// Check valid Object with Interface
var _ cestdb.Repository[authors.Author, int64] = (*Authors)(nil)
var _ IAuthors = (*Authors)(nil)

// IAuthors define for mockup cestdb
type IAuthors interface {
	Create(ctx context.Context, author authors.Author) (authors.Author, error)
	Update(ctx context.Context, id int64, author authors.Author) (authors.Author, error)
	Get(ctx context.Context, id int64) (authors.Author, error)
	GetAll(ctx context.Context) ([]authors.Author, error)
	Delete(ctx context.Context, id int64) error
	Exists(ctx context.Context, id int64) (bool, error)
	Count(ctx context.Context) (int64, error)
}

// NewAuthor creates a new cestdb repository.
func NewAuthor(pgCon *pgx.Conn) *Authors {
	storage := sql.New[authors.Author, int64](pgCon, "authors")

	return &Authors{
		SQL:   storage,
		query: authors.New(pgCon),
	}
}

// Authors repository
type Authors struct {
	*sql.SQL[authors.Author, int64]
	query *authors.Queries
}

// Create method implement sql.SQL.Create
func (a *Authors) Create(ctx context.Context, entity authors.Author) (authors.Author, error) {
	resp, err := a.query.Create(ctx, authors.CreateParams{
		Name: entity.Name,
		Bio:  entity.Bio,
	})

	return resp, err
}

// Update method implement sql.SQL.Update
func (a *Authors) Update(ctx context.Context, id int64, author authors.Author) (authors.Author, error) {
	resp, err := a.query.Update(ctx, authors.UpdateParams{
		ID:   id,
		Name: author.Name,
		Bio:  author.Bio,
	})

	return resp, err
}
