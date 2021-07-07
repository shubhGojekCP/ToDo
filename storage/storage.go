package storage

type ToDoList struct {
	Id     int    `json:"Id"`
	Task   string `json:"Task"`
	Status bool   `json:"Status"`
}

var Data = make(map[int]ToDoList)

func checkExistence(Id int) bool {
	_, keyExists := Data[Id]
	return keyExists
}

func AddTask(newTask ToDoList) (ToDoList, int) {
	if !checkExistence(newTask.Id) {
		Data[newTask.Id] = newTask
		return Data[newTask.Id], 201

	}
	return Data[newTask.Id], 200
}

func AllTask() []ToDoList {
	var values []ToDoList
	for _, val := range Data {
		values = append(values, val)
	}
	return values
}

func GetById(id int) (interface{}, int) {
	if checkExistence(id) {
		return Data[id], 200
	}
	return nil, 404
}

func RemoveById(id int) (interface{}, int) {
	if checkExistence(id) {
		instance := Data[id]
		delete(Data, id)
		return instance, 200
	}
	return nil, 404
}

func UpdateTask(newTask ToDoList) (interface{}, int) {
	if !checkExistence(newTask.Id) {
		return nil, 404
	}
	Data[newTask.Id] = newTask
	return Data[newTask.Id], 200
}
