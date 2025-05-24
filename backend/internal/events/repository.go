package events

import "github.com/jinzhu/gorm"

type Repository interface {
	List() ([]Event, error)
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) List() ([]Event, error) {
	var evs []Event
	err := r.db.Find(&evs).Error
	return evs, err
}
