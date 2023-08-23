package main

import (
	"github.com/RaymondCode/simple-demo/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//go service.RunMessageServer()

	database.Init()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
