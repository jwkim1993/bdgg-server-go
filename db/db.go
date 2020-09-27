package db

import (
	"github.com/jwkim1993/bdgg-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	if db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Member{}, &models.Group{}, &models.Match{}, &models.GroupMember{}, &models.MemberMatch{})
	if err != nil {
		return err
	}

	return nil
}

func DBManager() *gorm.DB {
	return db
}
