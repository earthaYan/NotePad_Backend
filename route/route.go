package route

import (
	"NotePad/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/user/register", api.RegisterUserHandler)
	}
	return r
}
