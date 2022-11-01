package main

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(ctx)

	//mongodb1 := client.Database("go-test")
	//employeeCollection := mongodb1.Collection("employee")

	emp, _ := FindById("1")

	str, _ := json.Marshal(emp)
	log.Println("str: ", string(str))

	/*
		result, err1 := employeeCollection.InsertOne(ctx, emp)
		log.Println("result", result)
		log.Println("Eorror", err)
		if err1 != nil {
			log.Println(err)
		}
		log.Println("result..........")
	*/
}

type Employee struct {
	Id      string `bson:"id" json:"id"`
	Name    string `bson:"name" json:"name"`
	DeptId  int    `bson:"deptId" json:"deptId"`
	Phone   int64  `bson:"phone" json:"phone"`
	Email   string `bson:"email" json:"email"`
	Address string `bson:"address" json:"email"`
}

var employees = []Employee{
	Employee{
		Id:      "1",
		Name:    "Suraj",
		DeptId:  1,
		Phone:   9527957483,
		Email:   "suraj.more@github.com",
		Address: "washim",
	},
	Employee{
		Id:      "2",
		Name:    "Sanket",
		DeptId:  2,
		Phone:   9899799697,
		Email:   "sanket.kasar@github.com",
		Address: "Sangamner",
	},
}

func FindAll() []Employee {
	return employees
}

func FindById(id string) (Employee, error) {
	for _, emp := range employees {
		if emp.Id == id {
			return emp, nil
		}
	}
	return Employee{}, errors.New("Employee not found for ID : " + id)
}
