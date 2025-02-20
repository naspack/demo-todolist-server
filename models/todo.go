package models

import (
	"time"
)

type Todo struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	Title       string     `gorm:"size:100" json:"title"`
	Description string     `gorm:"size:500" json:"description"`
	Completed   bool       `json:"completed" gorm:"default:false"`
}
