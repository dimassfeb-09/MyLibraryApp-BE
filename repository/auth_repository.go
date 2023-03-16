package repository

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, tx *gorm.DB, user *domain.User) (isSuccess bool, msg string, err error)
	LoginUser(ctx context.Context, tx *gorm.DB, email string) (user *domain.User, err error)
}

type AuthRepositoryImplementation struct {
}

func NewAuthRepositoryImplementation() AuthRepository {
	return &AuthRepositoryImplementation{}
}

func (a *AuthRepositoryImplementation) RegisterUser(ctx context.Context, tx *gorm.DB, user *domain.User) (bool, string, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Table("user").Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, "Gagal daftar user.", err
	}

	return true, "Berhasil daftar user.", nil
}

func (a *AuthRepositoryImplementation) LoginUser(ctx context.Context, tx *gorm.DB, email string) (*domain.User, error) {
	var user *domain.User
	if err := tx.Table("user").Select("id, email, password").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
