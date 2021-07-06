package views

import (
	"ToDo/services"
	"encoding/json"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := services.AddTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := services.GetDataById(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	response := services.RemoveTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	response := services.SvcUpdateTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := services.GetAllData()
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}
