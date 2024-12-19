package main

import (
	"crud_server/internal/user/application/interfaces/input"
	"crud_server/internal/user/application/interfaces/output"
	"crud_server/internal/user/application/services"
	"crud_server/internal/user/insfrastructure/db"
	"crud_server/internal/user/insfrastructure/entity"
	"crud_server/internal/user/insfrastructure/handlers"
	"crud_server/internal/user/insfrastructure/persistence"
	"crud_server/internal/user/insfrastructure/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
			log.Println("Error loading .env file, using system environment variables")
	}

	db.DbConnection()

	db.DB.AutoMigrate(&entity.User{})

	userRepo := repository.NewUserRepository(db.DB)
	var userPersistence output.UserPersistence = persistence.NewUserPersistence(*userRepo)
	var userService input.UserService = services.NewUserServiceImpl(userPersistence)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/user", userHandler.ListUsers).Methods("GET")
	apiRouter.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	
	http.Handle("/", r)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)

}