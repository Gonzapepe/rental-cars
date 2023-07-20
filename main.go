package main

import (
	"net/http"

	"github.com/Gonzapepe/cars-rental/config"
	"github.com/Gonzapepe/cars-rental/helper"
	"github.com/Gonzapepe/cars-rental/migrations"
	"github.com/Gonzapepe/cars-rental/repository"
	"github.com/Gonzapepe/cars-rental/router"
)

func main() {
	// Database
	db := config.DatabaseConnection()

	// Run Migrations
	err := migrations.Run(db)
	helper.ErrorPanic(err)

	carRepository := repository.NewCarRepository(db)

	routes := router.NewRouter(carRepository)

	server := &http.Server{
		Addr: ":8080",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}