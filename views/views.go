package views

import (
	"ToDo/models"
	"encoding/json"
	"net/http"
)

var Data []models.ToDoList

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.ToDoList
	var response http.Response
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&newTask)
	Data = append(Data, newTask)
	w.WriteHeader(201)
	response.Status = "Creates Successfully"
	response.StatusCode = 201
	json.NewEncoder(w).Encode(response)

}
