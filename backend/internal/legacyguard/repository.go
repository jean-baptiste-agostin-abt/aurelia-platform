package legacyguard

import "github.com/jinzhu/gorm"

type Repository interface {
	Trigger(lg *LegacyGuard) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Trigger(lg *LegacyGuard) error {
	return r.db.Create(lg).Error
}
