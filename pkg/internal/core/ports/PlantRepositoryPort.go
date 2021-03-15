package ports

import (
	"rent_it_hex/internal/core/domain"
	"time"
)

//Driven port
type PlantRepository interface {
	GetAll() ([]*domain.Plant, error)
	CalculatePrice(startDate time.Time, endDate time.Time, id int32) (float64, error)
	CheckAvailability(startDate time.Time, endDate time.Time) (bool, error)
}
