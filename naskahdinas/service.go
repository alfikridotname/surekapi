package naskahdinas

type Service interface {
	FindNaskahDinas() ([]NaskahDinas, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindNaskahDinas() ([]NaskahDinas, error) {
	naskahDinas, err := s.repository.FindAll()
	if err != nil {
		return naskahDinas, err
	}

	return naskahDinas, nil
}
