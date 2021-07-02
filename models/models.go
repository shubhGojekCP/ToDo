package models

type ToDoList struct {
	Id     int    `json:"Id"`
	Task   string `json:"Task"`
	Status bool   `json:"Status"`
}
