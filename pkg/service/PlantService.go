package service

//here implement the service which makes the db operations (implements core/PlantServicePort)
//Service has both Mongo annd Postgres Repository implementation which it uses for Data acess.

import "github.com/orkhan17/hw2-rent_it_hex_arch/pkg/internal/core/domain"

type plantRepository interface {
	GetAllPlants() ([]*domain.Plant, error)
}

type PlantService struct {
	plantRepository plantRepository
}

func NewPlantService(bR plantRepository) *PlantService {
	return &PlantService{
		plantRepository: bR,
	}
}


func (s *PlantService) GetAllPlants() ([]*domain.Book, error){
	books, err := s.plantRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}
