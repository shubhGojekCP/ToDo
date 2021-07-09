package services

import (
	"ToDo/controller"
	"ToDo/model"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct {
	mockAddTask    func(data model.ToDoList) model.ToDoList
	mockGetById    func(id int) (model.ToDoList, error)
	mockRemoveById func(id int) (model.ToDoList, error)
	mockUpdateTask func(data model.ToDoList) (model.ToDoList, error)
	mockAllTask    func() ([]model.ToDoList, error)
}

func (m mockStorage) AddTask(data model.ToDoList) model.ToDoList {
	return m.mockAddTask(data)
}

func (m mockStorage) GetById(id int) (model.ToDoList, error) {
	return m.mockGetById(id)
}

func (m mockStorage) RemoveById(id int) (model.ToDoList, error) {
	return m.mockRemoveById(id)
}

func (m mockStorage) UpdateTask(data model.ToDoList) (model.ToDoList, error) {
	return m.mockUpdateTask(data)
}

func (m mockStorage) AllTask() ([]model.ToDoList, error) {
	return m.mockAllTask()
}

func TestSvcAddTask(t *testing.T) {
	s := Service{DataStore: mockStorage{mockAddTask: func(data model.ToDoList) model.ToDoList { return data }}}
	res := s.SvcAddTask(controller.ToDo{Id: 1, Status: true, Task: "Running"})
	assert.Equal(t, controller.ToDo{Id: 1, Status: true, Task: "Running"}, res)

}

func TestSvcGetAllData(t *testing.T) {
	s := Service{DataStore: mockStorage{mockAllTask: func() ([]model.ToDoList, error) { return []model.ToDoList{{Id: 1, Task: "Running", Status: true}}, nil }}}
	res, _ := s.SvcGetAllData()
	assert.Equal(t, []controller.ToDo{{Id: 1, Status: true, Task: "Running"}}, res)

}

func TestSvcGetDataById(t *testing.T) {
	var mockedData = []struct {
		id      int
		service Service
		err     error
		res     controller.ToDo
	}{
		{1, Service{DataStore: mockStorage{mockGetById: func(id int) (model.ToDoList, error) {
			return model.ToDoList{Id: id, Status: true, Task: "Running"}, nil
		}}}, nil, controller.ToDo{Id: 1, Task: "Running", Status: true}},
		{2, Service{DataStore: mockStorage{mockGetById: func(id int) (model.ToDoList, error) {
			return model.ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id))
		}}}, errors.New(fmt.Sprintf("Task with ID %d Not Found", 2)), controller.ToDo{}},
	}
	for _, rr := range mockedData {
		res, err := rr.service.SvcGetDataById(rr.id)
		assert.Equal(t, res, rr.res)
		assert.Equal(t, err, rr.err)
	}

}

func TestSvcRemoveTask(t *testing.T) {
	var mockedData = []struct {
		id      int
		service Service
		err     error
		res     controller.ToDo
	}{
		{1, Service{DataStore: mockStorage{mockRemoveById: func(id int) (model.ToDoList, error) {
			return model.ToDoList{Id: id, Status: true, Task: "Running"}, nil
		}}}, nil, controller.ToDo{Id: 1, Task: "Running", Status: true}},
		{2, Service{DataStore: mockStorage{mockRemoveById: func(id int) (model.ToDoList, error) {
			return model.ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id))
		}}}, errors.New(fmt.Sprintf("Task with ID %d Not Found", 2)), controller.ToDo{}},
	}
	for _, rr := range mockedData {
		res, err := rr.service.SvcRemoveTask(rr.id)
		assert.Equal(t, res, rr.res)
		assert.Equal(t, err, rr.err)
	}
}

func TestSvcUpdateTask(t *testing.T) {
	var mockedData = []struct {
		data    controller.ToDo
		service Service
		err     error
		res     controller.ToDo
	}{
		{controller.ToDo{Id: 1, Task: "Swimming", Status: true}, Service{DataStore: mockStorage{mockUpdateTask: func(data model.ToDoList) (model.ToDoList, error) {
			return data, nil
		}}}, nil, controller.ToDo{Id: 1, Task: "Swimming", Status: true}},
		{controller.ToDo{Id: 2, Task: "Jogging", Status: false}, Service{DataStore: mockStorage{mockUpdateTask: func(data model.ToDoList) (model.ToDoList, error) {
			return model.ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", data.Id))
		}}}, errors.New(fmt.Sprintf("Task with ID %d Not Found", 2)), controller.ToDo{}},
	}
	for _, rr := range mockedData {
		res, err := rr.service.SvcUpdateTask(rr.data)
		assert.Equal(t, res, rr.res)
		assert.Equal(t, err, rr.err)
	}
}
