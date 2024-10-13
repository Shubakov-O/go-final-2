package tests

import (
	"net/http/httptest"

	"github.com/Shubakov-O/go-final-2/internal/config"
	"github.com/Shubakov-O/go-final-2/pkg/router"

	_ "github.com/mattn/go-sqlite3"
)

func createTestServer() *httptest.Server {
	config.MustLoad()

	router := router.SetupRouter()

	return httptest.NewServer(router)
}
