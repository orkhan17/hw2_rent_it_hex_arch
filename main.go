package main

import (
	"fmt"
	repository2 "github.com/orkhan17/hw2-rent_it_hex_arch/pkg/repositories/plant"
	"github.com/orkhan17/hw2-rent_it_hex_arch/pkg/service"
	http2 "github.com/orkhan17/hw2-rent_it_hex_arch/pkg/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	// SQL driver
	_ "github.com/lib/pq"
)

const (
	logLevel        = "debug"
	httpServicePort = 8080
	websocketServicePort = 8081
)

//TODO: Here start the 3 servers asynchrounusly using goroutines.
func main() {
	// begin setup
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	log.Info("Start server")

	// construct application
	bookRepo := repository2.NewPlantRepostory()
	bookService := service.NewPlantService(bookRepo)
	bookHTTPHandler := http2.NewPlantHandler(bookService)


	httpRouter := mux.NewRouter()
	websocketRouter := mux.NewRouter()

	bookHTTPHandler.RegisterRoutes(httpRouter)


	// setup http server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}



	err = httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Infof("Stoped server")
}
