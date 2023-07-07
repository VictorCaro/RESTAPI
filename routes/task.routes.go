package routes

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCaro/go-restapi/db"
	"github.com/VictorCaro/go-restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}
