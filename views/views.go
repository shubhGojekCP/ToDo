package views

import (
	"ToDo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(400)
		return
	}
	for _, item := range Data {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var response http.Response
	w.Header().Set("Context-Type", "application/json")
	params, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		return
	}
	for index, item := range Data {
		if item.Id == params {
			Data = append(Data[:index], Data[index+1:]...)
			response.Status = "Deleted Successfully"
			response.StatusCode = 200

			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(404)

}
