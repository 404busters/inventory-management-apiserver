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
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

type userHandler struct {
	service core.UserService
}

func (h *userHandler) list(c *gin.Context) {
	users, err := h.service.List(c)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: users,
		})
	}
}

func (h *userHandler) get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.Get(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else if user == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_found",
			Message: fmt.Sprintf("user %s not exists", id),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: user,
		})
	}
}

func (h *userHandler) create(c *gin.Context) {
	var input core.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}
	user, err := h.service.Create(c, &input)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ApiRes{
			Data: user,
		})
	}
}

func (h *userHandler) update(c *gin.Context) {
	id := c.Param("id")
	var input core.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	entry, err := h.service.Update(c, id, &input)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorRes{
			Code:    "database_error",
			Message: err.Error(),
		})
		return
	} else if entry == nil {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_Found",
			Message: fmt.Sprintf("user %s not exists", id),
		})
		return
	}

	c.JSON(http.StatusOK, ApiRes{Data: entry})
}

func (h *userHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c, id); err == core.ErrRecordNotExists {
		c.JSON(http.StatusNotFound, ErrorRes{
			Code:    "item_not_found",
			Message: fmt.Sprintf("user %s is not exists", id),
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
