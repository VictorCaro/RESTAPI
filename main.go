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
	http.ListenAndServe(":3000", r)

}
