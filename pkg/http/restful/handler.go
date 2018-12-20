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

package restful

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/inject"
)

type resourceHandler interface {
	list(*gin.Context)
	get(*gin.Context)
	create(*gin.Context)
	update(*gin.Context)
	delete(*gin.Context)
}

func CreateHandler(ctx context.Context) http.Handler {
	app := gin.New()
	basePath := "/api"

	v1 := app.Group(basePath + "/v1")

	handlers := map[string]resourceHandler{
		"user": &userHandler{
			service: inject.GetUserServiceFromContext(ctx),
		},
		"itemType": &itemTypeHandler{
			Service: inject.GetItemTypeServiceFromContext(ctx),
		},
		"location": &locationHandler{
			Service: inject.GetLocationServiceFromContext(ctx),
		},
	}

	for prefix, handler := range handlers {
		regularPath := "/" + prefix
		idPath := regularPath + "/:id"
		v1.GET(regularPath, handler.list)
		v1.GET(idPath, handler.get)
		v1.POST(regularPath, handler.create)
		v1.PATCH(idPath, handler.update)
		v1.DELETE(idPath, handler.delete)
	}

	{
		inventoryHandler := inventoryHandler{
			Service: inject.GetInventoryServiceFromContext(ctx),
		}

		inventoryPath := "/inventory"
		idPath := inventoryPath + "/:id"
		v1.GET(inventoryPath + "/location/:locationId", inventoryHandler.locationList)
		v1.GET(inventoryPath + "/itemType/:itemTypeId", inventoryHandler.itemTypeList)
		v1.GET(idPath, inventoryHandler.get)
		v1.POST(inventoryPath, inventoryHandler.create)
		v1.PATCH(idPath, inventoryHandler.update)
		v1.DELETE(idPath, inventoryHandler.delete)
	}

	return app
}
