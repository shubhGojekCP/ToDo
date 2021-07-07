package controller

import (
	"ToDo/services"
	"encoding/json"
	"net/http"
)

// CreateTask creates task based on valid data provided,else return 400 for
// invalid Data,or 200 if Task with ID already exists.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> CreateTask")
	commonReturn(w, r, services.SvcAddTask)

}

// GetTaskById gives the task based on Id provided,else return 400 for Invalid ID ,
// or 404 if Task with ID does not exists.
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> GetTaskById")
	commonReturn(w, r, services.SvcGetDataById)

}

// DeleteTask deletes a task based on valid ID provided,else return 400 for
// invalid ID,or 404 if Task with ID does not exists.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> DeleteTask")
	commonReturn(w, r, services.SvcRemoveTask)

}

// UpdateTaskStatus updates a task based on valid data provided,else return 400 for
// invalid data,or 404 if Task with ID does not exists.
func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> UpdatetaskStatus")
	commonReturn(w, r, services.SvcUpdateTask)

}

//GetAllTask gets all the tasks.
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	services.InfoLogger.Println(">> GetAllTask")
	commonReturn(w, r, services.SvcGetAllData)

}

func commonReturn(w http.ResponseWriter, r *http.Request, f func(*http.Request) services.Responses) {

	w.Header().Set("Content-Type", "application/json")
	response := f(r)
	w.WriteHeader(response.KnowStatus())
	json.NewEncoder(w).Encode(response)

}
