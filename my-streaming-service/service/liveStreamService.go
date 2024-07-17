package service

import (
	"github.com/kjk/betterguid"
	"my-streaming-service/dao"
	"my-streaming-service/internal/models"
	//引入joy5
)

// CreateLiveRoomService 创建直播间
func CreateLiveRoomService(roomName string, userID string) {
	roomId := betterguid.New()
	dao.CreateLiveRoom(roomName, userID, roomId)

}

// QueryLiveRoomService 查询直播间
func QueryLiveRoomService(roomName string, roomId string) *models.LiveRoom {

	room, err := dao.QueryLiveRoom(roomName, roomId)
	if err != nil {
		return nil
	}
	return room
}

// DeleteLiveRoomService 删除直播间
func DeleteLiveRoomService(roomId string) bool { return dao.DeleteLiveRoom(roomId) }
