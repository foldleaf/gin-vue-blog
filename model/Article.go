package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category    Category
	Title       string `gorm:"type:varchar(20);not null" json:"title"`
	CategoryId  int    `gorm:"type:int;not null" json:"category_id"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Content     string `gorm:"type:text" json:"content"`
	Img         string `gorm:"type:varchar(200)" json:"img"`
}
