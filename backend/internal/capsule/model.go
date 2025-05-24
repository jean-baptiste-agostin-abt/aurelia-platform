package capsule

import "github.com/jinzhu/gorm"

type Capsule struct {
	gorm.Model
	Title    string
	Content  string
	FamilyID uint
}
