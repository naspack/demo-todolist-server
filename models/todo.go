package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"size:100" json:"title"`
	Description string `gorm:"size:500" json:"description"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}