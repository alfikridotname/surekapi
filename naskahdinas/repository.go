package naskahdinas

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]NaskahDinas, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]NaskahDinas, error) {
	var naskahDinas []NaskahDinas
	err := r.db.Where("is_active = ?", true).Find(&naskahDinas).Error
	if err != nil {
		return naskahDinas, err
	}

	return naskahDinas, nil
}
