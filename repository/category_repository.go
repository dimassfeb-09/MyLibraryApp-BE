package repository

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	AddCategory(ctx context.Context, tx *gorm.DB, category *domain.Category) (isSuccess bool, msg string, err error)
	UpdateCategory(ctx context.Context, tx *gorm.DB, category *domain.Category) (isSuccess bool, msg string, err error)
	DeleteCategory(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
	GetCategoryByID(ctx context.Context, db *gorm.DB, ID int) (category *domain.Category, msg string, err error)
	GetCategoryByName(ctx context.Context, db *gorm.DB, name string) (category *domain.Category, msg string, err error)
	GetCategories(ctx context.Context, db *gorm.DB) (categories []*domain.Category, msg string, err error)
}

type CategoryRepositoryImplementation struct {
}

func NewCategoryRepositoryImplementation() CategoryRepository {
	return &CategoryRepositoryImplementation{}
}

func (c *CategoryRepositoryImplementation) AddCategory(ctx context.Context, tx *gorm.DB, category *domain.Category) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("category").Create(&category).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal tambah kategori.", err
	}

	return true, "Berhasil tambah kategori.", nil
}

func (c *CategoryRepositoryImplementation) UpdateCategory(ctx context.Context, tx *gorm.DB, category *domain.Category) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("category").Where("id = ?", category.ID).Updates(&category).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update kategori.", err
	}

	return true, "Berhasil update kategori.", nil
}

func (c *CategoryRepositoryImplementation) DeleteCategory(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("category").Where("id = ?", ID).Delete(ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal hapus kategori.", err
	}

	return true, "Berhasil hapus kategori.", nil
}

func (c *CategoryRepositoryImplementation) GetCategoryByID(ctx context.Context, db *gorm.DB, ID int) (category *domain.Category, msg string, err error) {
	if err := db.WithContext(ctx).Table("category").Where("id = ?", ID).First(&category).Error; err != nil {
		return nil, "Gagal get data kategori.", err
	}

	return category, "Berhasil get data kategori.", nil
}

func (c *CategoryRepositoryImplementation) GetCategoryByName(ctx context.Context, db *gorm.DB, name string) (category *domain.Category, msg string, err error) {
	if err := db.WithContext(ctx).Table("category").Where("name = ?", name).First(&category).Error; err != nil {
		return nil, "Gagal get data kategori.", err
	}

	return category, "Berhasil get data kategori.", nil
}

func (c *CategoryRepositoryImplementation) GetCategories(ctx context.Context, db *gorm.DB) (categories []*domain.Category, msg string, err error) {
	if err := db.WithContext(ctx).Table("category").Find(&categories).Error; err != nil {
		return nil, "Gagal get data kategoris.", err
	}
	return categories, "Berhasil get data kategoris.", nil
}
