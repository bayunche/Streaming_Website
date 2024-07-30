package apis

import (
	"github.com/gin-gonic/gin"
	"my-streaming-service/internal/models"
	"my-streaming-service/service"
)

// NewLiveRoom 新增直播间
func NewLiveRoom(c *gin.Context) {
	userId, _ := c.Get("userId")

	err := service.CreateLiveRoomService(c.PostForm("roomName"), userId.(string))
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success"})
}

// GetLiveRoomList 获取直播间列表
func GetLiveRoomList(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	rooms, err := service.GetLiveRoomListService(page, size)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success", "data": rooms})
}

// GetLiveRoom 查询直播间
func GetLiveRoom(c *gin.Context) {
	roomId := c.Query("roomId")
	roomName := c.Query("roomName")
	room, err := service.QueryLiveRoomService(roomName, roomId)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success", "data": room})
}

// UpdateLiveRoom 更新直播间信息
func UpdateLiveRoom(c *gin.Context) {
	roomId := c.PostForm("roomId")
	userId, _ := c.Get("userId")
	roomName := c.PostForm("roomName")
	description := c.PostForm("description")
	roomUrl := c.PostForm("roomUrl")
	roomInfo := models.LiveRoom{
		RoomName:    roomName,
		Description: description,
		RoomUrl:     roomUrl,
	}
	err := service.UpdateLiveRoomService(roomId, userId.(string), roomInfo)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success"})

}

// DeleteLiveRoom 删除直播间
func DeleteLiveRoom(c *gin.Context) {
	roomId := c.PostForm("roomId")
	userId, _ := c.Get("userId")
	ok, err := service.DeleteLiveRoomService(roomId, userId.(string))
	if ok == false && err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success"})
}

// 推流至直播间
func PushLiveRoom(c *gin.Context) {
	roomId := c.PostForm("roomId")
	userId, _ := c.Get("userId")
	err := service.PushLiveRoomService(roomId, userId.(string))
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success"})
}
