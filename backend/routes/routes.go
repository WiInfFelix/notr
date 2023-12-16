package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/WiInfFelix/notr/ent"
	"github.com/WiInfFelix/notr/handlers"
	"github.com/WiInfFelix/notr/services"
	"github.com/WiInfFelix/notr/utils"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func SetupRoutes(app chi.Router) {
	// client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=notr_db password=password sslmode=disable")

	//client with sqlite file
	log.Println("Connecting to sqlite...")
	client, err := ent.Open("sqlite3", "./ent.db?_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	log.Println("Connected to sqlite...")

	// run migration
	log.Println("Running migration...")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Migration successful...")

	// services
	userService := services.NewUserService(client)

	// handlers
	userHandler := handlers.NewUserHandler(*userService)

	// routes
	app.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "pong"})
	})

	app.Get("/users", userHandler.GetAllUsersHandler)
	app.Post("/users", userHandler.CreateUserHandler)

}
