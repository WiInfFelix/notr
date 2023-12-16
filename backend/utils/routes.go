package utils

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PrintRoutes(app chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(app, walkFunc); err != nil {
		log.Printf("Logging err: %s\n", err.Error())
	}
}
