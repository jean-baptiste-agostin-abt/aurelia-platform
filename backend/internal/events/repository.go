package events

import "github.com/jinzhu/gorm"

// Repository defines DB operations for Event
//
//go:generate mockery --name Repository
type Repository interface {
	List() ([]Event, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) List() ([]Event, error) {
	var evs []Event
	if err := r.db.Find(&evs).Error; err != nil {
		return nil, err
	}
	return evs, nil
}
