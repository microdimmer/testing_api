package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	// "github.com/google/uuid"
)

var db *gorm.DB
var logger *logrus.Logger

func InitConnection(dsnUrl string, logger *logrus.Logger) error {

	conn, err := gorm.Open("postgres", dsnUrl)
	if err != nil {
		return err
	}

	db = conn
	db.Debug().AutoMigrate(&Course{}, &Course_expiring{})

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func ClearTables() {
	db.Exec("TRUNCATE TABLE courses")
	db.Exec("TRUNCATE TABLE course_expirings")
}
