package repository

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type RatingRepository interface {
	AddRating(ctx context.Context, tx *gorm.DB, rating *domain.Rating) (isSuccess bool, msg string, err error)
	UpdateRating(ctx context.Context, tx *gorm.DB, rating *domain.Rating) (isSuccess bool, msg string, err error)
	DeleteRating(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
	GetRatingByBookID(ctx context.Context, db *gorm.DB, bookID int) (rating *domain.Rating, msg string, err error)
	GetRatingByID(ctx context.Context, db *gorm.DB, ID int) (rating *domain.Rating, msg string, err error)
}

type RatingRepositoryImplementation struct {
}

func NewRatingRepositoryImplementation() RatingRepository {
	return &RatingRepositoryImplementation{}
}

func (r *RatingRepositoryImplementation) AddRating(ctx context.Context, tx *gorm.DB, rating *domain.Rating) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("rating").Create(&rating).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal tambah rating.", err
	}

	return true, "Berhasil tambah rating.", nil
}

func (r *RatingRepositoryImplementation) UpdateRating(ctx context.Context, tx *gorm.DB, rating *domain.Rating) (bool, string, error) {
	err := tx.WithContext(ctx).Where("id = ?", rating.ID).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("rating").Updates(&rating).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update rating.", err
	}

	return true, "Berhasil update rating.", nil
}

func (r *RatingRepositoryImplementation) DeleteRating(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	err := tx.WithContext(ctx).Where("user_id = ?", ID).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("rating").Where("id = ?", ID).Delete(ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update rating.", err
	}

	return true, "Berhasil update rating.", nil
}

func (r *RatingRepositoryImplementation) GetRatingByBookID(ctx context.Context, db *gorm.DB, bookID int) (*domain.Rating, string, error) {
	var rating *domain.Rating
	if err := db.WithContext(ctx).Table("rating").Select("AVG(rating) as rating").Where("book_id = ?", bookID).Find(&rating).Error; err != nil {
		return nil, "Gagal get rating by buku", err
	} else {
		return rating, "Berhasil get rating by buku", nil
	}
}

func (r *RatingRepositoryImplementation) GetRatingByID(ctx context.Context, db *gorm.DB, ID int) (*domain.Rating, string, error) {
	var rating *domain.Rating
	if err := db.WithContext(ctx).Table("rating").Where("id = ?", ID).First(&rating).Error; err != nil {
		return nil, "Gagal get rating by id.", err
	} else {
		return rating, "Berhasil get rating by id.", nil
	}
}
