package booking

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

var s InMemStorage

type service interface {
	CreateBooking(b Booking)
	GetBookings() []Booking
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{
		s: s,
	}
}

func (h Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var b Booking

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON request body")
		return
	}

	b.ReservationNumber = strconv.Itoa(rand.Intn(100000))
	b.Status = "pending"

	h.s.CreateBooking(b)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}
}

func (h Handler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings := h.s.GetBookings()

	err := json.NewEncoder(w).Encode(bookings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
