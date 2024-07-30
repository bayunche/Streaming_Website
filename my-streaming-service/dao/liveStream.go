package dao

import (
	"errors"
	"log"
	"my-streaming-service/internal/models"
	"strconv"
)

// CreateLiveRoom 创建直播间
func CreateLiveRoom(roomName string, userID string, roomId string) error {
	// 检查RoomId是否已经存在
	var existingRoom models.LiveRoom
	result := DB.Where("RoomId = ?", roomId).First(&existingRoom)
	if result.Error == nil {
		return errors.New("RoomId already exists")
	}

	LiveRoom := models.LiveRoom{RoomName: roomName, UserID: userID, RoomId: roomId}
	result = DB.Create(&LiveRoom)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// QueryLiveRoom 查询直播间
// QueryLiveRoom 根据roomName和roomId查询直播间
func QueryLiveRoom(roomName string, roomId string) (*models.LiveRoom, error) {
	// 初始化LiveRoom
	LiveRoom := models.LiveRoom{}
	// 判断使用roomName查询还是roomId查询
	if roomName != "" {
		// 使用roomName查询
		result := DB.Where("RoomName LIKE ?", "%"+roomName+"%").First(&LiveRoom)
		if result.Error != nil {
			// 查询出错，返回错误
			return nil, result.Error
		}
		// 查询成功，返回LiveRoom
		return &LiveRoom, nil
	} else if roomId != "" {
		// 使用roomId查询
		result := DB.Where("roomId = ?", roomId).First(&LiveRoom)
		if result.Error != nil {
			// 查询出错，返回错误
			return nil, result.Error
		}
		// 查询成功，返回LiveRoom
		return &LiveRoom, nil
	}

	// roomName和roomId都为空，返回错误
	return nil, errors.New("roomName and roomId are both empty")
}

// DeleteLiveRoom 删除直播间
func DeleteLiveRoom(roomId string, userId string) (bool, error) {
	var LiveRoom models.LiveRoom
	result := DB.Where("room_id = ? And userID = ? ", roomId, userId).Delete(&LiveRoom)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetLiveRoomList 获取直播间列表
func GetLiveRoomList(page string, size string) ([]models.LiveRoom, error) {
	var err error
	var pageInt int
	var sizeInt int

	// 将字符串转换为整数
	pageInt, err = strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	sizeInt, err = strconv.Atoi(size)
	if err != nil {
		return nil, err
	}

	// 声明一个新的变量来存储查询结果
	var liveRooms []models.LiveRoom

	// 分页查找直播间列表
	result := DB.Offset((pageInt - 1) * sizeInt).Limit(sizeInt).Find(&liveRooms)
	if result.Error != nil {
		// 处理错误
		log.Println(result.Error)
		return nil, result.Error
	}

	// 返回查询结果
	return liveRooms, nil
}

// UpdateLiveRoomInfo 更新直播间信息
func UpdateLiveRoomInfo(roomId string, userId string, roomInfo models.LiveRoom) (bool, error) {
	r := DB.Model(&models.LiveRoom{}).Where("room_id = ? AND user_id = ?", roomId, userId).Updates(roomInfo)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}
