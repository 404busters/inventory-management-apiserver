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

type itemTypeHandler struct {
	Service core.ItemTypeService
}

func (h *itemTypeHandler) list(c *gin.Context) {
	entries, err := h.Service.List(c)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{Data: entries})
	}
}

func (h *itemTypeHandler) get(c *gin.Context) {
	id := c.Param("id")
	entry, err := h.Service.Get(c, id)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
		return
	} else if entry == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_found",
			Message: fmt.Sprintf("item type %s is not exists", id),
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{Data: entry})
}

func (h *itemTypeHandler) create(c *gin.Context) {
	var input core.ItemType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	entry, err := h.Service.Create(c, &input)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{Data: entry})
}

func (h *itemTypeHandler) update(c *gin.Context) {
	id := c.Param("id")
	var input core.ItemType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	entry, err := h.Service.Update(c, id, &input)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
		return
	} else if entry == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("item type %s not exists", id),
		})
		return
	}

	c.JSON(http.StatusOK, ApiRes{Data: entry})
}

func (h *itemTypeHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(c, id); err == core.ErrRecordNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_found",
			Message: fmt.Sprintf("item type %s is not exists", id),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
