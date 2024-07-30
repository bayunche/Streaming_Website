package apis

import (
	"my-streaming-service/internal/models"
	"my-streaming-service/modles"
	"my-streaming-service/service"

	"github.com/gin-gonic/gin"
)

// Login 登录接口
func Login(c *gin.Context) {
	// 获取用户名和密码
	userName := c.PostForm("username")
	password := c.PostForm("password")
	//验证用户名和密码
	complete, loginService := service.LoginService(userName, password)
	if loginService != nil {
		return
	}
	if complete {

		c.JSON(200, gin.H{"message": "登录成功", "data": nil})
	}

}

// 注册接口

func Resign(c *gin.Context) {

	info := &modles.UserInfo{
		UserName:    c.PostForm("username"),
		Password:    c.PostForm("password"),
		Email:       c.PostForm("email"),
		PhoneNumber: c.PostForm("phoneNumber"),
	}

	//调用注册服务
	err := service.RegisterService(info)
	if err != nil {
		return
	}
}

//查询用户信息

func QueryUserInfo(c *gin.Context) {
	//获取存储在上下文的userid
	userId, _ := c.Get("userID")
	//调用查询服务
	userInfo, err := service.QueryUserService(userId.(string))
	if err != nil {
		c.JSON(200, gin.H{"message": "查询失败", "data": nil})
		return

	}
	c.JSON(200, gin.H{"data": userInfo, "message": "查询成功"})
}

// 修改用户信息

func UpdateUserInfo(c *gin.Context) {
	userInfo := &models.User{}
	//从postform中获取数据
	err := c.ShouldBindJSON(userInfo)
	if err != nil {
		return
	}
	//从上下文中获取userid
	userId, _ := c.Get("userID")

	//调用修改服务
	editErr := service.UpdateUserService(userId.(string), userInfo)
	if editErr != nil {
		return
	}

	c.JSON(200, gin.H{"message": "登录成功", "data": nil})
}
