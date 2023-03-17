package service

import (
	"context"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type CategoryService interface {
	AddCategory(ctx context.Context, category *request.Category) (isSuccess bool, msg string, err error)
	UpdateCategory(ctx context.Context, category *request.Category) (isSuccess bool, msg string, err error)
	DeleteCategory(ctx context.Context, ID int) (isSuccess bool, msg string, err error)
	GetCategoryByID(ctx context.Context, ID int) (category *response.Category, msg string, err error)
	GetCategoryByName(ctx context.Context, name string) (category *response.Category, msg string, err error)
	GetCategories(ctx context.Context) (categories []*response.Category, msg string, err error)
}

type CategoryServiceImplementation struct {
	DB                 *gorm.DB
	CategoryRepository repository.CategoryRepository
}

func NewCategoryServiceImplementation(DB *gorm.DB, categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImplementation{DB: DB, CategoryRepository: categoryRepository}
}

func (c *CategoryServiceImplementation) AddCategory(ctx context.Context, r *request.Category) (bool, string, error) {
	category := &domain.Category{
		Name: r.Name,
	}
	isSuccess, msg, err := c.CategoryRepository.AddCategory(ctx, c.DB, category)
	if err != nil {
		return false, msg, err
	}

	return isSuccess, msg, nil
}

func (c *CategoryServiceImplementation) UpdateCategory(ctx context.Context, r *request.Category) (bool, string, error) {
	_, msg, err := c.GetCategoryByID(ctx, r.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Kategori ID tidak ditemukan.", err
	}

	category := &domain.Category{
		ID:   r.ID,
		Name: r.Name,
	}

	isSuccess, msg, err := c.CategoryRepository.UpdateCategory(ctx, c.DB, category)
	if err != nil {
		return false, msg, err
	}

	return isSuccess, msg, nil
}

func (c *CategoryServiceImplementation) DeleteCategory(ctx context.Context, ID int) (bool, string, error) {
	_, msg, err := c.GetCategoryByID(ctx, ID)
	if err != nil {
		return false, msg, err
	}

	isSuccess, msg, err := c.CategoryRepository.DeleteCategory(ctx, c.DB, ID)
	if err != nil {
		return false, msg, err
	}

	return isSuccess, msg, nil
}

func (c *CategoryServiceImplementation) GetCategoryByID(ctx context.Context, ID int) (*response.Category, string, error) {
	result, msg, err := c.CategoryRepository.GetCategoryByID(ctx, c.DB, ID)
	if err == gorm.ErrRecordNotFound {
		return nil, "Kategori ID tidak ditemukan.", err
	}

	category := &response.Category{
		ID:   result.ID,
		Name: result.Name,
	}
	return category, msg, nil
}

func (c *CategoryServiceImplementation) GetCategoryByName(ctx context.Context, name string) (*response.Category, string, error) {
	result, msg, err := c.CategoryRepository.GetCategoryByName(ctx, c.DB, name)
	if err == gorm.ErrRecordNotFound {
		return nil, "Kategori Name tidak ditemukan.", err
	}

	category := &response.Category{
		ID:   result.ID,
		Name: result.Name,
	}
	return category, msg, nil
}

func (c *CategoryServiceImplementation) GetCategories(ctx context.Context) ([]*response.Category, string, error) {
	results, msg, err := c.CategoryRepository.GetCategories(ctx, c.DB)
	if err != nil {
		return nil, msg, err
	}

	var categories []*response.Category
	for _, result := range results {
		category := &response.Category{
			ID:   result.ID,
			Name: result.Name,
		}
		categories = append(categories, category)
	}

	return categories, msg, nil
}
