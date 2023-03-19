package repository

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	AddWishlist(ctx context.Context, tx *gorm.DB, wishlist *domain.Wishlist) (isSuccess bool, msg string, err error)
	UpdateWishlist(ctx context.Context, tx *gorm.DB, wishlist *domain.Wishlist) (isSuccess bool, msg string, err error)
	DeleteWishlist(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
	GetWishlistByID(ctx context.Context, db *gorm.DB, ID int) (wishlist *domain.Wishlist, msg string, err error)
	GetWishlistByUserID(ctx context.Context, db *gorm.DB, userID int) (wishlist []*domain.Wishlist, msg string, err error)
	GetWishlists(ctx context.Context, db *gorm.DB) (wishlists []*domain.Wishlist, msg string, err error)
}

type WishlistRepositoryImplementation struct {
}

func NewWishlistRepositoryImplementation() WishlistRepository {
	return &WishlistRepositoryImplementation{}
}

func (w *WishlistRepositoryImplementation) AddWishlist(ctx context.Context, tx *gorm.DB, wishlist *domain.Wishlist) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("wishlist").Create(&wishlist).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal tambah wishlist.", err
	}

	return true, "Berhasil tambah wishlist.", nil
}

func (w *WishlistRepositoryImplementation) UpdateWishlist(ctx context.Context, tx *gorm.DB, wishlist *domain.Wishlist) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("wishlist").Where("id = ?", wishlist.ID).Updates(&wishlist).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update wishlist.", err
	}

	return true, "Berhasil update wishlist.", nil
}

func (w *WishlistRepositoryImplementation) DeleteWishlist(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("wishlist").Where("id = ?", ID).Delete(ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal hapus wishlist.", err
	}

	return true, "Berhasil hapus wishlist.", nil
}

func (w *WishlistRepositoryImplementation) GetWishlistByID(ctx context.Context, db *gorm.DB, ID int) (wishlist *domain.Wishlist, msg string, err error) {
	if err := db.WithContext(ctx).Table("wishlist").Where("id = ?", ID).First(&wishlist).Error; err != nil {
		return nil, "Gagal get wishlist.", err
	}

	return wishlist, "Berhasil get wishlist.", nil
}

func (w *WishlistRepositoryImplementation) GetWishlistByUserID(ctx context.Context, db *gorm.DB, userID int) (wishlist []*domain.Wishlist, msg string, err error) {
	if err := db.WithContext(ctx).Table("wishlist").Where("user_id = ?", userID).First(&wishlist).Error; err != nil {
		return nil, "Gagal get wishlist.", err
	}

	return wishlist, "Berhasil get wishlist.", nil
}

func (w *WishlistRepositoryImplementation) GetWishlists(ctx context.Context, db *gorm.DB) ([]*domain.Wishlist, string, error) {
	var wishlists []*domain.Wishlist
	if err := db.Table("wishlist").Find(&wishlists).Error; err != nil {
		return nil, "Gagal get wishlists.", err
	}

	return wishlists, "Berhasil get wishlists.", nil
}
