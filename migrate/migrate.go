package main

import (
	"github.com/aayush/go-crud/inititalizers"
	"github.com/aayush/go-crud/models"
)

func init() {
	inititalizers.LoadEnvVariables()
	inititalizers.ConnectTODB()
}

func main() {
	inititalizers.DB.AutoMigrate(&models.Post{})
}
