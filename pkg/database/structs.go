package database

import "gorm.io/gorm"

type Leech struct {
	gorm.Model
	SubjectId int `gorm:"uniqueIndex"`
	Type      string
	Level     int
	Url       string
	Value     string
	Meaning   string
	Reading   *string
}
