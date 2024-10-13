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

	log.Printf("Server is running at %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Fatalf("server stopped")
}
