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

func CreateHandler(ctx context.Context) http.Handler {
	app := gin.New()
	basePath := "/api"

	v1 := app.Group(basePath + "/v1")

	logger := inject.GetLoggerFromContext(ctx)

	{
		service := inject.GetLocationServiceFromContext(ctx)
		handler := locationHandler{
			Service: service,
			Logger:  logger.WithField("controller", "location"),
		}
		v1.GET("/location", handler.list)
		v1.GET("/location/:id", handler.Get)
		v1.POST("/location", handler.Create)
		v1.PATCH("/location/:id", handler.Update)
		v1.DELETE("/location/:id", handler.Delete)
	}

	{
		handler := itemTypeHandler{
			Service: inject.GetItemTypeServiceFromContext(ctx),
		}
		v1.GET("/itemType", handler.list)
		v1.GET("/itemType/:item_type", handler.get)
		v1.POST("/itemType", handler.create)
		v1.PATCH("/itemType/:item_type", handler.update)
		v1.DELETE("/itemType/:item_type", handler.delete)
	}

	{
		handler := userHandler{
			service: inject.GetUserServiceFromContext(ctx),
		}
		v1.GET("/user", handler.list)
		v1.GET("/user/:id", handler.get)
		v1.POST("/user", handler.create)
		v1.PATCH("/user/:id", handler.update)
		v1.DELETE("/user/:id", handler.delete)
	}

	return app
}
