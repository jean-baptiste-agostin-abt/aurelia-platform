package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index"`
	Password string
	FamilyID uint
}

func NewUser(email, password string, familyID uint) *User {
	return &User{Email: email, Password: password, FamilyID: familyID}
}
