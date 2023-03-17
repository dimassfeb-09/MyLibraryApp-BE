package service

import (
	"context"
	"errors"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	AuthRegister(ctx context.Context, register *request.AuthRegister) (isSuccess bool, msg string, err error)
	AuthLogin(ctx context.Context, login *request.AuthLogin) (isSuccess bool, msg string, err error)
}

type AuthServiceImplementation struct {
	DB             *gorm.DB
	AuthRepository repository.AuthRepository
	UserService    UserService
}

func NewAuthServiceImplementation(DB *gorm.DB, authRepository repository.AuthRepository, userService UserService) AuthService {
	return &AuthServiceImplementation{DB: DB, AuthRepository: authRepository, UserService: userService}
}

func (a *AuthServiceImplementation) AuthRegister(ctx context.Context, r *request.AuthRegister) (bool, string, error) {

	_, msg, err := a.UserService.GetUserByEmail(ctx, r.Email)
	if err != gorm.ErrRecordNotFound {
		return false, "Email telah digunakan.", gorm.ErrRegistered
	}

	_, msg, err = a.UserService.GetUserByNPM(ctx, r.NPM)
	if err != gorm.ErrRecordNotFound {
		return false, "NPM telah digunakan.", gorm.ErrRegistered
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, "Internal Server Error", errors.New("Internal Server Error")
	}

	authRegister := &domain.AuthRegister{
		Name:     r.Name,
		NPM:      r.NPM,
		Email:    r.Email,
		Password: string(hashPassword),
		IsGoogle: r.IsGoogle,
	}
	_, msg, err = a.AuthRepository.AuthRegister(ctx, a.DB, authRegister)
	if err != nil {
		return false, msg, err
	}

	return false, msg, err
}

func (a *AuthServiceImplementation) AuthLogin(ctx context.Context, login *request.AuthLogin) (bool, string, error) {

	loginDetail, err := a.AuthRepository.AuthLogin(ctx, a.DB, login.Email)
	if err != nil {
		return false, "Internal Server Error.", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginDetail.Password), []byte(login.Password))
	if err != nil {
		return false, "Password salah, silahkan cek kembali.", err
	}

	return true, "Berhasil Login.", nil
}
