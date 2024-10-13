package router

import (
	"log"
	"net/http"

	"github.com/Shubakov-O/go-final-2/internal/scheduler"
	"github.com/Shubakov-O/go-final-2/internal/storage/sqlite"

	"github.com/Shubakov-O/go-final-2/internal/config"
	"github.com/Shubakov-O/go-final-2/internal/http-server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {

	webDir := config.WebDirPath

	storage, err := sqlite.New(config.DBFilePath)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	planner := scheduler.NewScheduler(storage)

	router := chi.NewRouter()

	router.Use(middleware.URLFormat)

	router.Handle("/*", http.FileServer(http.Dir(webDir)))

	handlers.RegisterRoutes(router, planner)

	return router
}
