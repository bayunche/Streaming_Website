package Util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
) // 导入jwt-go包

// JwtToken 生成JWT Token
func JwtToken(userId string, username string) (string, error) {

	// 创建一个Claims对象，用于存储需要签名的数据
	claims := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 48).Unix(), // 设置过期时间为48小时
	}

	// 使用HS256算法生成JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("HasuneMiku")) // 使用自定义的密钥进行签名
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// JwtParse  解析JWT Token
func JwtParse(tokenString string, r *gin.Context) error {
	// 解析JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回自定义的密钥
		return []byte("secret"), nil
	})

	// 检查解析结果
	if err != nil {
		return err
	}

	// 获取Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	r.Set("userId", claims["userId"])
	r.Set("username", claims["username"])
	return nil
}

// jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		err := JwtParse(tokenString, c)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
