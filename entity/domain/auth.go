package domain

type AuthRegister struct {
	ID       int    `json:"id" gorm:"->;primaryKey"`
	Name     string `json:"name" gorm:"<-;not null;varchar(50)"`
	NPM      string `json:"npm" gorm:"<-;not null;varchar(15)"`
	Email    string `json:"email" gorm:"<-;not null;varchar(150)"`
	Password string `json:"password" gorm:"<-;not null;varchar(150)"`
	IsGoogle bool   `json:"is_google" gorm:"<-;tinyint(1)"`
}
