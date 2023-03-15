package repository

import (
	"context"
	"fmt"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx *gorm.DB, user *domain.User) (isSuccess bool, msg string, err error)
	UpdateUser(ctx context.Context, tx *gorm.DB, user *domain.User) (isSuccess bool, msg string, err error)
	GetUserByID(ctx context.Context, db *gorm.DB, ID int) (user *domain.User, msg string, err error)
	GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (user *domain.User, msg string, err error)
	GetUserByNPM(ctx context.Context, db *gorm.DB, NPM string) (user *domain.User, msg string, err error)
	DeleteUser(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepositoryImplementation() UserRepository {
	return &UserRepositoryImplementation{}
}

func (u *UserRepositoryImplementation) AddUser(ctx context.Context, tx *gorm.DB, user *domain.User) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Create(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, "Gagal tambah user.", err
	}

	return true, "Berhasil tambah user.", nil
}

func (u *UserRepositoryImplementation) UpdateUser(ctx context.Context, tx *gorm.DB, user *domain.User) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Updates(&user).Where("id = ?", &user.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, "Gagal update user.", err
	}

	return true, "Berhasil update user.", nil
}

func (u *UserRepositoryImplementation) GetUserByID(ctx context.Context, db *gorm.DB, ID int) (*domain.User, string, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, "Gagal Get User by ID.", err
	}
	fmt.Println(user)
	return user, "Berhasil Get User by ID.", nil
}

func (u *UserRepositoryImplementation) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*domain.User, string, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "Gagal Get User by Email.", err
	}
	return user, "Berhasil Get User by Email.", nil
}

func (u *UserRepositoryImplementation) GetUserByNPM(ctx context.Context, db *gorm.DB, npm string) (*domain.User, string, error) {
	var user *domain.User
	if err := db.WithContext(ctx).Table("user").Where("npm = ?", npm).First(&user).Error; err != nil {
		return nil, "Gagal Get User by NPM.", err
	}
	return user, "Berhasil Get User by NPM.", nil
}

func (u *UserRepositoryImplementation) DeleteUser(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	var user *domain.User
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("user").Where("id = ?").Delete(&user).Error; err != nil {
			return err
		} else {
			return nil
		}
	})

	if err != nil {
		return false, "Gagal hapus data user.", err
	}

	return true, "Berhasil hapus data user,", nil
}
