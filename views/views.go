package views

import (
	"ToDo/services"
	"encoding/json"
	"net/http"
)

// CreateTask creates task based on valid data provided,else return 400 for
// invalid Data,or 200 if Task with ID already exists.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> CreateTask")
	w.Header().Set("Content-Type", "application/json")
	response := services.SvcAddTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

// GetTaskById gives the task based on Id provided,else return 400 for Invalid ID ,
// or 404 if Task with ID does not exists.
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> GetTaskById")
	w.Header().Set("Content-Type", "application/json")
	response := services.SvcGetDataById(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

// DeleteTask deletes a task based on valid ID provided,else return 400 for
// invalid ID,or 404 if Task with ID does not exists.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> DeleteTask")
	w.Header().Set("Context-Type", "application/json")
	response := services.SvcRemoveTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

// UpdateTaskStatus updates a task based on valid data provided,else return 400 for
// invalid data,or 404 if Task with ID does not exists.
func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> UpdatetaskStatus")
	w.Header().Set("Context-Type", "application/json")
	response := services.SvcUpdateTask(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}

//GetAllTask gets all the tasks.
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> GetAllTask")
	w.Header().Set("Content-Type", "application/json")
	response := services.SvcGetAllData()
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}
