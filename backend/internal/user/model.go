package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index"`
	Password string
	FamilyID uint
}
