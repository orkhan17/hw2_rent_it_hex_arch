package repositories

//TODO: here we implement the ports/PlantRepository Interface for Mongodb
import "github.com/orkhan17/hw2-rent_it_hex_arch/pkg/internal/core/domain"

type PostgresPlantRepository struct {
	data []*domain.Plant
}

func NewPlantRepostory() *PostgresPlantRepository {
	return &PostgresPlantRepository{
		data: []*domain.Plant{},
	}
}

func (r *PostgresPlantRepository) GetAllPlants() ([]*domain.Plant, error) {
	return r.data, nil
}