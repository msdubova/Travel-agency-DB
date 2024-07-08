package tour

type storage interface {
	GetTours() []Tour
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{
		s: s,
	}
}

func (s *Service) GetTours() []Tour {
	return s.s.GetTours()
}
