package apis

import (
	"github.com/gin-gonic/gin"
	"my-streaming-service/modles"
	"my-streaming-service/service"
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

		c.JSON(200, gin.H{"message": "登录成功"})
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
