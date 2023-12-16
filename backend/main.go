package main

import (
	"log"
	"net/http"

	"github.com/WiInfFelix/notr/routes"
	"github.com/WiInfFelix/notr/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./api/docs"))))

	routes.SetupRoutes(r)

	// print all routes
	utils.PrintRoutes(r)

	log.Println(http.ListenAndServe(":8000", r))
}
