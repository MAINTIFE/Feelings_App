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
	r.GET("/notes", controllers.AllNotes)
	r.GET("/notes/:id", controllers.SingleNote)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)
	
	
	r.Run()
}
