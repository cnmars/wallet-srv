package db

import (
	"wallet-srv/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {

	var err error
	db, err = gorm.Open("mysql", conf.MYSQL_DSN)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	db.SingularTable(true)

	db.LogMode(false)

}

func DB() *gorm.DB {
	return db
}
