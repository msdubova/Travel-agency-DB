package booking

type storage interface {
	Create(b Booking)
	GetBookings() []Booking
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{
		s: s,
	}
}

func (s *Service) CreateBooking(b Booking) {
	s.s.Create(b)
}

func (s *Service) GetBookings() []Booking {
	return s.s.GetBookings()
}
