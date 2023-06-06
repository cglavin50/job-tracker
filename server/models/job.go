package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string
	Company     string
	AppliedDate string // user passed in yyyy/mm/dd, could look to use time package later but no date interface
	Status      string
	Interest    string
	Link        string
	CompanyPage string
	Notes       string
}

// https://gorm.io/docs/models.html details field tags
// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }
// for the gorm.Model key
