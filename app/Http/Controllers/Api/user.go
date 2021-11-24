package Api

import (
	"gebi/app/Http/Serializer"
	"gebi/app/Http/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

// @Tags user
// @Description user register
// @Param user_info body Services.UsersRegisterService true "user info"
// @Success 200 {object} Serializer.Response
// @Router /user/register [post]
func (a User) Register(c *gin.Context) {
	var service Services.UsersRegisterService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(service.Register()))
}

// @Tags user
// @Description user login
// @Param user_info body Services.UsersLoginService true "user info"
// @Success 200 {object} Serializer.Response
// @Router /user/login [post]
func (a User) Login(c *gin.Context) {
	var service Services.UsersLoginService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	c.JSON(http.StatusOK, SuccessResponse(service.Login()))
}

// @Tags user
// @Description user me
// @Param Authorization header string true "Authorization"
// @Success 200 {object} Serializer.Response
// @Router /user/me [get]
func (a User) UserMe(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse(Serializer.BuildUser(CurrentUser(c))))
}
