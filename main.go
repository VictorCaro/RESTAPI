package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/VictorCaro/go-restapi/db"
	"github.com/VictorCaro/go-restapi/models"
	"github.com/VictorCaro/go-restapi/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("Post")
	r.HandleFunc("/tasks", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
