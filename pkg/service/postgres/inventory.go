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

var _ core.InventoryService = &InventoryService{}

type InventoryService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *InventoryService) ItemTypeList(ctx context.Context, itemTypeId string) ([]core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) LocationList(ctx context.Context, locationId string) ([]core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) Get(ctx context.Context, id string) (*core.Inventory, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	var location core.Inventory

	row := conn.QueryRowContext(ctx, "SELECT id, item_type, last_seen_location, status, last_seen_time  FROM inventory WHERE id = $1", id)
	if err := row.Scan(&location.Id); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &location, nil
}

func (s *InventoryService) Create(ctx context.Context, input *core.Inventory) (*core.Inventory, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	result, err := conn.ExecContext(ctx, "SELECT name FROM location WHERE id = $1", input.Location)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	cnt, err := result.RowsAffected()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	} else if cnt < 1 {
		return nil, core.ErrReferencrNotExists
	}

	result, err = conn.ExecContext(ctx, "SELECT name FROM item_type WHERE id = $1", input.ItemType)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	cnt, err = result.RowsAffected()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	} else if cnt < 1 {
		return nil, core.ErrReferencrNotExists
	}

	var inventory core.Inventory

	tx, err := conn.BeginTx(ctx, nil)
	row := tx.QueryRowContext(ctx, "INSERT INTO inventory (id, item_type, last_seen_location, status, last_seen_time, created_at, updated_at) VALUES ($1, $2, $3, $4,current_timestamp ,current_timestamp,current_timestamp) RETURNING id, item_type, last_seen_location, status,last_seen_time",
		uuid.NewV4(), input.ItemType, input.Location, input.Status)
	defer tx.Rollback()

	if err := row.Scan(&inventory.Id, &inventory.ItemType, &inventory.Location, &inventory.Status, &inventory.LastSeenTime); err != nil {
		s.Logger.Error(err)
		return nil, nil
	}

	if err := tx.Commit(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return &inventory, nil
}

func (s *InventoryService) Update(ctx context.Context, id string, input *core.Inventory) (*core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) Delete(ctx context.Context, id string) error {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)

	if err != nil {
		s.Logger.Error(err)
		return err
	}

	result, err := tx.Exec("UPDATE inventory SET deleted_at= current_timestamp WHERE id = $1", id)
	defer tx.Rollback()

	if err != nil {
		s.Logger.Error(err)
		return err
	}

	cnt, err := result.RowsAffected()
	if err != nil {
		s.Logger.Error(err)
		return err
	} else if cnt < 1 {
		return core.ErrRecordNotExists
	}

	if err := tx.Commit(); err != nil {
		s.Logger.Error(err)
		return err
	}

	return nil
}
