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
	"errors"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

// For static type checking
var _ core.LocationService = &LocationService{}

type LocationService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *LocationService) List(ctx context.Context) ([]core.Location, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT id, name FROM location")
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	var locations = make([]core.Location, 0)

	for rows.Next() {
		var location core.Location
		if err := rows.Scan(&location.Id, &location.Name); err != nil {
			s.Logger.Error(err)
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return locations, nil
}

func (s *LocationService) Get(ctx context.Context, id string) (*core.Location, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	var location core.Location

	row := conn.QueryRowContext(ctx, "SELECT id, name FROM location WHERE id = $1", id)
	if err := row.Scan(&location.Id, &location.Name); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &location, nil
}

func (s *LocationService) Create(ctx context.Context, input *core.Location) (*core.Location, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	var location core.Location

	tx, err := conn.BeginTx(ctx, nil)

	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	row := tx.QueryRowContext(ctx, "INSERT INTO location (id ,name) VALUES ( $1, $2) RETURNING id, name", uuid.NewV4(), input.Name)
	defer tx.Rollback()

	if err := row.Scan(&location.Id, &location.Name); err != nil {
		s.Logger.Error(err)
		return nil, nil
	}

	if err := tx.Commit(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return &location, nil
}

func (s *LocationService) Update(ctx context.Context, id string, input *core.Location) (*core.Location, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	var location core.Location

	tx, err := conn.BeginTx(ctx, nil)

	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	row := tx.QueryRowContext(ctx, "UPDATE location SET name = $1, updated_at = current_timestamp WHERE id = $2 RETURNING id, name", input.Name, id)
	defer tx.Rollback()

	if err := row.Scan(&location.Id, &location.Name); err != nil {
		s.Logger.Error(err)
		return nil, nil
	}

	if err := tx.Commit(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return &location, nil
}

func (s *LocationService) Delete(ctx context.Context, id string) error {
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

	result, err := tx.Exec("DELETE FROM location WHERE id = $1", id)
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
		return errors.New("item_not_Found")
	}

	if err := tx.Commit(); err != nil {
		s.Logger.Error(err)
		return err
	}

	return nil
}
