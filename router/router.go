package router

import (
	"ToDo/views"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", views.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", views.GetTaskById).Methods("GET")
	return router
}
