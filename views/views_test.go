package views

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	var mockedData = []struct {
		data     string
		status   int
		response string
	}{
		{`{"Id":1,"Task":"Running","Status":true}`, http.StatusCreated, `{"Message":"Created Successfully","Body":{"Id":1,"Task":"Running","Status":true},"Status":201}` + "\n"},     // Creating Task For the First Time
		{`{"Id":1,"Task":"Running","Status":true}`, http.StatusOK, `{"Message":"Task with ID 1 Already Exists","Body":{"Id":1,"Task":"Running","Status":true},"Status":200}` + "\n"}, // Creating Task With Same ID
		{`{"Id":1,"Task":"Running","Status":"true"}`, http.StatusBadRequest, `{"Message":"Bad Request,Invalid Data","Body":null,"Status":400}` + "\n"},                               // Creating Task With Invalid Data
	}

	for _, e := range mockedData {
		var mockedTask = []byte(e.data)
		req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

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
	}{
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/2`, http.StatusOK, `{"Message":"OK","Body":{"Id":2,"Task":"Running","Status":true},"Status":200}` + "\n"}, // Creating Task For the First Time
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/pi`, http.StatusBadRequest, `{"message":"Invalid ID","errorcode":400}` + "\n"},                            // Creating Task For the First Time
		{`{"Id":2,"Task":"Running","Status":true}`, `/api/task/3`, http.StatusNotFound, `{"Message":"Task with ID 3 Not Found","Body":null,"Status":404}` + "\n"},        // Creating Task For the First Time

	}

	for _, e := range mockedData {
		var mockedTask = []byte(e.data)
		req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

		handler.ServeHTTP(rr, req)

		req, err = http.NewRequest("GET", e.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/task/{id}", GetTaskById)

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
	}{
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/3`, http.StatusOK, `{"Message":"Deleted","Body":{"Id":3,"Task":"Running","Status":true},"Status":200}` + "\n"}, // Creating Task For the First Time
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/pi`, http.StatusBadRequest, `{"message":"Invalid ID","errorcode":400}` + "\n"},                                 // Creating Task For the First Time
		{`{"Id":3,"Task":"Running","Status":true}`, `/api/task/4`, http.StatusNotFound, `{"Message":"Task with ID 4 Not Found","Body":null,"Status":404}` + "\n"},             // Creating Task For the First Time

	}

	for _, e := range mockedData {
		var mockedTask = []byte(e.data)
		req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

		handler.ServeHTTP(rr, req)

		req, err = http.NewRequest("DELETE", e.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/task/{id}", DeleteTask)

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
	}{
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":4,"Task":"Swimming","Status":false}`, http.StatusOK, `{"Message":"Deleted","Body":{"Id":3,"Task":"Running","Status":true},"Status":200}` + "\n"}, // Creating Task For the First Time
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":4,"Task":"Swimming","Status":"false"}`, http.StatusBadRequest, `{"message":"Invalid ID","errorcode":400}` + "\n"},                                // Creating Task For the First Time
		{`{"Id":4,"Task":"Running","Status":true}`, `{"Id":5,"Task":"Swimming","Status":false}`, http.StatusNotFound, `{"Message":"Task with ID 4 Not Found","Body":null,"Status":404}` + "\n"},             // Creating Task For the First Time

	}
	for _, e := range mockedData {
		var mockedTask = []byte(e.data)
		req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

		handler.ServeHTTP(rr, req)
		mockedTask = []byte(e.updatedData)
		req, err = http.NewRequest("PUT", `/api/task/update`, bytes.NewBuffer(mockedTask))
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(UpdateTaskStatus)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, e.status, rr.Code)

	}

}

func TestGetAllTask(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllTask)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}
