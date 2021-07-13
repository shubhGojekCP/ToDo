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
	Message string `json:"Message"`
	Data    ToDo   `json:"Body"`
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

func checkDataValidation(body io.Reader) (ToDo, error) {
	utils.InfoLogger.Println(">> checkDataValidation")
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
	data, err := checkDataValidation(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
		return
	}
	res, err := h.Service.SvcGetDataById(data.Id)
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response{Message: "OK", Data: res})
		return
	}
	res, err = h.Service.SvcAddTask(data)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: res})

}

// GetTaskById gives the task based on Id provided,else return 400 for Invalid ID ,
// or 404 if Task with ID does not exists.
func (h Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> GetTaskById")
	params := mux.Vars(r)
	err := checkParamValidation(params)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
		return
	}
	id, _ := strconv.Atoi(params["id"])
	res, err := h.Service.SvcGetDataById(id)
	if err != nil {
		if err.Error() == "Internal Server Error" {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
			return
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 404})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: res})

}

// DeleteTask deletes a task based on valid ID provided,else return 400 for
// invalid ID,or 404 if Task with ID does not exists.
func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> DeleteTask")
	params := mux.Vars(r)
	err := checkParamValidation(params)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
		return
	}
	id, _ := strconv.Atoi(params["id"])
	res, err := h.Service.SvcGetDataById(id)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 404})
		return
	}
	res, err = h.Service.SvcRemoveTask(id)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: res})
}

// UpdateTaskStatus updates a task based on valid data provided,else return 400 for
// invalid data,or 404 if Task with ID does not exists.
func (h Handler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> UpdatetaskStatus")
	data, err := checkDataValidation(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 400})
		return
	}
	_, err = h.Service.SvcGetDataById(data.Id)
	if err != nil {
		if err.Error() == "Internal Server Error" {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
			return
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 404})
		return
	}
	res, err := h.Service.SvcUpdateTask(data)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: res})
}

//GetAllTask gets all the tasks.
func (h Handler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println(">> GetAllTask")
	type response struct {
		Message string `json:"Message"`
		Data    []ToDo `json:"Data"`
	}
	res, err := h.Service.SvcGetAllData()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error(), ErrorCode: 500})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response{Message: "OK", Data: res})

}
