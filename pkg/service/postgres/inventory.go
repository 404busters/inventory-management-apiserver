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

var _ core.InventoryService = &InventoryService{}

type InventoryService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s *InventoryService) Get(ctx context.Context, id string) (*core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) Create(ctx context.Context, input *core.Inventory) (*core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) Update(ctx context.Context, id string, input *core.Inventory) (*core.Inventory, error) {
	panic("implement me")
}

func (s *InventoryService) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
