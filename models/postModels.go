package models

import "gorm.io/gorm"

type Feelings struct {
	gorm.Model
	Title    string
	Body     string
	Date     uint
	ImageUrl string
}
