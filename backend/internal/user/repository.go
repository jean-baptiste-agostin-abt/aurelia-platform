package user

import "github.com/jinzhu/gorm"

// Repository abstracts DB operations on User
//
//go:generate mockery --name Repository
type Repository interface {
	Create(u *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
}

// GormRepository implements Repository with GORM
type GormRepository struct {
	db *gorm.DB
}

// NewRepository returns a repository backed by GORM
func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *GormRepository) FindByEmail(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *GormRepository) FindByID(id uint) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
