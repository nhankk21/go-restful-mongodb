package data

import (
	"context"
	"fmt"
	"go-restful/pkg/dto"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var dataBase string = "test2"
var col string = "todo"

func CloseDb() {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func ConnectDb(uri string) {
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

}

func GetDataByID(id int64) (dto.Todo, error) {

	collection := client.Database(dataBase).Collection(col)
	size, err := collection.CountDocuments(ctx, bson.D{})
	if id > size {
		return dto.Todo{}, err
	}
	cursor, err := collection.Find(ctx, bson.D{{"id", bson.D{{"$eq", id}}}}, options.Find().SetProjection(bson.D{{"_id", 0}}))
	if err != nil {
		return dto.Todo{}, err
	}
	var bsonResults []bson.D
	err = cursor.All(ctx, &bsonResults)
	fmt.Println("Get by id oke")
	if err != nil {
		return dto.Todo{}, err
	}
	bsonBytes, _ := bson.Marshal(bsonResults[len(bsonResults)-1])
	var result dto.Todo
	bson.Unmarshal(bsonBytes, &result)
	return result, nil

}

func GetAllData() ([]dto.Todo, error) {
	collection := client.Database(dataBase).Collection(col)
	cursor, err := collection.Find(ctx, bson.D{
		// {"_id", 0}
	}, options.Find().SetProjection(bson.D{{"_id", 0}}))
	if err != nil {
		return nil, err
	}
	var bsonResults []bson.D
	results := make([]dto.Todo, 0)
	err = cursor.All(ctx, &bsonResults)
	for _, item := range bsonResults {
		bsonBytes, _ := bson.Marshal(item)
		var result dto.Todo
		bson.Unmarshal(bsonBytes, &result)
		results = append(results, result)
	}

	fmt.Println("Get all oke")
	fmt.Println(results)
	if err != nil {
		return nil, err
	}
	return results, err
}
func CreateData(todo dto.Todo) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	id, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	result, err := collection.InsertOne(ctx, dto.Todo{
		ID:      id + 1,
		Name:    todo.Name,
		Content: todo.Content,
		Status:  "ok",
	})
	fmt.Println("Insert oke")
	return result, err
}

func UpdateDataByID(id int, update interface{}) (*mongo.UpdateResult, error) {
	collection := client.Database(dataBase).Collection(col)
	filter := bson.D{{"id", bson.D{{"$eq", id}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(result)
	return result, err
}
func DeleteDataByID(id int) (*mongo.DeleteResult, error) {
	collection := client.Database(dataBase).Collection(col)
	filter := bson.D{{"id", bson.D{{"$eq", id}}}}
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println(result)
	return result, err
}
