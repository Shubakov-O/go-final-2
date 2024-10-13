package main

import (
	"log"
	"net/http"

	"github.com/Shubakov-O/go-final-2/internal/config"
	"github.com/Shubakov-O/go-final-2/pkg/router"
)

func main() {
	config.MustLoad()

	port := ":" + config.Port

	router := router.SetupRouter()

	log.Printf("Сервер работает по адресу %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	log.Fatalf("сервер остановлен")
}
