package repository

import (
	"context"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type GenreRepository interface {
	AddGenre(ctx context.Context, tx *gorm.DB, genre *domain.Genre) (isSuccess bool, msg string, err error)
	UpdateGenre(ctx context.Context, tx *gorm.DB, genre *domain.Genre) (isSuccess bool, msg string, err error)
	DeleteGenre(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
	GetGenreByID(ctx context.Context, db *gorm.DB, ID int) (genre *domain.Genre, msg string, err error)
	GetGenreByCategoryID(ctx context.Context, db *gorm.DB, ID int) (genre []*domain.Genre, msg string, err error)
	GetGenres(ctx context.Context, db *gorm.DB) (genres []*domain.Genre, msg string, err error)
}

type GenreRepositoryImplementation struct {
}

func NewGenreRepositoryImplementation() GenreRepository {
	return &GenreRepositoryImplementation{}
}

func (g *GenreRepositoryImplementation) AddGenre(ctx context.Context, tx *gorm.DB, genre *domain.Genre) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("genre").Create(&genre).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal tambah genre.", err
	}

	return true, "Berhasil tambah genre.", nil
}

func (g *GenreRepositoryImplementation) UpdateGenre(ctx context.Context, tx *gorm.DB, genre *domain.Genre) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("genre").Where("id = ?", genre.ID).Updates(&genre).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update genre.", err
	}

	return true, "Berhasil update genre.", nil
}

func (g *GenreRepositoryImplementation) DeleteGenre(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("genre").Where("id = ?", ID).Delete(&ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal hapus genre.", err
	}

	return true, "Berhasil hapus genre.", nil
}

func (g *GenreRepositoryImplementation) GetGenreByID(ctx context.Context, db *gorm.DB, ID int) (genre *domain.Genre, msg string, err error) {
	if err := db.WithContext(ctx).Table("genre").Where("id = ?", ID).First(&genre).Error; err != nil {
		return nil, "Gagal get genre.", err
	}
	return genre, "Berhasil get genre.", nil
}

func (g *GenreRepositoryImplementation) GetGenreByCategoryID(ctx context.Context, db *gorm.DB, ID int) (genre []*domain.Genre, msg string, err error) {
	if err := db.WithContext(ctx).Table("genre").Where("category_id = ?", ID).Limit(4).Find(&genre).Error; err != nil {
		return nil, "Gagal get genre.", err
	}
	return genre, "Berhasil get genre.", nil
}

func (g *GenreRepositoryImplementation) GetGenres(ctx context.Context, db *gorm.DB) (genres []*domain.Genre, msg string, err error) {
	if err := db.WithContext(ctx).Table("genre").Find(&genres).Error; err != nil {
		return nil, "Gagal get genres.", nil
	}

	return genres, "Berhasil get genres.", err
}
