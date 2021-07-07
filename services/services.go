package services

import (
	"ToDo/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Data = make(map[int]models.ToDoList)

func checkExistence(Id int) bool {
	InfoLogger.Println(">> checkExistence")
	_, keyExists := Data[Id]
	return keyExists
}

func checkDataValidation(body io.ReadCloser) models.Response {
	InfoLogger.Println(">> checkDataValidation")
	var data models.ToDoList
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		ErrorLogger.Println(">> Invalid Data")
		return models.Response{Message: "Invalid Data", Status: 400}
	}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		ErrorLogger.Println(">> Invalid Data")
		return models.Response{Message: "Invalid Data", Status: 400}
	}
	return models.Response{Message: "OK", Body: data, Status: 200}

}

func SvcAddTask(r *http.Request) models.Responses {
	InfoLogger.Println(">> AddTask")
	var response models.Response
	var newTask models.ToDoList
	checkResponse := checkDataValidation(r.Body)
	if !checkResponse.IsOk() {
		return models.Error{Status: 400, Message: "Bad Request,Invalid Data"}
	}
	newTask = checkResponse.Body.(models.ToDoList)
	if !checkExistence(newTask.Id) {
		Data[newTask.Id] = newTask
		response.Status = 201
		response.Message = "Created Successfully"
		response.Body = Data[newTask.Id]
		return response

	}
	response.Message = fmt.Sprintf("Task with ID %d Already Exists", newTask.Id)
	response.Body = Data[newTask.Id]
	response.Status = 200
	return response

}

func SvcGetAllData() models.Response {
	InfoLogger.Println(">> GetAllData")
	var values []models.ToDoList
	for _, val := range Data {
		values = append(values, val)
	}
	return models.Response{Message: "OK", Body: values, Status: 200}
}

func SvcGetDataById(r *http.Request) models.Responses {
	InfoLogger.Println(">> GetDataById")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ErrorLogger.Println(">> Invalid ID")
		return models.Error{Message: "Invalid ID", Status: 400}
	}
	if checkExistence(id) {
		return models.Response{Message: "OK", Body: Data[id], Status: 200}
	}
	return models.Error{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: 404}

}

func SvcRemoveTask(r *http.Request) models.Responses {
	InfoLogger.Println(">> RemoveTask")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ErrorLogger.Println(">> Invalid ID")
		return models.Error{Message: "Invalid ID", Status: 400}
	}
	if checkExistence(id) {
		instance := Data[id]
		delete(Data, id)
		return models.Response{Message: "Deleted", Body: instance, Status: 200}
	}
	return models.Error{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: 404}

}

func SvcUpdateTask(r *http.Request) models.Responses {
	InfoLogger.Println(">> SvcUpdateTask")
	checkResponse := checkDataValidation(r.Body)
	var response models.Response
	var newTask models.ToDoList
	if !checkResponse.IsOk() {
		return checkResponse
	}
	newTask = checkResponse.Body.(models.ToDoList)
	if !checkExistence(newTask.Id) {
		return models.Error{Status: 404, Message: "Data Does not Exists"}
	}
	Data[newTask.Id] = newTask
	response.Status = 200
	response.Body = Data[newTask.Id]
	response.Message = "Updated Successfully"
	return response

}
