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

package mock

import (
	"context"
	"errors"

	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

var ErrNotExists = errors.New("target id not exists")

// For static type checking
var _ core.LocationService = &LocationService{}

var locations = []core.Location{
	{Id: "3e7654f0-cf60-484a-afa1-43a8837ebd7b", Name: "Location 1"},
	{Id: "6b3a5aac-dbf6-408f-9248-b087080ef939", Name: "Location 2"},
}

type LocationService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *LocationService) List(ctx context.Context) ([]core.Location, error) {
	return locations, nil
}

func (s *LocationService) Get(ctx context.Context, id string) (*core.Location, error) {
	for _, location := range locations {
		if location.Id == id {
			return &location, nil
		}
	}
	return nil, nil
}

func (s *LocationService) Create(ctx context.Context, input *core.Location) (*core.Location, error) {
	id := uuid.NewV4().String()
	location := core.Location{Id: id, Name: input.Name}
	locations = append(locations, location)
	return &location, nil
}

func (s *LocationService) Update(ctx context.Context, id string, input *core.Location) (*core.Location, error) {
	for _, location := range locations {
		if location.Id == id {
			location.Name = input.Name
			return &location, nil
		}
	}
	return nil, ErrNotExists
}

func (s *LocationService) Delete(ctx context.Context, id string) error {
	for idx, location := range locations {
		if location.Id == id {
			locations = append(locations[:idx], locations[idx+1:]...)
		}
	}
	return ErrNotExists
}
