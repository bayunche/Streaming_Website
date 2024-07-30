package main

import (
	"fmt"
	"my-streaming-service/config"
)

func main() {
	// 初始化数据库
	//dbs.InitDB()
	config.Initproject()
	fmt.Println("Hello, World!")

}
