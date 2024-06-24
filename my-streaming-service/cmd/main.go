package main

import (
	"fmt"
	"my-streaming-service/dbs"
)

func main() {
	dbs.InitDB()
	fmt.Println("Hello, World!")

}
