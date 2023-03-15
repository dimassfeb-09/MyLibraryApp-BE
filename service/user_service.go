package service

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(ctx context.Context, user *request.AddUser) (isSuccess bool, msg string, err error)
	UpdateUser(ctx context.Context, user *domain.User) (isSuccess bool, msg string, err error)
	DeleteUser(ctx context.Context, ID int) (isSuccess bool, msg string, err error)
	GetUserByID(ctx context.Context, ID int) (user *domain.User, msg string, err error)
	GetUserByNPM(ctx context.Context, NPM string) (user *domain.User, msg string, err error)
	GetUserByEmail(ctx context.Context, NPM string) (user *domain.User, msg string, err error)
}

type UserServiceImplementation struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func NewUserServiceImplementation(DB *gorm.DB, userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{DB: DB, UserRepository: userRepository}
}

func (u *UserServiceImplementation) AddUser(ctx context.Context, r *request.AddUser) (bool, string, error) {

	userByNPM, _, _ := u.GetUserByNPM(ctx, r.NPM)
	if userByNPM != nil {
		fmt.Println(userByNPM)
		return false, "NPM telah digunakan.", gorm.ErrRegistered
	}

	userByEmail, _, _ := u.GetUserByEmail(ctx, r.Email)
	if userByEmail != nil {
		return false, "Email telah digunakan.", gorm.ErrRegistered
	}

	user := &domain.User{
		ID:       r.ID,
		Name:     r.Name,
		NPM:      r.NPM,
		Email:    r.Email,
		Password: r.Password,
		IsGoogle: r.IsGoogle,
	}

	return u.UserRepository.AddUser(ctx, u.DB, user)
}

func (u *UserServiceImplementation) UpdateUser(ctx context.Context, user *domain.User) (bool, string, error) {

	_, _, err := u.GetUserByID(ctx, user.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "User ID tidak ditemukan.", gorm.ErrRecordNotFound
	}

	userByNPM, _, _ := u.GetUserByNPM(ctx, user.NPM)
	if userByNPM != nil {
		if userByNPM.ID != user.ID {
			return false, "NPM telah digunakan.", gorm.ErrRegistered
		}
	}

	userByEmail, _, _ := u.GetUserByEmail(ctx, user.Email)
	if userByEmail != nil {
		if userByEmail.ID != user.ID {
			return false, "Email telah digunakan.", gorm.ErrRegistered
		}
	}

	return u.UserRepository.UpdateUser(ctx, u.DB, user)
}

func (u *UserServiceImplementation) DeleteUser(ctx context.Context, ID int) (bool, string, error) {
	_, _, err := u.GetUserByID(ctx, ID)
	if err == gorm.ErrRecordNotFound {
		return false, "User ID tidak ditemukan.", gorm.ErrRecordNotFound
	}

	return u.UserRepository.DeleteUser(ctx, u.DB, ID)
}

func (u *UserServiceImplementation) GetUserByID(ctx context.Context, ID int) (*domain.User, string, error) {

	userByID, msg, err := u.UserRepository.GetUserByID(ctx, u.DB, ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "User ID tidak ditemukan.", gorm.ErrRecordNotFound
		}
		return nil, msg, err
	}

	return userByID, "User ID ditemukan.", nil
}

func (u *UserServiceImplementation) GetUserByNPM(ctx context.Context, NPM string) (*domain.User, string, error) {
	user, msg, err := u.UserRepository.GetUserByNPM(ctx, u.DB, NPM)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "User NPM tidak ditemukan.", gorm.ErrRecordNotFound
		}
		return nil, msg, err
	}

	return user, "User NPM ditemukan.", nil
}

func (u *UserServiceImplementation) GetUserByEmail(ctx context.Context, email string) (*domain.User, string, error) {
	user, msg, err := u.UserRepository.GetUserByEmail(ctx, u.DB, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "User Email tidak ditemukan.", gorm.ErrRecordNotFound
		}
		return nil, msg, err
	}
	return user, "User Email ditemukan.", nil
}
