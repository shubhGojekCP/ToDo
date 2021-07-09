package controller

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type mockedService struct {
	mockSvcAddTask     func(req ToDo) ToDo
	mockSvcGetDataById func(id int) (ToDo, error)
	mockSvcRemoveTask  func(id int) (ToDo, error)
	mockSvcUpdateTask  func(req ToDo) (ToDo, error)
	mockSvcGetAllData  func() ([]ToDo, error)
}

func (m mockedService) SvcAddTask(req ToDo) ToDo {
	return m.mockSvcAddTask(req)
}

func (m mockedService) SvcGetDataById(id int) (ToDo, error) {
	return m.mockSvcGetDataById(id)
}

func (m mockedService) SvcRemoveTask(id int) (ToDo, error) {
	return m.mockSvcRemoveTask(id)
}

func (m mockedService) SvcUpdateTask(req ToDo) (ToDo, error) {
	return m.mockSvcUpdateTask(req)
}
func (m mockedService) SvcGetAllData() ([]ToDo, error) {
	return m.mockSvcGetAllData()
}
func TestCreateTask(t *testing.T) {

	var mockedData = []struct {
		data     string
		status   int
		response string
		handler  Handler
	}{ // Creating Task For the First Time
		{`{"Id":1,"Task":"Running","Status":true}`, http.StatusCreated, `{"Message":"OK","Body":{"Id":1,"Task":"Running","Status":true}}` + "\n",
			Handler{Service: mockedService{mockSvcAddTask: func(req ToDo) ToDo { return req },
				mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{}, errors.New(fmt.Sprintf("Task with ID %d not Found", id)) }}}},

		// Creating Task With Same ID
		{`{"Id":1,"Task":"Running","Status":true}`, http.StatusOK, `{"Message":"OK","Body":{"Id":1,"Task":"Running","Status":true}}` + "\n",
			Handler{Service: mockedService{mockSvcAddTask: func(req ToDo) ToDo { return req },
				mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{Id: 1, Task: "Running", Status: true}, nil }}}},

		// Creating Task With Invalid Data
		{`{"Id":1,"Task":"Running","Status":"true"}`, http.StatusBadRequest, `{"Message":"Invalid Data,Bad Request","ErrorCode":400}` + "\n",
			Handler{}},
	}

	for _, e := range mockedData {
		var mockedTask = []byte(e.data)
		req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler.CreateTask)

		handler.ServeHTTP(rr, req)
		status := rr.Code
		assert.Equal(t, e.status, status)
		assert.Equal(t, e.response, rr.Body.String())

	}

}

func TestGetTaskById(t *testing.T) {
	var mockedData = []struct {
		data     string
		url      string
		status   int
		response string
		handler  Handler
	}{
		// Geting Task For the Valid and Existing ID
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/2`, http.StatusOK, `{"Message":"OK","Body":{"Id":2,"Task":"Running","Status":true}}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{Id: 2, Task: "Running", Status: true}, nil }}}},

		// Getting Task For the Invalid ID
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/pi`, http.StatusBadRequest, `{"Message":"Invalid ID","ErrorCode":400}` + "\n",
			Handler{}},

		// Getting Task For The Non Existent ID
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/3`, http.StatusNotFound, `{"Message":"Task with ID 3 Not Found","ErrorCode":404}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id)) }}}},
	}

	for _, e := range mockedData {
		req, err := http.NewRequest("GET", e.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/task/{id}", e.handler.GetTaskById)

		router.ServeHTTP(rr, req)
		status := rr.Code
		assert.Equal(t, e.status, status)
		assert.Equal(t, e.response, rr.Body.String())

	}

}

func TestDeleteTask(t *testing.T) {
	var mockedData = []struct {
		data     string
		url      string
		status   int
		response string
		handler  Handler
	}{
		// Deleting Task for Valid and Existing ID
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/3`, http.StatusOK, `{"Message":"OK","Body":{"Id":3,"Task":"Running","Status":true}}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{Id: 3, Task: "Running", Status: true}, nil },
				mockSvcRemoveTask: func(id int) (ToDo, error) { return ToDo{Id: 3, Task: "Running", Status: true}, nil }}}},

		// Deleting Task For the Invalid ID
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/pi`, http.StatusBadRequest, `{"Message":"Invalid ID","ErrorCode":400}` + "\n",
			Handler{}},

		// Deleting Task For The Non Existent ID
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/4`, http.StatusNotFound, `{"Message":"Task with ID 4 Not Found","ErrorCode":404}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id)) }}}},
	}

	for _, e := range mockedData {
		req, err := http.NewRequest("DELETE", e.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/task/{id}", e.handler.DeleteTask)

		router.ServeHTTP(rr, req)
		assert.Equal(t, e.status, rr.Code)
		assert.Equal(t, e.response, rr.Body.String())
	}

}

func TestUpdateTask(t *testing.T) {
	var mockedData = []struct {
		data        string
		updatedData string
		status      int
		response    string
		handler     Handler
	}{
		// Updating Task for Valid and Existing ID and Data
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":4,"Task":"Swimming","Status":false}`, http.StatusOK, `{"Message":"Deleted","Body":{"Id":3,"Task":"Running","Status":true}}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{Id: id, Status: true, Task: "Running"}, nil },
				mockSvcUpdateTask: func(data ToDo) (ToDo, error) { return ToDo{Id: 4, Task: "Swimming", Status: false}, nil }}}},

		// Updating Task For the Invalid Data
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":4,"Task":"Swimming","Status":"false"}`, http.StatusBadRequest, `{"Message":"Invalid ID","ErrorCode":400}` + "\n",
			Handler{Service: mockedService{}}},

		// Updating Task For The Non Existent ID
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":5,"Task":"Swimming","Status":false}`, http.StatusNotFound, `{"Message":"Task with ID 4 Not Found","ErrorCode":404}` + "\n",
			Handler{Service: mockedService{mockSvcGetDataById: func(id int) (ToDo, error) { return ToDo{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id)) }}}},
	}
	for _, e := range mockedData {
		var mockedTask = []byte(e.updatedData)
		req, err := http.NewRequest("PUT", `/api/task/update`, bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler.UpdateTaskStatus)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, e.status, rr.Code)

	}

}

func TestGetAllTask(t *testing.T) {
	mockHandler := Handler{Service: mockedService{mockSvcGetAllData: func() ([]ToDo, error) { return []ToDo{{1, "Running", true}}, nil }}}
	req, err := http.NewRequest("GET", "/api/task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetAllTask)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"Message":"OK","Data":[{"Id":1,"Task":"Running","Status":true}]}`+"\n", rr.Body.String())

}
