package main

import (
	"github.com/themaintife/feelings/initializers"
	"github.com/themaintife/feelings/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// initializers.DB.Migrator().DropTable(
	// 	models.Feelings{},
	// )
	initializers.DB.AutoMigrate(&models.Feelings{})
}