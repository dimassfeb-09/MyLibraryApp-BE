package request

type User struct {
	ID       int    `json:"id" gorm:"->;primaryKey"`
	Name     string `binding:"required" json:"name" gorm:"<-;not null"`
	NPM      string `binding:"required" json:"npm" gorm:"<-;not null"`
	Email    string `binding:"required" json:"email" gorm:"<-;not null"`
	Password string `binding:"required" json:"password" gorm:"->:false;<-;not null"`
	IsGoogle bool   `json:"is_google" gorm:"<-"`
}
