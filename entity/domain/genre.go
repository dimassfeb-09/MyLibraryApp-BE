package domain

type Genre struct {
	ID         int    `json:"id"`
	Name       string `json:"name" gorm:"<-"`
	CategoryID int    `json:"category_id" gorm:"<-"`
}