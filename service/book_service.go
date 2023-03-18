package service

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"gorm.io/gorm"
)

type BookService interface {
	AddBook(ctx context.Context, r *request.Book) (bool, string, error)
	UpdateBook(ctx context.Context, r *request.Book) (bool, string, error)
	DeleteBook(ctx context.Context, ID int) (bool, string, error)
	GetBookByID(ctx context.Context, ID int) (book *response.Book, msg string, err error)
	GetBookByTitle(ctx context.Context, title string) (book []*response.Book, msg string, err error)
	GetBooks(ctx context.Context) (books []*response.Book, msg string, err error)
}

type BookServiceImplementation struct {
	db              *gorm.DB
	BookRepository  repository.BookRepository
	CategoryService CategoryService
}

func NewBookServiceImplementation(db *gorm.DB, bookRepository repository.BookRepository, categoryService CategoryService) BookService {
	return &BookServiceImplementation{db: db, BookRepository: bookRepository, CategoryService: categoryService}
}

func (b *BookServiceImplementation) AddBook(ctx context.Context, r *request.Book) (bool, string, error) {

	_, _, err := b.CategoryService.GetCategoryByID(ctx, r.CategoryID)
	if err == gorm.ErrRecordNotFound {
		return false, "Kategori tidak ditemukan.", err
	}

	book := &domain.Book{
		Title:       r.Title,
		Description: r.Description,
		Stok:        r.Stok,
		Writer:      r.Writer,
		ImgURL:      r.ImgURL,
		Rating:      r.Rating,
		CategoryID:  r.CategoryID,
	}

	return b.BookRepository.AddBook(ctx, b.db, book)
}

func (b *BookServiceImplementation) UpdateBook(ctx context.Context, r *request.Book) (bool, string, error) {

	_, _, err := b.CategoryService.GetCategoryByID(ctx, r.CategoryID)
	if err == gorm.ErrRecordNotFound {
		return false, "Kategori tidak ditemukan.", err
	}

	_, _, err = b.GetBookByID(ctx, r.ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Data buku tidak ditemukan.", err
	}

	book := &domain.Book{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Stok:        r.Stok,
		Writer:      r.Writer,
		ImgURL:      r.ImgURL,
		Rating:      r.Rating,
		CategoryID:  r.CategoryID,
	}

	fmt.Println(book)

	isSuccess, msg, err := b.BookRepository.UpdateBook(ctx, b.db, book)
	if err != nil {
		return false, msg, err
	}

	return isSuccess, msg, nil
}

func (b *BookServiceImplementation) DeleteBook(ctx context.Context, ID int) (bool, string, error) {
	_, _, err := b.GetBookByID(ctx, ID)
	if err == gorm.ErrRecordNotFound {
		return false, "Data buku tidak ditemukan.", err
	}

	return b.BookRepository.DeleteBook(ctx, b.db, ID)
}

func (b *BookServiceImplementation) GetBookByID(ctx context.Context, ID int) (*response.Book, string, error) {
	result, msg, err := b.BookRepository.GetBookByID(ctx, b.db, ID)
	if err == gorm.ErrRecordNotFound {
		return nil, "Data buku tidak ditemukan.", err
	}

	book := &response.Book{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Stok:        result.Stok,
		Writer:      result.Writer,
		ImgURL:      result.ImgURL,
		Rating:      result.Rating,
		CategoryID:  result.CategoryID,
	}

	return book, msg, nil

}

func (b *BookServiceImplementation) GetBookByTitle(ctx context.Context, title string) (books []*response.Book, msg string, err error) {
	results, msg, err := b.BookRepository.GetBookByTitle(ctx, b.db, title)
	if err != nil {
		return nil, msg, err
	}

	for _, result := range results {
		book := &response.Book{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Stok:        result.Stok,
			Writer:      result.Writer,
			ImgURL:      result.ImgURL,
			Rating:      result.Rating,
			CategoryID:  result.CategoryID,
		}
		books = append(books, book)
	}

	return books, "Sukses get books by name", nil
}

func (b *BookServiceImplementation) GetBooks(ctx context.Context) (books []*response.Book, msg string, err error) {

	results, msg, err := b.BookRepository.GetBooks(ctx, b.db)
	if err != nil {
		return nil, msg, err
	}

	for _, result := range results {
		book := &response.Book{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Stok:        result.Stok,
			Writer:      result.Writer,
			ImgURL:      result.ImgURL,
			Rating:      result.Rating,
			CategoryID:  result.CategoryID,
		}
		books = append(books, book)
	}

	return books, msg, nil
}
