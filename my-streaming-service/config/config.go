package config

import (
	"my-streaming-service/routers"
	"my-streaming-service/service"
)

type Config struct {
	DbUrl string
}

var AppConfig Config

//加载配置

func loadConfig() {

}

// Initproject 初始化项目
func Initproject() {
	AppConfig.DbUrl = "127.0.0.1:3306"
	//初始化数据库
	//dbs.InitDB()
	//初始化直播manager
	service.NewManager()
	//启动gin
	err := routers.Routers().Run(":8080")
	if err != nil {
		return
	}

}
