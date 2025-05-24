package capsule

import "github.com/jinzhu/gorm"

type Repository interface {
	Create(c *Capsule) error
	GetByID(id uint) (*Capsule, error)
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(c *Capsule) error {
	return r.db.Create(c).Error
}

func (r *gormRepository) GetByID(id uint) (*Capsule, error) {
	var cap Capsule
	if err := r.db.First(&cap, id).Error; err != nil {
		return nil, err
	}
	return &cap, nil
}
