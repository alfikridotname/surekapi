package script

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Script, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Script, error) {
	var script []Script
	err := r.db.Where("is_active = ?", 1).Find(&script).Error
	if err != nil {
		return script, err
	}

	return script, nil
}
