package service

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"my-streaming-service/dao"
	"my-streaming-service/modles"
)

// LoginService 登录服务
// @params userName string,password string
// @returns error
func LoginService(userName string, password string) (bool, error) {
	passwordEncrypt, err := dao.LoginUser(userName)
	if err != nil {
		return false, err
	}
	ok := ComparePassword(password, passwordEncrypt)
	if !ok {
		return false, err
	} else {
		return true, nil
	}
}

//注册服务

func RegisterService(info *modles.UserInfo) error {
	EncryptPassword := Encrypt(info.Password)
	// 将加密后的密码存入info
	info.Password = EncryptPassword
	err := dao.RegisterUser(info.UserName, info.PhoneNumber, info.Email, EncryptPassword)
	if err != nil {
		return err
	}
	return nil
}

// Encrypt 加密函数
func Encrypt(password string) string {
	EncryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)

		return ""
	}
	return string(EncryptPassword)
}

// ComparePassword 比较密码
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
