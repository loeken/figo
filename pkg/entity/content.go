package entity

import (
	"gorm.io/gorm"
)

// Content type
type Content struct {
	gorm.Model
	Title      string `json:"title"`
	Body     string `json:"body"`
}
