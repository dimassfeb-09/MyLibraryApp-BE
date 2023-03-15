package domain

type User struct {
	ID       int
	Name     string `gorm:"<-"`
	NPM      string `gorm:"<-"`
	Email    string `gorm:"<-"`
	Password string `gorm:"<-;->:false"`
	IsGoogle bool   `gorm:"<-"`
}
