package kategoripenerima

type Service interface {
	FindKategoriPenerima() ([]KategoriPenerima, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindKategoriPenerima() ([]KategoriPenerima, error) {
	katPenerima, err := s.repository.FindAll()
	if err != nil {
		return katPenerima, err
	}

	return katPenerima, nil
}
