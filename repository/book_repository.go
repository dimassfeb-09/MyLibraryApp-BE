package repository

import (
	"context"
	"fmt"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	AddBook(ctx context.Context, tx *gorm.DB, book *domain.Book) (isSuccess bool, msg string, err error)
	UpdateBook(ctx context.Context, tx *gorm.DB, book *domain.Book) (isSuccess bool, msg string, err error)
	DeleteBook(ctx context.Context, tx *gorm.DB, ID int) (isSuccess bool, msg string, err error)
	GetBookByID(ctx context.Context, db *gorm.DB, ID int) (book *domain.Book, msg string, err error)
	GetBooksByCategoryID(ctx context.Context, db *gorm.DB, categoryID int) (book []*domain.Book, msg string, err error)
	GetBooksByGenreID(ctx context.Context, db *gorm.DB, genreID int) (book []*domain.Book, msg string, err error)
	GetBookByTitle(ctx context.Context, db *gorm.DB, name string) (book []*domain.Book, msg string, err error)
	GetBookByQuery(ctx context.Context, db *gorm.DB, query string) (book []*domain.Book, msg string, err error)
	GetBooks(ctx context.Context, db *gorm.DB) (books []*domain.Book, msg string, err error)
}

type BookRepositoryImplementation struct {
}

func NewBookRepositoryImplementation() BookRepository {
	return &BookRepositoryImplementation{}
}

func (b *BookRepositoryImplementation) AddBook(ctx context.Context, tx *gorm.DB, book *domain.Book) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Create(&book).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal tambah buku.", err
	}

	return true, "Berhasil tambah buku.", nil
}

func (b *BookRepositoryImplementation) UpdateBook(ctx context.Context, tx *gorm.DB, book *domain.Book) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Where("id = ?", book.ID).Updates(&book).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal update buku.", err
	}

	return true, "Berhasil update buku.", nil
}

func (b *BookRepositoryImplementation) DeleteBook(ctx context.Context, tx *gorm.DB, ID int) (bool, string, error) {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("book").Where("id = ?", ID).Delete(ID).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return false, "Gagal hapus buku.", err
	}

	return true, "Berhasil hapus buku.", nil
}

func (b *BookRepositoryImplementation) GetBooksByCategoryID(ctx context.Context, db *gorm.DB, categoryID int) (books []*domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Where("category_id = ?", categoryID).Find(&books).Error; err != nil {
		return nil, "Gagal Get Book by Category ID.", err
	}
	fmt.Println(books)
	return books, "Berhasil Get Book by Category ID.", nil
}

func (b *BookRepositoryImplementation) GetBooksByGenreID(ctx context.Context, db *gorm.DB, genreID int) (books []*domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Where("genre_id = ?", genreID).Find(&books).Error; err != nil {
		return nil, "Gagal Get Book by Genre ID.", err
	}
	fmt.Println(books)
	return books, "Berhasil Get Book by Genre ID.", nil
}

func (b *BookRepositoryImplementation) GetBookByID(ctx context.Context, db *gorm.DB, ID int) (book *domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Where("id = ?", ID).First(&book).Error; err != nil {
		return nil, "Gagal Get User by ID.", err
	}
	return book, "Berhasil Get User by ID.", nil
}

func (b *BookRepositoryImplementation) GetBookByTitle(ctx context.Context, db *gorm.DB, title string) (books []*domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Where("title = ?", "%", title).Find(&books).Error; err != nil {
		return nil, "Gagal get data kategoris.", err
	}
	return books, "Berhasil get data kategoris.", nil
}

func (b *BookRepositoryImplementation) GetBookByQuery(ctx context.Context, db *gorm.DB, query string) (books []*domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Joins("JOIN category ON category.id = book.category_id").Where("(title LIKE ? OR writer LIKE ? OR category.name = ?)", "%"+query+"%", "%"+query+"%", query).Find(&books).Error; err != nil {
		return nil, "Gagal get data kategoris.", err
	}
	return books, "Berhasil get data kategoris.", nil
}

func (b *BookRepositoryImplementation) GetBooks(ctx context.Context, db *gorm.DB) (books []*domain.Book, msg string, err error) {
	if err := db.WithContext(ctx).Table("book").Find(&books).Error; err != nil {
		return nil, "Gagal get data kategoris.", err
	}
	return books, "Berhasil get data kategoris.", nil
}
