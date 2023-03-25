package domain

import "time"

type Book struct {
	ID          int       `json:"id" gorm:"<-:false"`
	Title       string    `json:"title" gorm:"<-"`
	Description string    `json:"description" gorm:"<-"`
	Stok        int       `json:"stok" gorm:"<-"`
	Writer      string    `json:"writer" gorm:"<-"`
	ImgURL      string    `json:"img_url" gorm:"<-"`
	Rating      float32   `json:"rating" gorm:"<-"`
	CategoryID  int       `json:"category_id" gorm:"<-"`
	GenreID     int       `json:"genre_id" gorm:"<-"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"<-"`
}
