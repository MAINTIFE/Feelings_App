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
		"note": note,
	})
}

func AllNotes(c *gin.Context) {
	// Get all the note records
	var notes []models.Feelings
	initializers.DB.Find(&notes)

	//respond with them
	c.JSON(200, gin.H{
		"notes": notes,
	})
}

func SingleNote(c *gin.Context) {
	// Get all the note records
	id := c.Param("id")
	var note models.Feelings
	initializers.DB.First(&note, id)

	//respond with them
	c.JSON(200, gin.H{
		"notes": note,
	})
}

func UpdateNote(c *gin.Context) {
	// Get all the note records
	id := c.Param("id")
	var note models.Feelings
	initializers.DB.First(&note, id)

	//get data off req body
	var body struct {
		Title    string
		Body     string
		ImageUrl string
	}
	c.Bind(&body)

	//update it
	initializers.DB.Model(&note).Updates(models.Feelings{
		Title:    body.Title,
		Body:     body.Body,
		Date:     time.Now().UTC(),
		ImageUrl: body.ImageUrl,
	})

	//respond with them
	c.JSON(200, gin.H{
		"notes": note,
	})
}

func DeleteNote(c *gin.Context) {
	// Get all the note records
	id := c.Param("id")
	initializers.DB.Delete(&models.Feelings{}, id)
	

	//respond with them
	c.Status(200)
}
