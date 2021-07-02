package router

import (
	"ToDo/views"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", views.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", views.GetTaskById).Methods("GET")
	router.HandleFunc("/api/task/delete/{id}", views.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task/update", views.UpdateTaskStatus).Methods("PUT")

	return router
}
