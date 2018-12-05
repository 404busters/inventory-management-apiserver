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
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"net/http"
)

type inventoryHandler struct {
	Service core.InventoryService
	Logger  logrus.FieldLogger
}

func (h *inventoryHandler) Create(c *gin.Context) {
	var inventoryInput core.Inventory
	err := c.ShouldBindJSON(&inventoryInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	inventory, err := h.Service.Create(c, &inventoryInput)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: inventory,
		})
	}
}

func (h *inventoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var inventoryInput core.Inventory
	err := c.ShouldBindJSON(&inventoryInput)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	inventory, err := h.Service.Update(c, id, &inventoryInput)
	if err == nil && inventory == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("inventory %s is not exists", id),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: inventory,
		})
	}
}

func (h *inventoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.Delete(c, id); err == core.ErrRecordNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("inventory %s is not exists", id),
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
