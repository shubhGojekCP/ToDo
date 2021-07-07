package router

import (
	"ToDo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.GetTaskById).Methods("GET")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task", controller.UpdateTaskStatus).Methods("PUT")
	router.HandleFunc("/api/task", controller.GetAllTask).Methods("GET")

	return router
}
