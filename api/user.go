package api

import (
	"NotePad/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	var userRegister service.RegisterUserService
	
	if  err := c.ShouldBind(&userRegister);err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
