package models

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Name string `json:"name"`
}

func (course *Course) Create() bool {
	status := GetDB().Create(course)
	if status.Error != nil {
		logger.Error(status.Error)
		return false
	}
	return true
}
