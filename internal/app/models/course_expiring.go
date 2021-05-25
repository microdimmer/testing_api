package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Course_expiring struct {
	gorm.Model
	Expiring  time.Time `json:"expiring"`
	Link      string    `json:"link"`
	Empl_name string    `json:"empl_name"`
	Empl_dep  string    `json:"empl_dep"`
	Passing   byte      `json:"passing"`
	CourseID  int       `json:"course_id"`
}

func (course_exp *Course_expiring) Validate() bool {

	if course_exp.Link == "" {
		return false
	}
	count := 0
	timestamp := time.Now()
	GetDB().Model(&Course_expiring{}).Where("link = ? AND expiring >= ?", course_exp.Link, timestamp).Count(&count)

	return count != 0
}

func (course_exp *Course_expiring) Start() bool {
	c := &Course_expiring{}
	GetDB().Model(&Course_expiring{}).Where("Link = ? AND Passing = 0", course_exp.Link).First(&c)

	if c.ID == 0 || course_exp.Empl_name == "" || course_exp.Empl_dep == "" {
		return false
	}

	course_exp.Expiring = time.Now().Add(1 * time.Hour)
	course_exp.Passing = 1
	course_exp.ID = c.ID
	course_exp.CourseID = c.CourseID
	status := GetDB().Save(course_exp)
	if status.Error != nil {
		logger.Error(status.Error)
		return false
	}
	return true
}

func (course_exp *Course_expiring) Process() bool {
	c := &Course_expiring{}
	GetDB().Model(&Course_expiring{}).Where("Link = ? AND Passing <> 0", course_exp.Link).First(&c)

	if c.ID == 0 || c.Passing >= course_exp.Passing {
		return false
	}

	c.Passing = course_exp.Passing

	timestamp := time.Now()
	if c.Passing > 10 || c.Expiring.Before(timestamp) {
		return false
	}

	if c.Passing == 10 {
		c.Expiring = time.Now()
		status := GetDB().Save(c)
		if status.Error != nil {
			logger.Error(status.Error)
			return false
		}
		return true
	}
	status := GetDB().Save(c)
	if status.Error != nil {
		logger.Error(status.Error)
		return false
	}

	return true
}

func (course_exp *Course_expiring) FindCourse() *Course {
	course := &Course{}
	c_exp := &Course_expiring{}
	GetDB().Model(&Course_expiring{}).Where("link = ?", course_exp.Link).First(c_exp)
	GetDB().Model(&Course{}).Where("ID = ?", c_exp.CourseID).First(course)
	// GetDB().Model(&Course_expiring{}).Select("courses.*").Joins("left join courses on courses.id = course_expirings.course_id").First(course)

	if course.Name == "" {
		return nil
	}

	return course
}

func (course_exp *Course_expiring) Create() bool {
	status := GetDB().Create(course_exp)
	if status.Error != nil {
		logger.Error(status.Error)
		return false
	}
	return true
}
