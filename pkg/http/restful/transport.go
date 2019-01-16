package restful

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"net/http"
)

type transportHandler struct {
	Service core.TransportService
}

func (h *transportHandler) list(c *gin.Context) {
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

func (h *transportHandler) checkIn(c *gin.Context) {
	var transport core.Transport
	err := c.ShouldBindJSON(&transport)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	location, err := h.Service.CheckIn(c, &transport)
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

func (h *transportHandler) checkOut(c *gin.Context) {
	var transport core.Transport
	err := c.ShouldBindJSON(&transport)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorRes{
			Code:    "invalid_input",
			Message: err.Error(),
		})
		return
	}

	location, err := h.Service.CheckOut(c, &transport)
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
