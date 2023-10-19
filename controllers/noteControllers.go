package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/themaintife/feelings/initializers"
	"github.com/themaintife/feelings/models"
)

func CreateNotes(c *gin.Context) {
	//get data off req body
	var body struct {
		Title string
		Body  string
		// Date     time.Time `gorm:"type:TIMESTAMP"`
		ImageUrl string
	}
	c.Bind(&body)

	//create a note
	note := models.Feelings{Title: body.Title, Body: body.Body, Date: time.Now().UTC(), ImageUrl: body.ImageUrl}
	result := initializers.DB.Create(&note)

	if result.Error != nil {
		log.Println(result.Error.Error())
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": note,
	})
}

// "The First Record of My Feelings"
