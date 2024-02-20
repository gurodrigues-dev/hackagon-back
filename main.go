package main

import (
	"gin/config"
	"gin/internal/controllers"
	"gin/internal/service"
	"gin/repository"
	"log"
)

func main() {
	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	repo, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("error creating repository: %s", err.Error())
	}

	service := service.New(repo)

	controller := controllers.New(service)

	log.Printf("initing service: %s", config.Name)
	controller.Start()

}
