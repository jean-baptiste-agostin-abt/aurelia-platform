package family

import "github.com/jinzhu/gorm"

// Repository defines DB operations for Family
//
//go:generate mockery --name Repository
type Repository interface {
	Create(f *Family) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(f *Family) error {
	return r.db.Create(f).Error
}
