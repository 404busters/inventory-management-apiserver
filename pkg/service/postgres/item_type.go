/*
	Copyright 2018 Carmen Chan & Tony Yip

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package postgres

import (
	"context"
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

// For static type checking
var _ core.ItemTypeService = &ItemTypeService{}

type ItemTypeService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *ItemTypeService) List(ctx context.Context) (_ []core.ItemType, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	const query = "SELECT id, name, description FROM item_type WHERE deleted_at IS NULL"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]core.ItemType, 0)
	for rows.Next() {
		var itemType core.ItemType
		if err := rows.Scan(&itemType.Id, &itemType.Name, &itemType.Description); err != nil {
			s.Logger.Error(err)
			return nil, err
		}
		result = append(result, itemType)
	}

	if err := rows.Err(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *ItemTypeService) Get(ctx context.Context, id string) (result *core.ItemType, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	const query = "SELECT id, name, description FROM item_type WHERE id = $1 AND deleted_at IS NULL"
	row := conn.QueryRowContext(ctx, query, id)
	result = new(core.ItemType)
	if err = row.Scan(&result.Id, &result.Name, &result.Description); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return
}

func (s *ItemTypeService) Create(ctx context.Context, input *core.ItemType) (result *core.ItemType, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer tx.Rollback()

	const query = "INSERT INTO item_type (id, name, description) VALUES ($1, $2, $3) RETURNING id, name, description"
	row := tx.QueryRowContext(ctx, query, uuid.NewV4(), input.Name, input.Description)

	result = new(core.ItemType)
	if err = row.Scan(&result.Id, &result.Name, &result.Description); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return
}

func (s *ItemTypeService) Update(ctx context.Context, id string, input *core.ItemType) (result *core.ItemType, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer tx.Rollback()

	const query = "UPDATE item_type SET name = $2, description = $3, updated_at = current_timestamp WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, description"
	row := tx.QueryRowContext(ctx, query, id, input.Name, input.Description)

	result = new(core.ItemType)
	if err = row.Scan(&result.Id, &result.Name, &result.Description); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return
}

func (s *ItemTypeService) Delete(ctx context.Context, id string) (err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer tx.Rollback()

	const query = "UPDATE item_type SET deleted_at = current_timestamp WHERE id = $1 AND deleted_at IS NULL"
	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		s.Logger.Error(err)
		return
	}

	if affected, err := result.RowsAffected(); err != nil {
		s.Logger.Error(err)
		return err
	} else if affected != 1 {
		return core.ErrRecordNotExists
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return
	}

	return nil
}
