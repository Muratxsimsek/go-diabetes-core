package main

import (
	"context"
	//"errors"
	"fmt"
	"log"
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection Retrieves a client to the MongoDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	//username := "admin"//os.Getenv("MONGODB_USERNAME")
	//password := "bp8w3wEY6u8Ugws" //os.Getenv("MONGODB_PASSWORD")
	//clusterEndpoint := "mongodb+srv://admin:bp8w3wEY6u8Ugws@cluster0.z2zxm.mongodb.net/diabetes?retryWrites=true&w=majority" //os.Getenv("MONGODB_ENDPOINT")

	//connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:bp8w3wEY6u8Ugws@cluster0.z2zxm.mongodb.net/diabetes?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	//client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	//if err != nil {
	//	log.Printf("Failed to create client: %v", err)
	//}
	//
	//
	//if err != nil {
	//	log.Printf("Failed to connect to cluster: %v", err)
	//}

	// Force a connection to verify our connection string
	//err = client.Ping(ctx, nil)
	//if err != nil {
	//	log.Printf("Failed to ping cluster: %v", err)
	//}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

// GetAllTasks Retrives all tasks from the db
func GetAllDiabetes() ([]*Diabetes, error) {
	var diabetesList []*Diabetes

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database("diabetes").Collection("diabetes")

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	//collection := db.Collection("diabetes")
	//cursor, err := collection.Find(ctx, bson.D{})
	//if err != nil {
	//	return nil, err
	//}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &diabetesList)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	// Close the connection once no longer needed
	//err = client.Disconnect(context.TODO())
	//
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Connection to MongoDB closed.")
	//}

	return diabetesList, nil
}

//
//// GetTaskByID Retrives a task by its id from the db
//func GetTaskByID(id primitive.ObjectID) (*Task, error) {
//	var task *Task
//
//	client, ctx, cancel := getConnection()
//	defer cancel()
//	defer client.Disconnect(ctx)
//	db := client.Database("tasks")
//	collection := db.Collection("tasks")
//	result := collection.FindOne(ctx, bson.D{})
//	if result == nil {
//		return nil, errors.New("Could not find a Task")
//	}
//	err := result.Decode(&task)
//
//	if err != nil {
//		log.Printf("Failed marshalling %v", err)
//		return nil, err
//	}
//	log.Printf("Tasks: %v", task)
//	return task, nil
//}
//
////Create creating a task in a mongo
//func Create(task *Task) (primitive.ObjectID, error) {
//	client, ctx, cancel := getConnection()
//	defer cancel()
//	defer client.Disconnect(ctx)
//	task.ID = primitive.NewObjectID()
//
//	result, err := client.Database("tasks").Collection("tasks").InsertOne(ctx, task)
//	if err != nil {
//		log.Printf("Could not create Task: %v", err)
//		return primitive.NilObjectID, err
//	}
//	oid := result.InsertedID.(primitive.ObjectID)
//	return oid, nil
//}
//
////Update updating an existing task in a mongo
//func Update(task *Task) (*Task, error) {
//	var updatedTask *Task
//	client, ctx, cancel := getConnection()
//	defer cancel()
//	defer client.Disconnect(ctx)
//
//	update := bson.M{
//		"$set": task,
//	}
//
//	upsert := true
//	after := options.After
//	opt := options.FindOneAndUpdateOptions{
//		Upsert:         &upsert,
//		ReturnDocument: &after,
//	}
//
//	err := client.Database("tasks").Collection("tasks").FindOneAndUpdate(ctx, bson.M{"_id": task.ID}, update, &opt).Decode(&updatedTask)
//	if err != nil {
//		log.Printf("Could not save Task: %v", err)
//		return nil, err
//	}
//	return updatedTask, nil
//}
