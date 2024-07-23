package main

import (
	"fmt"
	"my-streaming-service/config"
	"my-streaming-service/dbs"

	"github.com/gin-gonic/gin"
)

func main() {

	dbs.InitDB()
	r := gin.Default()
	r.BasePath()

	config.Initproject()
	fmt.Println("Hello, World!")

}
