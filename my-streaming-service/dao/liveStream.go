package dao

import (
	"my-streaming-service/internal/models"

	"gorm.io/gorm"
)

// 创建直播间

func CreateLiveRoom(roomName string, userID string, roomId string) error {

	LiveRoom := models.LiveRoom{RoomName: roomName, UserID: userID, RoomId: roomId}
	result := DB.Create(&LiveRoom)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// QueryLiveRoom 查询直播间
func QueryLiveRoom(roomName string, roomId string) (*gorm.DB, error) {
	LiveRoom := models.LiveRoom{}

	// 判断使用roomName查询还是roomId查询
	if roomName != "" {

		result := DB.Where("RoomName Like ?", "%"+roomName+`%`).First(&LiveRoom)
		if result.Error != nil {
			return nil, result.Error
		}
		return result, nil
	} else if roomId != "" {
		result := DB.Where("roomId = ?", roomId).First(&LiveRoom)
		if result.Error != nil {
			return nil, result.Error
		}
		return result, nil
	}

	return nil, nil
}

// 删除直播间

func DeleteLiveRoom(roomId string) error {

	var LiveRoom models.LiveRoom
	result := DB.Where("room_id = ?", roomId).Delete(&LiveRoom)
	return result.Error
}
