package api

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConn() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mylibraryapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println(err)
	}

	db.Session(&gorm.Session{SkipDefaultTransaction: true})

	var user User
	db.Table("user").Find(&user).Where("id = ?", 20).Find(&user)
	//tx.Scan(&user.ID, &user.Name, &user.NPM, &user.Email, &user.Password, &user.IsGoogle)
	fmt.Println(user)

}
