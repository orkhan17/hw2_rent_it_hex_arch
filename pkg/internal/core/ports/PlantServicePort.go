package ports

import (
	"rent_it_hex/internal/core/domain"
	"time"
)

//Driver port
//This contains the footprint that different adapters must call
type PlantServicePort interface {
	GetAll() ([]*domain.Plant, error)
	CalculatePrice(startDate time.Time, endDate time.Time, id int32) (float64, error)
	CheckAvailability(startDate time.Time, endDate time.Time) (bool, error)
}
