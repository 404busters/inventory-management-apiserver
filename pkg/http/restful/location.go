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
		c.Status(http.StatusServiceUnavailable)
	} else {
		c.JSON(http.StatusOK, locations)
	}
}
