package model

import (
	"ToDo/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDoList struct {
	Id     int    `json:"Id"`
	Task   string `json:"Task"`
	Status bool   `json:"Status"`
}

var Data = make(map[int]ToDoList)

type Storage struct {
	Client *mongo.Client
}

func Connect(ctx context.Context) Storage {
	databaseURI := os.Getenv("DATABASE_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURI))
	if err != nil {
		log.Fatal(fmt.Sprintf("error while connecting to MongoDB: %v", err))
	}
	log.Println("DataBase Connected")

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return Storage{Client: client}
}

func (s Storage) AddTask(newTask ToDoList) (ToDoList, error) {
	utils.InfoLogger.Println(">> AddTask")
	collection := s.Client.Database("ToDo").Collection("ToDos")
	bsonTaskList, err := bson.Marshal(newTask)
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return ToDoList{}, errors.New("Internal Server Error")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, bsonTaskList)
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return ToDoList{}, errors.New("Internal Server Error")
	}
	return newTask, nil

}

func (s Storage) AllTask() ([]ToDoList, error) {
	utils.InfoLogger.Println(">> AddTask")
	var res []ToDoList
	collection := s.Client.Database("ToDo").Collection("ToDos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return []ToDoList{}, errors.New("Internal Server Error")
	}
	if err = cursor.All(ctx, &res); err != nil {
		utils.ErrorLogger.Println(err.Error())
		return []ToDoList{}, errors.New("Internal Server Error")
	}
	return res, nil
}

func (s Storage) GetById(id int) (ToDoList, error) {
	utils.InfoLogger.Println(">> GetById")
	var res ToDoList
	collection := s.Client.Database("ToDo").Collection("ToDos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := make(chan error, 1)
	mssg := collection.FindOne(ctx, bson.M{"id": id}).Decode(&res)
	err <- mssg
	select {
	case <-ctx.Done():
		return ToDoList{}, errors.New("Internal Server Error")
	case mssg := <-err:
		if mssg != nil {
			utils.ErrorLogger.Println(mssg.Error())
			return ToDoList{}, errors.New(fmt.Sprintf("Task with ID %d Not Found", id))
		} else {
			return res, nil
		}

	}
}

func (s Storage) RemoveById(id int) (ToDoList, error) {
	utils.InfoLogger.Println(">> RemoveById")
	var res ToDoList
	collection := s.Client.Database("ToDo").Collection("ToDos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOneAndDelete(ctx, bson.M{"id": id}).Decode(&res)
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return ToDoList{}, errors.New("Internal Server Error")

	}
	return res, nil
}

func (s Storage) UpdateTask(newTask ToDoList) (ToDoList, error) {
	utils.InfoLogger.Println(">> UpdateTask")
	var res ToDoList
	collection := s.Client.Database("ToDo").Collection("ToDos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOneAndUpdate(ctx, bson.M{"id": newTask.Id}, bson.D{{"$set", bson.D{{"task", newTask.Task}, {"status", newTask.Status}}}}).Decode(&res)
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
		return ToDoList{}, errors.New("Internal Server Error")

	}
	return newTask, nil
}
