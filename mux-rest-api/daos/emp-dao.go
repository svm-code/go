package daos

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var RecordNotFoundError = errors.New("Record Not Found!")

type EmployeeModel struct {
	Name    string `json:"name"`
	DeptId  int    `json:"deptId"`
	Phone   int64  `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Employee struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Name    string             `bson:"name" json:"name"`
	DeptId  int                `bson:"deptId" json:"deptId"`
	Phone   int64              `bson:"phone" json:"phone"`
	Email   string             `bson:"email" json:"email"`
	Address string             `bson:"address" json:"address"`
}

func (Employee) GetCollectionName() string {
	return "employees"
}

func (emp Employee) Save() error {
	f := func(collection *mongo.Collection, ctx context.Context) error {
		_, err := collection.InsertOne(ctx, emp)
		return err
	}
	err := DoSome(f, emp)
	return err
}

func (emp Employee) Update() error {
	f := func(collection *mongo.Collection, ctx context.Context) error {
		id := primitive.ObjectID(emp.ID)
		filter := bson.D{{"_id", id}}
		//update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}

		var setElements bson.D

		if emp.Name != "" {
			setElements = append(setElements, bson.E{"name", emp.Name})
		}
		if emp.Phone != 0 {
			setElements = append(setElements, bson.E{"phone", emp.Phone})
		}
		if emp.DeptId != 0 {
			setElements = append(setElements, bson.E{"deptId", emp.DeptId})
		}
		if emp.Email != "" {
			setElements = append(setElements, bson.E{"email", emp.Email})
		}
		if emp.Address != "" {
			setElements = append(setElements, bson.E{"address", emp.Address})
		}

		update := bson.D{{"$set", setElements}}

		result, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}
		if result.ModifiedCount == 0 {
			return RecordNotFoundError
		}
		log.Println("Update Result:", result.ModifiedCount)
		return err
	}
	return DoSome(f, Employee{})
}

func (Employee) Delete(id string) error {
	f := func(collection *mongo.Collection, ctx context.Context) error {
		_id, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{"_id", _id}}
		result, err := collection.DeleteOne(ctx, filter)
		if result.DeletedCount == 0 {
			return RecordNotFoundError
		}
		log.Println("Deleted Result:", result.DeletedCount)
		return err
	}
	err := DoSome(f, Employee{})
	return err
}

func (Employee) FindAll() []*Employee {
	var res []*Employee
	f := func(collection *mongo.Collection, ctx context.Context) error {

		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			return err
		}

		for cur.Next(ctx) {
			var emp Employee
			err := cur.Decode(&emp)
			if err != nil {
				return err
			}
			res = append(res, &emp)
		}

		if err := cur.Err(); err != nil {
			return err
		}
		return nil
	}
	err := DoSome(f, Employee{})
	if err == nil {
		log.Println("Selected records counts: ", len(res))
	}
	return res
}

func (Employee) FindByName(name string) (Employee, error) {
	var res Employee
	f := func(collection *mongo.Collection, ctx context.Context) error {
		filter := bson.D{{"name", name}}
		return collection.FindOne(context.Background(), filter).Decode(&res)
	}
	err := DoSome(f, Employee{})
	if err == nil {
		log.Println("Selected record:", res)
	}
	return res, err
}
