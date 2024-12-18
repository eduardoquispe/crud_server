package main

import (
	"crud_server/internal/user/application/interfaces/input"
	"crud_server/internal/user/application/interfaces/output"
	"crud_server/internal/user/application/services"
	"crud_server/internal/user/insfrastructure/db"
	"crud_server/internal/user/insfrastructure/entity"
	"crud_server/internal/user/insfrastructure/handlers"
	"crud_server/internal/user/insfrastructure/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.DbConnection()

	db.DB.AutoMigrate(&entity.User{})

	var userRepo output.UserRepository = repository.NewUserRepository(db.DB)
	var userService input.UserService = services.NewUserServiceImpl(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler.ListUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}