package Api

import (
	"gebi/app/Http/Serializer"
	"gebi/app/Http/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Address struct {
}

func (a Address) List(c *gin.Context) {
	var service Services.AddressListService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(Serializer.BuildAddress(service.List())))
}

func (a *Address) Update(c *gin.Context) {
	var service Services.AddressUpdateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(service.Update()))
}

func (a *Address) Del(c *gin.Context) {
	var service Services.AddressDelService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(service.Del()))
}

func (a Address) Add(c *gin.Context) {
	var service Services.AddressAddService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(service.Add()))
}
