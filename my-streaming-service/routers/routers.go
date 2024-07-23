package routers

import (
	"github.com/gin-gonic/gin"
	"my-streaming-service/apis"
)

func Routers() *gin.Engine {
	r := gin.Default()
	user := r.Group("/users")
	{
		user.POST("/userLogin", apis.Login)
		user.POST("/userRegister", apis.Resign)
	}

	return r
}
