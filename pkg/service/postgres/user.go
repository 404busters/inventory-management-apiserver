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

type UserService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *UserService) List(ctx context.Context) (_ []core.User, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	const query = "SELECT id, name FROM user"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer rows.Close()

	users := make([]core.User, 0)
	for rows.Next() {
		var user core.User
		if err = rows.Scan(&user.Id, &user.Name); err != nil {
			s.Logger.Error(err)
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Error(err)
		return
	}

	return users, nil
}

func (s *UserService) Get(ctx context.Context, id string) (_ *core.User, err error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	defer conn.Close()

	const query = "SELECT id, name FROM user WHERE id = $1"
	row := conn.QueryRowContext(ctx, query, id)
	user := new(core.User)
	if err = row.Scan(&user.Id, &user.Name); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		s.Logger.Error(err)
		return
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, input *core.User) (_ *core.User, err error) {
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

	const query = "INSERT INTO user (id, name) VALUES ($1, $2) RETURNING id, name"
	row := tx.QueryRowContext(ctx, query, uuid.NewV4(), input.Name)
	user := new(core.User)
	if err = row.Scan(&user.Id, &user.Name); err != nil {
		s.Logger.Error(err)
		return
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id string, input *core.User) (_ *core.User, err error) {
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

	const query = "UPDATE user SET name = $2, updated_at = current_timestamp WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, uuid.NewV4(), input.Name)
	user := new(core.User)
	if err = row.Scan(&user.Id, &user.Name); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		s.Logger.Error(err)
		return
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id string) (err error) {
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

	const query = "DELETE FROM user WHERE id = $1"
	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		s.Logger.Error(err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		s.Logger.Error(err)
		return
	} else if affected != 1 {
		return core.ErrRecordNotExists
	}

	if err = tx.Commit(); err != nil {
		s.Logger.Error(err)
		return
	}

	return nil
}
