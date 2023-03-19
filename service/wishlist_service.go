package service

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type WishlistService interface {
	AddWishlist(ctx context.Context, wishlist *request.Wishlist) (isSuccess bool, msg string, err error)
	UpdateWishlist(ctx context.Context, wishlist *request.Wishlist) (isSuccess bool, msg string, err error)
	DeleteWishlist(ctx context.Context, ID int) (isSuccess bool, msg string, err error)
	GetWishlistByID(ctx context.Context, ID int) (wishlist *response.Wishlist, msg string, err error)
	GetWishlistByUserID(ctx context.Context, userID int) (wishlist []*response.Wishlist, msg string, err error)
	GetWishlists(ctx context.Context) (wishlists []*response.Wishlist, msg string, err error)
}

type WishlistServiceImplementation struct {
	db                 *gorm.DB
	WishlistRepository repository.WishlistRepository
	UserService        UserService
	BookService        BookService
}

func NewWishlistServiceImplementation(db *gorm.DB, wishlistRepository repository.WishlistRepository, userRepository UserService, bookService BookService) WishlistService {
	return &WishlistServiceImplementation{db: db, WishlistRepository: wishlistRepository, UserService: userRepository, BookService: bookService}
}

func (w *WishlistServiceImplementation) AddWishlist(ctx context.Context, r *request.Wishlist) (bool, string, error) {

	_, _, err := w.UserService.GetUserByID(ctx, r.UserID)
	if err == gorm.ErrRecordNotFound {
		return false, "User tidak ditemukan.", err
	}

	_, _, err = w.BookService.GetBookByID(ctx, r.BookID)
	if err == gorm.ErrRecordNotFound {
		return false, "Buku tidak ditemukan.", err
	}

	wishlist := &domain.Wishlist{
		BookID: r.BookID,
		UserID: r.UserID,
	}

	return w.WishlistRepository.AddWishlist(ctx, w.db, wishlist)
}

func (w *WishlistServiceImplementation) UpdateWishlist(ctx context.Context, r *request.Wishlist) (bool, string, error) {

	_, _, err := w.GetWishlistByID(ctx, r.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Wishlist tidak ditemukan", err
	}

	_, _, err = w.UserService.GetUserByID(ctx, r.UserID)
	if err == gorm.ErrRecordNotFound {
		return false, "User tidak ditemukan.", err
	}

	_, _, err = w.BookService.GetBookByID(ctx, r.BookID)
	if err == gorm.ErrRecordNotFound {
		return false, "Buku tidak ditemukan.", err
	}

	wishlist := &domain.Wishlist{
		ID:     r.ID,
		BookID: r.BookID,
		UserID: r.UserID,
	}

	return w.WishlistRepository.UpdateWishlist(ctx, w.db, wishlist)
}

func (w *WishlistServiceImplementation) DeleteWishlist(ctx context.Context, ID int) (isSuccess bool, msg string, err error) {
	_, _, err = w.GetWishlistByID(ctx, ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Wishlist tidak ditemukan.", err
	}

	return w.WishlistRepository.DeleteWishlist(ctx, w.db, ID)
}

func (w *WishlistServiceImplementation) GetWishlistByID(ctx context.Context, ID int) (*response.Wishlist, string, error) {
	result, _, err := w.WishlistRepository.GetWishlistByID(ctx, w.db, ID)
	if err == gorm.ErrRecordNotFound {
		return nil, "Wishlist tidak ditemukan.", err
	}

	wishlist := &response.Wishlist{
		ID:     result.ID,
		BookID: result.BookID,
		UserID: result.UserID,
	}

	return wishlist, "Berhasil get wishlist.", nil
}

func (w *WishlistServiceImplementation) GetWishlistByUserID(ctx context.Context, userID int) ([]*response.Wishlist, string, error) {
	results, msg, err := w.WishlistRepository.GetWishlistByUserID(ctx, w.db, userID)
	if err == gorm.ErrRecordNotFound {
		return nil, msg, err
	}

	var wishlists []*response.Wishlist
	for _, result := range results {
		wishlist := &response.Wishlist{
			ID:     result.ID,
			BookID: result.BookID,
			UserID: result.UserID,
		}
		wishlists = append(wishlists, wishlist)
	}

	return wishlists, "Berhasil get wishlists.", nil
}

func (w *WishlistServiceImplementation) GetWishlists(ctx context.Context) ([]*response.Wishlist, string, error) {
	results, msg, err := w.WishlistRepository.GetWishlists(ctx, w.db)
	if err != nil {
		return nil, msg, err
	}

	var wishlists []*response.Wishlist
	for _, result := range results {
		wishlist := &response.Wishlist{
			ID:     result.ID,
			BookID: result.BookID,
			UserID: result.UserID,
		}
		wishlists = append(wishlists, wishlist)
	}

	return wishlists, "Berhasil get wishlists", nil
}
