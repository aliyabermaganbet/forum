package app

import (
	"log"

	"forum/internal/delivery"
	"forum/internal/server"
	"forum/internal/service"
	"forum/internal/storage"
)

func Run() {
	db, err := storage.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.CreateTables(db); err != nil {
		log.Fatal(err)
	}
	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := delivery.NewHandler(services)
	server := new(server.Server)
	if err := server.Start(handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
