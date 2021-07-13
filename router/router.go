package router

import (
	"ToDo/controller"
	"ToDo/model"
	"ToDo/services"
	"context"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	ctx, _ := context.WithCancel(context.Background())
	Storage := model.Connect(ctx)
	controller := controller.Handler{
		Service: services.Service{DataStore: Storage},
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.GetTaskById).Methods("GET")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task", controller.UpdateTaskStatus).Methods("PUT")
	router.HandleFunc("/api/task", controller.GetAllTask).Methods("GET")

	return router
}
