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
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

type locationHandler struct {
	Service core.LocationService
	Logger  logrus.FieldLogger
}

func (h *locationHandler) list(c *gin.Context) {
	locations, err := h.Service.List(c)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: locations,
		})
	}
}

func (h *locationHandler) Get(c *gin.Context) {
	id := c.Param("id")
	location, err := h.Service.Get(c, id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: err.Error(),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: location,
		})
	}
}

func (h *locationHandler) Create(c *gin.Context) {
	var locationInput core.Location
	err := c.ShouldBindJSON(&locationInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "",
			Message: err.Error(),
		})
		return
	}

	location, err := h.Service.Create(c, &locationInput)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: location,
		})
	}
}

func (h *locationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var locationInput core.Location
	err := c.ShouldBindJSON(&locationInput)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "",
			Message: err.Error(),
		})
		return
	}

	location, err := h.Service.Update(c, id, &locationInput)
	if err == nil && location == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: err.Error(),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: location,
		})
	}
}

func (h *locationHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.Delete(c, id)
	if err.Error() == "item_not_Found" {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: err.Error(),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}
