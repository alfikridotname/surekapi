package script

type Service interface {
	GetScripts() ([]Script, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetScripts() ([]Script, error) {
	scripts, err := s.repository.FindAll()
	if err != nil {
		return scripts, err
	}

	return scripts, nil
}
