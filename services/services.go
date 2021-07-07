package services

import (
	"ToDo/storage"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func checkDataValidation(body io.ReadCloser) Response {
	InfoLogger.Println(">> checkDataValidation")
	var data storage.ToDoList
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		ErrorLogger.Println(">> Invalid Data")
		return Response{Message: "Invalid Data", Status: 400}
	}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		ErrorLogger.Println(">> Invalid Data")
		return Response{Message: "Invalid Data", Status: 400}
	}
	return Response{Message: "OK", Body: data, Status: 200}

}

func SvcAddTask(r *http.Request) Responses {
	InfoLogger.Println(">> AddTask")
	var newTask storage.ToDoList
	var response Response
	checkResponse := checkDataValidation(r.Body)
	if !checkResponse.IsOk() {
		return Error{Status: 400, Message: "Bad Request,Invalid Data"}
	}
	newTask = checkResponse.Body.(storage.ToDoList)
	res, status := storage.AddTask(newTask)
	response.Status = status
	response.Message = "OK"
	response.Body = res
	return response
}

func SvcGetAllData(r *http.Request) Responses {
	InfoLogger.Println(">> GetAllData")
	values := storage.AllTask()
	return Response{Message: "OK", Body: values, Status: 200}
}

func SvcGetDataById(r *http.Request) Responses {
	InfoLogger.Println(">> GetDataById")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ErrorLogger.Println(">> Invalid ID")
		return Error{Message: "Invalid ID", Status: 400}
	}
	res, status := storage.GetById(id)
	if status == 404 {
		return Error{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: status}
	}
	return Response{Message: "OK", Body: res, Status: status}

}

func SvcRemoveTask(r *http.Request) Responses {
	InfoLogger.Println(">> RemoveTask")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ErrorLogger.Println(">> Invalid ID")
		return Error{Message: "Invalid ID", Status: 400}
	}
	res, status := storage.RemoveById(id)
	if status == 404 {
		return Error{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: status}
	}
	return Response{Message: "Deleted", Body: res, Status: status}

}

func SvcUpdateTask(r *http.Request) Responses {
	InfoLogger.Println(">> SvcUpdateTask")
	checkResponse := checkDataValidation(r.Body)
	var newTask storage.ToDoList
	if !checkResponse.IsOk() {
		return checkResponse
	}
	newTask = checkResponse.Body.(storage.ToDoList)
	res, status := storage.UpdateTask(newTask)
	if status == 404 {
		return Error{Message: "Data Does not Exists", Status: status}
	}
	return Response{Message: "Updated Successfully", Body: res, Status: status}

}
