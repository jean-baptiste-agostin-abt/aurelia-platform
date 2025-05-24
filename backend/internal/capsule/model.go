package capsule

import "github.com/jinzhu/gorm"

type Capsule struct {
	gorm.Model
	Title    string
	Content  string
	FamilyID uint
}

func NewCapsule(title, content string, familyID uint) *Capsule {
	return &Capsule{Title: title, Content: content, FamilyID: familyID}
}
