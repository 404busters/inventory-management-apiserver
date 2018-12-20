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
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"net/http"
)

type inventoryHandler struct {
	Service core.InventoryService
}

func (h *inventoryHandler) locationList(c *gin.Context) {
	locationId := c.Param("id")
	inventories, err := h.Service.LocationList(c, locationId)

	if err == core.ErrReferencrNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("location˝ %s are not exists", locationId),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: inventories,
		})
	}
}

func (h *inventoryHandler) itemTypeList(c *gin.Context) {
	itemTypeId := c.Param("id")
	inventories, err := h.Service.ItemTypeList(c, itemTypeId)

	if err == core.ErrReferencrNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("itemtype˝ %s are not exists", itemTypeId),
		})
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: inventories,
		})
	}
}

func (h *inventoryHandler) get(c *gin.Context) {
	id := c.Param("id")
	inventory, err := h.Service.Get(c, id)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else if inventory == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_found",
			Message: fmt.Sprintf("inventroy %s is not exists", id),
		})
		return
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: inventory,
		})
	}
}

func (h *inventoryHandler) create(c *gin.Context) {
	var inventoryInput core.Inventory
	err := c.ShouldBindJSON(&inventoryInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	if inventoryInput.Status != core.StatusStock && inventoryInput.Status != core.StatusRepair {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	inventory, err := h.Service.Create(c, &inventoryInput)
	if err == core.ErrReferencrNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("location˝ %s or item_type %s are not exists", inventoryInput.Location, inventoryInput.ItemType),
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

func (h *inventoryHandler) update(c *gin.Context) {
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

	if inventoryInput.Status != core.StatusStock && inventoryInput.Status != core.StatusRepair {
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
	} else if err == core.ErrReferencrNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("location %s or item_type %s are not exists", inventoryInput.Location, inventoryInput.ItemType),
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

func (h *inventoryHandler) delete(c *gin.Context) {
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
