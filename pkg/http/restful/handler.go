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
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler() http.Handler {
	app := gin.New()
	basePath := "/api"

	v1 := app.Group(basePath + "/v1")

	inventory := v1.Group("/inventory")
	{
		inventory.GET("/", getInventory)
	}

	return app
}

func getInventory(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
