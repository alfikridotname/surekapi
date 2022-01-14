package recipient_category

import "gorm.io/gorm"

type Repository interface {
	GetRecipientCategory() ([]RecipientCategory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetRecipientCategory() ([]RecipientCategory, error) {
	var recipientCategory []RecipientCategory
	err := r.db.Find(&recipientCategory).Error
	if err != nil {
		return recipientCategory, err
	}

	return recipientCategory, nil
}
