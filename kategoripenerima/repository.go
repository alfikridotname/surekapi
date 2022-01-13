package kategoripenerima

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]KategoriPenerima, error)
	FindByID(ID int) ([]KategoriPenerima, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]KategoriPenerima, error) {
	var katpenerima []KategoriPenerima
	err := r.db.Find(&katpenerima).Error
	if err != nil {
		return katpenerima, err
	}

	return katpenerima, nil
}

func (r *repository) FindByID(ID int) ([]KategoriPenerima, error) {
	var katpenerima []KategoriPenerima
	err := r.db.Where("id = ?", ID).Find(&katpenerima).Error
	if err != nil {
		return katpenerima, err
	}

	return katpenerima, nil
}
