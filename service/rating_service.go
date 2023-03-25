package service

import (
	"context"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type RatingService interface {
	AddRating(ctx context.Context, r *request.Rating) (isSuccess bool, msg string, err error)
	UpdateRating(ctx context.Context, r *request.Rating) (isSuccess bool, msg string, err error)
	DeleteRating(ctx context.Context, ID int) (isSuccess bool, msg string, err error)
	GetRatingByBookID(ctx context.Context, bookID int) (rating *response.RatingByBook, msg string, err error)
	GetRatingByID(ctx context.Context, ID int) (rating *response.Rating, msg string, err error)
}

type RatingServiceImplementation struct {
	DB               *gorm.DB
	RatingRepository repository.RatingRepository
	BookRepository   repository.BookRepository
	UserRepository   repository.UserRepository
}

func NewRatingServiceImplementation(DB *gorm.DB, microRepository repository.MicroRepository) RatingService {
	return &RatingServiceImplementation{
		DB:               DB,
		RatingRepository: microRepository.Rating(),
		BookRepository:   microRepository.Book(),
		UserRepository:   microRepository.User(),
	}
}

func (rs *RatingServiceImplementation) AddRating(ctx context.Context, r *request.Rating) (bool, string, error) {

	_, _, err := rs.BookRepository.GetBookByID(ctx, rs.DB, r.BookID)
	if err == gorm.ErrRecordNotFound {
		return false, "Buku tidak ditemukan.", err
	}

	_, _, err = rs.UserRepository.GetUserByID(ctx, rs.DB, r.UserID)
	if err == gorm.ErrRecordNotFound {
		return false, "User tidak ditemukan.", err
	}

	rating := &domain.Rating{
		Rating: r.Rating,
		BookID: r.BookID,
		UserID: r.UserID,
	}

	_, msg, err := rs.RatingRepository.AddRating(ctx, rs.DB, rating)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (rs *RatingServiceImplementation) UpdateRating(ctx context.Context, r *request.Rating) (bool, string, error) {
	_, _, err := rs.BookRepository.GetBookByID(ctx, rs.DB, r.BookID)
	if err == gorm.ErrRecordNotFound {
		return false, "Buku tidak ditemukan.", err
	}

	_, _, err = rs.UserRepository.GetUserByID(ctx, rs.DB, r.UserID)
	if err == gorm.ErrRecordNotFound {
		return false, "User tidak ditemukan.", err
	}

	_, _, err = rs.GetRatingByID(ctx, r.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Rating tidak ditemukan.", err
	}

	rating := &domain.Rating{
		ID:     r.ID,
		Rating: r.Rating,
		BookID: r.BookID,
		UserID: r.UserID,
	}

	_, msg, err := rs.RatingRepository.UpdateRating(ctx, rs.DB, rating)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (rs *RatingServiceImplementation) DeleteRating(ctx context.Context, ID int) (bool, string, error) {
	_, _, err := rs.RatingRepository.GetRatingByID(ctx, rs.DB, ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Rating tidak ditemukan.", err
	}

	_, msg, err := rs.RatingRepository.DeleteRating(ctx, rs.DB, ID)
	if err != nil {
		return false, msg, err
	}

	return true, msg, nil
}

func (rs *RatingServiceImplementation) GetRatingByBookID(ctx context.Context, bookID int) (*response.RatingByBook, string, error) {
	_, _, err := rs.BookRepository.GetBookByID(ctx, rs.DB, bookID)
	if err == gorm.ErrRecordNotFound {
		return nil, "Rating tidak ditemukan.", err
	}

	result, msg, err := rs.RatingRepository.GetRatingByBookID(ctx, rs.DB, bookID)
	if err != nil {
		return nil, msg, err
	}

	rating := &response.RatingByBook{
		ID:     result.ID,
		Rating: result.Rating,
		BookID: result.BookID,
	}

	return rating, msg, nil
}

func (rs *RatingServiceImplementation) GetRatingByID(ctx context.Context, ID int) (*response.Rating, string, error) {
	result, msg, err := rs.RatingRepository.GetRatingByID(ctx, rs.DB, ID)
	if err == gorm.ErrRecordNotFound {
		return nil, "Rating tidak ditemukan.", err
	}

	rating := &response.Rating{
		ID:     result.ID,
		Rating: result.Rating,
		BookID: result.BookID,
		UserID: result.UserID,
	}

	return rating, msg, nil
}
