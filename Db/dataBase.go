package Db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Mysql(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&parseTime=true",
		username,
		password,
		hostname,
		port,
		dbname)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	mysqlDb, _ := db.DB()
	mysqlDb.SetMaxIdleConns(20)
	mysqlDb.SetMaxOpenConns(100)

	DB = db
	return db, nil
}
