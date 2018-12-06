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

package http

import (
	"context"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/service/postgres"
	"os"

	_ "github.com/lib/pq"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/inject"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/logging"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

func bindService(base context.Context) (ctx context.Context) {
	connector := &dbutils.Connector{Driver: "postgres", DataSource: os.Getenv("DATABASE_URL")}
	logger := logging.GetRoot()

	{
		service := &postgres.LocationService{Connector: connector, Logger: logger.WithField("service", "location")}
		ctx = inject.BindLocationServiceToContext(base, service)
	}

	{
		service := &postgres.ItemTypeService{Connector: connector, Logger: logger.WithField("service", "item_type")}
		ctx = inject.BindItemTypeServiceToContext(ctx, service)
	}

	ctx = inject.BindUserServiceToContext(ctx, &postgres.UserService{
		Connector: connector,
		Logger:    logger.WithField("service", "user"),
	})

	// TODO Add Other Service
	return
}
