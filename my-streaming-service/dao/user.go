package dao

import (
	"my-streaming-service/internal/models"

	"github.com/kjk/betterguid"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 新增用户
// @params name:用户名 password:密码  email:邮箱
// @return string：用户id error：错误信息
func AddUser(userName string, password string, email string) (string, error) {

	userID := betterguid.New()
	user := models.User{Username: userName, Password: password, Email: email, userID}

	result := DB.Create(&user)

	if result.Error != nil {

		return "", result.Error

	}

	return userID, nil
}

// 查询用户
// @params id:用户id
// @return *models.User：用户信息 error：错误信息
func GetUser(id string) (*models.User, error) {

	user := &models.User{}
	result := DB.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}

// 删除用户
// @params id:用户id
// @return error：错误信息
func DeleteUser(id string) error {

	result := DB.Delete(&models.User{}, id)
	if result.Error != nil {

		return result.Error
	}
	return result.Error

}

// 更新用户
// @params id:用户id user:用户信息
// @return error：错误信息
func UpdateUser(id string, user *models.User) error {

	result := DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

// 查询所有用户
// 按照创建时间排序
// @params page:页码 pageSize:每页数量
// @return []models.User：用户列表 error：错误信息
func GetAllUser(page int, pageSize int) ([]models.User, error) {

	var users []models.User
	result := DB.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return
	}

	return users, result.Error

}
