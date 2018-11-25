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
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

// For static type checking
var _ core.LocationService = &LocationService{}

// TODO: implement
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
		s.Logger.Debug(location)
	}

	if err := rows.Err(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return locations, nil
}

func (s *LocationService) Get(ctx context.Context, id string) (*core.Location, error) {
	return nil, nil
}

func (s *LocationService) Create(ctx context.Context, input *core.LocationInput) (*core.Location, error) {
	return nil, nil
}

func (s *LocationService) Update(ctx context.Context, id string, input *core.LocationInput) (*core.Location, error) {
	return nil, nil
}

func (s *LocationService) Delete(ctx context.Context, id string) error {
	return nil
}
