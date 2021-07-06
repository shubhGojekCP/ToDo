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
	InfoLogger.Println(">>")
	_, keyExists := Data[Id]
	return keyExists
}

func checkDataValidation(body io.ReadCloser) models.Response {
	InfoLogger.Println(">>")
	var data models.ToDoList
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return models.Response{Message: "Invalid Data", Status: 400}
	}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return models.Response{Message: "Invalid Data", Status: 400}
	}
	return models.Response{Message: "OK", Body: data, Status: 200}

}

func AddTask(r *http.Request) models.Responses {
	InfoLogger.Println(">>")
	var response models.Response
	var newTask models.ToDoList
	checkResponse := checkDataValidation(r.Body)
	if !checkResponse.IsOk() {
		response.Status = 400
		response.Message = "Bad Request,Invalid Data"
		return response
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

func GetAllData() models.Response {
	InfoLogger.Println(">>")
	var values []models.ToDoList
	for _, val := range Data {
		values = append(values, val)
	}
	return models.Response{Message: "OK", Body: values, Status: 200}
}

func GetDataById(r *http.Request) models.Responses {
	InfoLogger.Println(">>")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return models.Error{Message: "Invalid ID", Status: 400}
	}
	if checkExistence(id) {
		return models.Response{Message: "OK", Body: Data[id], Status: 200}
	}
	return models.Response{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: 404}

}

func RemoveTask(r *http.Request) models.Responses {
	InfoLogger.Println(">>")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return models.Error{Message: "Invalid ID", Status: 400}
	}
	if checkExistence(id) {
		instance := Data[id]
		delete(Data, id)
		return models.Response{Message: "Deleted", Body: instance, Status: 200}
	}
	return models.Response{Message: fmt.Sprintf("Task with ID %d Not Found", id), Status: 404}

}

func SvcUpdateTask(r *http.Request) models.Responses {
	InfoLogger.Println(">>")
	checkResponse := checkDataValidation(r.Body)
	var response models.Response
	var newTask models.ToDoList
	if !checkResponse.IsOk() {
		return checkResponse
	}
	newTask = checkResponse.Body.(models.ToDoList)
	if !checkExistence(newTask.Id) {
		response.Status = 404
		response.Message = "Data Does not Exists"
		return response
	}
	Data[newTask.Id] = newTask
	response.Status = 200
	response.Body = Data[newTask.Id]
	response.Message = "Updated Successfully"
	return response

}
