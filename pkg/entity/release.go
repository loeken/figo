package entity

import (
	"gorm.io/gorm"
)

// Release type
type Release struct {
	gorm.Model
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	Label      string `json:"label"`
	Attachment string `json:"attachment"`
}
