package db

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

const (
	Db_name      = "todowork"
	Db_user_name = "root"
	Db_password  = "2002"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := Db_user_name + ":" + Db_password + "@tcp(localhost:3306)/" + Db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
