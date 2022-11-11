package main

import (
	"github.com/aayush/go-crud/controllers"
	"github.com/aayush/go-crud/inititalizers"
	"github.com/gin-gonic/gin"
)

func init() {
	inititalizers.LoadEnvVariables()
	inititalizers.ConnectTODB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
