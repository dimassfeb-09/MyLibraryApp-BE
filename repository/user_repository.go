package repository

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx *gorm.DB, user *domain.User) (isSuccess bool, err error)
	UpdateUser(ctx context.Context, tx *gorm.DB, user *domain.User) (isSuccess bool, err error)
	GetUserByID(ctx context.Context, db *gorm.DB, ID int) (user *domain.User, err error)
	GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (user *domain.User, err error)
	GetUserByNPM(ctx context.Context, db *gorm.DB, NPM string) (user *domain.User, err error)
	DeleteUser(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, err error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepositoryImplementation() UserRepository {
	return &UserRepositoryImplementation{}
}

func (u *UserRepositoryImplementation) AddUser(ctx context.Context, tx *gorm.DB, user *domain.User) (bool, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Create(&user).Error; err != nil {
			return err
		} else {
			return nil
		}
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserRepositoryImplementation) UpdateUser(ctx context.Context, tx *gorm.DB, user *domain.User) (bool, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Updates(&user).Where("id = ?", &user.ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserRepositoryImplementation) GetUserByID(ctx context.Context, db *gorm.DB, ID int) (*domain.User, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func (u *UserRepositoryImplementation) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*domain.User, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImplementation) GetUserByNPM(ctx context.Context, db *gorm.DB, npm string) (*domain.User, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("npm = ?", npm).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImplementation) DeleteUser(ctx context.Context, tx *gorm.DB, ID int) (bool, error) {
	var user *domain.User
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Where("id = ?").Delete(&user).Error; err != nil {
			return err
		} else {
			return nil
		}
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
