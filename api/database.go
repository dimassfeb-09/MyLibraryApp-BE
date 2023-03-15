package api

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/mylibraryapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println(err)
	}

	return db.Session(&gorm.Session{SkipDefaultTransaction: true})
}
