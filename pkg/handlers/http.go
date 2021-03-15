package http

// this file serves the role of the adapter or handler
//it receives an http request and based on it it is able to call the correct service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/orkhan17/hw2-rent_it_hex_arch/pkg/internal/core/domain"
	)

type plantService interface {
	GetAllPlants() ([]*domain.Plant, error)
}

type PlantHandler struct {
	plantService plantService
}

func NewPlantHandler(bS plantService) *PlantHandler {
	return &PlantHandler{
		plantService: bS,
	}
}

func (h *PlantHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plant", h.getAllPlants).Methods(http.MethodGet)
}



func (h *PlantHandler) getAllPlants(w http.ResponseWriter, _ *http.Request) {
	books, err := h.plantService.GetAllBooks()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&books)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}