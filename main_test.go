package main

import (
	"fmt"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/api"
	"testing"
)

func TestMains(t *testing.T) {
	db, err := api.DBConn()
	if err != nil {
		fmt.Println(err)
	}

	//genre := &domain.Genre{
	//	ID:         1,
	//	Name:       "",
	//	CategoryID: 0,
	//}
	//
	//genreRepository := repository.NewGenreRepositoryImplementation()
	//genreService := service.NewGenreServiceImplementation(db, genreRepository)
	//result, msg, err := genreService.DeleteGenre(context.Background(), 1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(msg)
	//fmt.Println(result)

}
