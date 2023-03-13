package repository

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx gorm.DB, user *domain.User) (*domain.User, error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepositoryImplementation() UserRepository {
	return &UserRepositoryImplementation{}
}

func (u *UserRepositoryImplementation) AddUser(ctx context.Context, tx gorm.DB, user *domain.User) (*domain.User, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Create(&user).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
