package legacyguard

import "github.com/jinzhu/gorm"

// Repository defines DB ops for LegacyGuard
//
//go:generate mockery --name Repository
type Repository interface {
	Create(l *LegacyGuard) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(l *LegacyGuard) error {
	return r.db.Create(l).Error
}
