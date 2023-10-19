package models

import (
	"time"

	"gorm.io/gorm"
)

type Feelings struct {
	gorm.Model
	Title    string
	Body     string
	Date     time.Time `gorm:"type:TIMESTAMP"`
	ImageUrl string
}
