package main

import (
	"github.com/gin-gonic/gin"
	"github.com/themaintife/feelings/controllers"
	"github.com/themaintife/feelings/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()

	r.POST("/notes", controllers.CreateNotes)

	r.Run()
}
