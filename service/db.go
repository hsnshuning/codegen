package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB



func init() {
	var err error
	db, err = gorm.Open(mysql.Open(Cmd.Dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Debug()
}


