package family

import "github.com/jinzhu/gorm"

type Repository interface {
	Create(f *Family) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(f *Family) error {
	return r.db.Create(f).Error
}
