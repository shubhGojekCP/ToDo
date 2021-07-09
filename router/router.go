package router

import (
	"ToDo/controller"
	"ToDo/model"
	"ToDo/services"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	controller := controller.Handler{
		Service: services.Service{DataStore: model.Storage{}},
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.GetTaskById).Methods("GET")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task", controller.UpdateTaskStatus).Methods("PUT")
	router.HandleFunc("/api/task", controller.GetAllTask).Methods("GET")

	return router
}
