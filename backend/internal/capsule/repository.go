package capsule

import "github.com/jinzhu/gorm"

// Repository abstracts Capsule persistence
//
//go:generate mockery --name Repository
type Repository interface {
	Create(c *Capsule) error
	Find(id uint) (*Capsule, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(c *Capsule) error {
	return r.db.Create(c).Error
}

func (r *GormRepository) Find(id uint) (*Capsule, error) {
	var cap Capsule
	if err := r.db.First(&cap, id).Error; err != nil {
		return nil, err
	}
	return &cap, nil
}
