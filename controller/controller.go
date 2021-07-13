package controller

import (
	"ToDo/utils"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ToDo struct {
	Id     int
	Task   string
	Status bool
}

type response struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Body"`
}

type errorResponse struct {
	Message   string `json:"Message"`
	ErrorCode int    `json:"ErrorCode"`
}

type Handler struct {
	Service ToDoService
}

type ToDoService interface {
	SvcAddTask(req ToDo) (ToDo, error)
	SvcGetDataById(id int) (ToDo, error)
	SvcRemoveTask(id int) (ToDo, error)
	SvcUpdateTask(req ToDo) (ToDo, error)
	SvcGetAllData() ([]ToDo, error)
}

func requestBodyValidation(body io.Reader) (ToDo, error) {
	utils.InfoLogger.Println(">> requestBodyValidation")
	var data ToDo
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		utils.ErrorLogger.Println(">> Invalid Data")
		return ToDo{}, err
	}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		utils.ErrorLogger.Println(err.Error())
		return ToDo{}, errors.New("Invalid Data,Bad Request")
	}
	return data, nil

}

func checkParamValidation(params map[string]string) error {
	utils.InfoLogger.Println(">> checkParamValidation")
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return errors.New("Invalid ID")
	}
	return nil
}

// CreateTask creates task based on valid data provided,else return 400 for
// invalid Data,or 200 if Task with ID already exists.
func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> CreateTask")
	data, err := requestBodyValidation(r.Body)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	res, err := h.Service.SvcGetDataById(data.Id)
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response{Message: "OK", Data: res})
		return
	} else if err.Error() == "Internal Server Error" {
		errorResponseHandler(err, w)
		return
	}
	res, err = h.Service.SvcAddTask(data)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	successResponseHandler(201, res, w)

}

// GetTaskById gives the task based on Id provided,else return 400 for Invalid ID ,
// or 404 if Task with ID does not exists.
func (h Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> GetTaskById")
	params := mux.Vars(r)
	err := checkParamValidation(params)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	id, _ := strconv.Atoi(params["id"])
	res, err := h.Service.SvcGetDataById(id)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	successResponseHandler(200, res, w)

}

// DeleteTask deletes a task based on valid ID provided,else return 400 for
// invalid ID,or 404 if Task with ID does not exists.
func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> DeleteTask")
	params := mux.Vars(r)
	err := checkParamValidation(params)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	id, _ := strconv.Atoi(params["id"])
	res, err := h.Service.SvcGetDataById(id)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	res, err = h.Service.SvcRemoveTask(id)

	successResponseHandler(200, res, w)
}

// UpdateTaskStatus updates a task based on valid data provided,else return 400 for
// invalid data,or 404 if Task with ID does not exists.
func (h Handler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> UpdatetaskStatus")
	data, err := requestBodyValidation(r.Body)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	_, err = h.Service.SvcGetDataById(data.Id)
	if err != nil {
		if err.Error() == "Internal Server Error" {
			errorResponseHandler(err, w)
			return
		}
		errorResponseHandler(err, w)
		return
	}
	res, err := h.Service.SvcUpdateTask(data)
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	successResponseHandler(200, res, w)
}

//GetAllTask gets all the tasks.
func (h Handler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> GetAllTask")
	res, err := h.Service.SvcGetAllData()
	if err != nil {
		errorResponseHandler(err, w)
		return
	}
	successResponseHandler(200, res, w)

}

func errorResponseHandler(err error, w http.ResponseWriter) {
	switch err.Error() {
	case "Internal Server Error":
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
	case "Invalid ID":
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
	case "Invalid Data,Bad Request":
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
	default:
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 404})
	}

}

func successResponseHandler(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: body})
}
