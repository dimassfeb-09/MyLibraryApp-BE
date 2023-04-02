package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type GenreService interface {
	AddGenre(ctx context.Context, r *request.Genre) (isSuccess bool, msg string, err error)
	UpdateGenre(ctx context.Context, r *request.Genre) (isSuccess bool, msg string, err error)
	DeleteGenre(ctx context.Context, ID int) (isSuccess bool, msg string, err error)
	GetGenreByID(ctx context.Context, ID int) (genre *response.Genre, msg string, err error)
	GetGenreByCategoryID(ctx context.Context, ID int) (genre []*response.Genre, msg string, err error)
	GetGenres(ctx context.Context) (genre []*response.Genre, msg string, err error)
}

type GenreServiceImplementation struct {
	DB              *gorm.DB
	GenreRepository repository.GenreRepository
	MicroRepository repository.MicroRepository
}

func NewGenreServiceImplementation(DB *gorm.DB, M repository.MicroRepository) GenreService {
	return &GenreServiceImplementation{DB: DB, GenreRepository: M.Genre(), MicroRepository: M}
}

func (g *GenreServiceImplementation) AddGenre(ctx context.Context, r *request.Genre) (bool, string, error) {

	genre := &domain.Genre{
		Name:       r.Name,
		CategoryID: r.CategoryID,
	}
	_, msg, err := g.GenreRepository.AddGenre(ctx, g.DB, genre)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (g *GenreServiceImplementation) UpdateGenre(ctx context.Context, r *request.Genre) (bool, string, error) {

	_, msg, err := g.GetGenreByID(ctx, r.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Genre ID tidak ditemukan.", err
	}

	genre := &domain.Genre{
		ID:         r.ID,
		Name:       r.Name,
		CategoryID: r.CategoryID,
	}

	fmt.Println("Exec here")

	_, msg, err = g.GenreRepository.UpdateGenre(ctx, g.DB, genre)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (g *GenreServiceImplementation) DeleteGenre(ctx context.Context, ID int) (bool, string, error) {
	_, msg, err := g.GetGenreByID(ctx, ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Genre ID tidak ditemukan.", err
	}

	bookByGenre, _, _ := g.MicroRepository.Book().GetBooksByGenreID(ctx, g.DB, ID)
	if len(bookByGenre) >= 1 {
		return false, "Tidatk dapat menghapus genre yang berelasi dengan buku.", errors.New("Tidak dapat menghapus genre.")
	}

	_, msg, err = g.GenreRepository.DeleteGenre(ctx, g.DB, ID)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (g *GenreServiceImplementation) GetGenreByID(ctx context.Context, ID int) (*response.Genre, string, error) {
	result, msg, err := g.GenreRepository.GetGenreByID(ctx, g.DB, ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "Genre ID tidak ditemukan.", err
		}
		return nil, msg, err
	}

	genre := &response.Genre{
		ID:         result.ID,
		Name:       result.Name,
		CategoryID: result.CategoryID,
	}

	return genre, "Berhasil get Genre.", nil
}

func (g *GenreServiceImplementation) GetGenreByCategoryID(ctx context.Context, ID int) (genres []*response.Genre, msg string, err error) {
	results, msg, err := g.GenreRepository.GetGenreByCategoryID(ctx, g.DB, ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "Genre Category ID tidak ditemukan.", err
		}
		return nil, msg, err
	}

	for _, result := range results {
		genre := &response.Genre{
			ID:         result.ID,
			Name:       result.Name,
			CategoryID: result.CategoryID,
		}
		genres = append(genres, genre)
	}

	return genres, "Berhasil get Genre.", nil
}

func (g *GenreServiceImplementation) GetGenres(ctx context.Context) (genres []*response.Genre, msg string, err error) {
	results, msg, err := g.GenreRepository.GetGenres(ctx, g.DB)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "Genre ID tidak ditemukan.", err
		}
		return nil, msg, err
	}

	for _, result := range results {
		genre := &response.Genre{
			ID:         result.ID,
			Name:       result.Name,
			CategoryID: result.CategoryID,
		}
		genres = append(genres, genre)
	}

	return genres, "Berhasil get genres", nil
}
