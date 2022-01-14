package recipient_category

type Service interface {
	GetRecipientCategory() ([]RecipientCategory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetRecipientCategory() ([]RecipientCategory, error) {
	recipientCategory, err := s.repository.GetRecipientCategory()
	if err != nil {
		return recipientCategory, err
	}

	return recipientCategory, nil
}
