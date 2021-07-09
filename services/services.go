package services

import (
	"ToDo/controller"
	"ToDo/model"
	"ToDo/utils"
)

type Service struct {
	DataStore Store
}

type Store interface {
	AddTask(data model.ToDoList) model.ToDoList
	GetById(id int) (model.ToDoList, error)
	RemoveById(id int) (model.ToDoList, error)
	UpdateTask(data model.ToDoList) (model.ToDoList, error)
	AllTask() ([]model.ToDoList, error)
}

func (s Service) SvcAddTask(data controller.ToDo) controller.ToDo {
	utils.InfoLogger.Println(">> AddTask")
	res := s.DataStore.AddTask(model.ToDoList{Id: data.Id, Status: data.Status, Task: data.Task})
	return controller.ToDo{Id: res.Id, Status: res.Status, Task: res.Task}
}

func (s Service) SvcGetAllData() ([]controller.ToDo, error) {
	utils.InfoLogger.Println(">> GetAllData")
	values, err := s.DataStore.AllTask()
	if err != nil {
		return []controller.ToDo{}, nil
	}
	var res []controller.ToDo
	for _, data := range values {
		res = append(res, controller.ToDo{Id: data.Id, Status: data.Status, Task: data.Task})
	}
	return res, nil
}

func (s Service) SvcGetDataById(id int) (controller.ToDo, error) {
	utils.InfoLogger.Println(">> GetDataById")

	res, err := s.DataStore.GetById(id)
	if err != nil {
		return controller.ToDo{}, err
	}

	return controller.ToDo{Id: res.Id, Status: res.Status, Task: res.Task}, nil

}

func (s Service) SvcRemoveTask(id int) (controller.ToDo, error) {
	utils.InfoLogger.Println(">> RemoveTask")
	res, err := s.DataStore.RemoveById(id)
	if err != nil {
		return controller.ToDo{}, err
	}
	return controller.ToDo{Id: res.Id, Status: res.Status, Task: res.Task}, nil

}

func (s Service) SvcUpdateTask(data controller.ToDo) (controller.ToDo, error) {
	utils.InfoLogger.Println(">> SvcUpdateTask")
	res, err := s.DataStore.UpdateTask(model.ToDoList{Id: data.Id, Status: data.Status, Task: data.Task})
	if err != nil {
		return controller.ToDo{}, err
	}
	return controller.ToDo{Id: res.Id, Status: res.Status, Task: res.Task}, nil

}
