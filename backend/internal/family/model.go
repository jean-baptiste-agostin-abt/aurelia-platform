package family

import (
	"github.com/jinzhu/gorm"
	"github.com/yourorg/aurelia-backend/internal/user"
)

type Family struct {
	gorm.Model
	Name  string
	Users []user.User
}
