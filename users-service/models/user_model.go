package models

type User struct {
	UserID   int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255);not null;unique"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Pwd      string `gorm:"type:varchar(255);not null"`
}

type Users []User
