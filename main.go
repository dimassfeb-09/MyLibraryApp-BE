package main

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/api"
	"log"
)

func main() {
	db, err := api.DBConn()
	if err != nil {
		log.Println(err)
	}

	r := api.GinRoute(db)
	r.Run(":8080")
}
