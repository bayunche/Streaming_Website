package routers

import (
	"github.com/gin-gonic/gin"
	"my-streaming-service/apis"
)

func Routers() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/userLogin", apis.Login)
		user.POST("/userRegister", apis.Resign)
		user.GET("/userInfo", apis.QueryUserInfo)
	}
	liveRoom := r.Group("/liveRoom")
	{
		liveRoom.GET("/getLiveRoomList", apis.GetLiveRoomList)
		liveRoom.POST("/creatLiveRoom", apis.NewLiveRoom)
		liveRoom.GET("/getLiveRoomInfo", apis.GetLiveRoom)
		liveRoom.POST("/updateLiveRoom", apis.UpdateLiveRoom)
		liveRoom.POST("/deleteLiveRoom", apis.DeleteLiveRoom)

	}
	return r
}
