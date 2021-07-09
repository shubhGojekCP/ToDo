package model

import (
	"ToDo/utils"
	"errors"
	"fmt"
)

type ToDoList struct {
	Id     int    `json:"Id"`
	Task   string `json:"Task"`
	Status bool   `json:"Status"`
}

var Data = make(map[int]ToDoList)

type Storage struct {
}

func checkExistence(Id int) bool {
	_, keyExists := Data[Id]
	return keyExists
}

func (s Storage) AddTask(newTask ToDoList) ToDoList {
	utils.InfoLogger.Println(">> AddTask")

	Data[newTask.Id] = newTask
	return Data[newTask.Id]

}

func (s Storage) AllTask() ([]ToDoList, error) {
	utils.InfoLogger.Println(">> AddTask")
	var values []ToDoList
	for _, val := range Data {
		values = append(values, val)
	}
	return values, nil
}

func (s Storage) GetById(id int) (ToDoList, error) {
	utils.InfoLogger.Println(">> GetById")
	if checkExistence(id) {
		return Data[id], nil
	}
	return ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id))
}

func (s Storage) RemoveById(id int) (ToDoList, error) {
	utils.InfoLogger.Println(">> RemoveById")
	if checkExistence(id) {
		instance := Data[id]
		delete(Data, id)
		return instance, nil
	}
	return ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id))
}

func (s Storage) UpdateTask(newTask ToDoList) (ToDoList, error) {
	utils.InfoLogger.Println(">> UpdateTask")
	if !checkExistence(newTask.Id) {
		return ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", newTask.Id))
	}
	Data[newTask.Id] = newTask
	return Data[newTask.Id], nil
}
