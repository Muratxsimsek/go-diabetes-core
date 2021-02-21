package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:bp8w3wEY6u8Ugws@cluster0.z2zxm.mongodb.net/diabetes?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func GetAllDiabetes() ([]*Diabetes, error) {
	var diabetesList []*Diabetes

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database("diabetes").Collection("diabetes")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"sugarDate", -1}})
	//findOptions.SetLimit(2)

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

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

func GetDiabetesByID(id string) (*Diabetes, error) {
	var diabetes *Diabetes

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("diabetes")
	collection := db.Collection("diabetes")
	docID, _ := primitive.ObjectIDFromHex(id)
	//fmt.Println(`bson.M{"_id"": docID}:`, bson.M{"_id": docID})
	//filter := bson.M{"_id": docID}
	//filter := bson.D{primitive.E{Key: "_id", Value: docID}}
	//filter := bson.M{
	//	"_id": docID,
	//}
	//err := collection.FindOne(ctx, filter).Decode(&diabetes)
	err := collection.FindOne(ctx, bson.D{{"_id", docID}}).Decode(&diabetes)
	if err != nil {
		panic(err)
	}

	log.Printf("Diabetes: %v", diabetes)
	return diabetes, nil
}

func CreateDiabetes(diabetes *Diabetes) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	diabetes.ID = primitive.NewObjectID()

	result, err := client.Database("diabetes").Collection("diabetes").InsertOne(ctx, diabetes)
	if err != nil {
		log.Printf("Could not create Diabetes: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func UpdateDiabetes(id string, diabetes *Diabetes) (interface{}, error) {

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	docID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{
		"$set": bson.M{
			"sugarValue":   diabetes.SugarValue,
			"hungerStatus": diabetes.HungerStatus,
			"sugarDate":    diabetes.SugarDate,
		},
	}

	//update := bson.M{
	//	"$set" : diabetes,
	//}

	var updatedDiabetes *Diabetes

	err := client.Database("diabetes").Collection("diabetes").
		FindOneAndUpdate(ctx, bson.D{{"_id", docID}}, update).Decode(&updatedDiabetes)

	if err != nil {
		panic(err)
	}

	return updatedDiabetes.ID, nil
}

func GetDiabetesChart() (*DiabetesChart, error) {
	var diabetesChart DiabetesChart
	var diabetesList []*Diabetes

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database("diabetes").Collection("diabetes")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"sugarDate", -1}})
	//findOptions.SetLimit(2)

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	//for cursor.Next(ctx) {
	//	var row bson.M
	//	if err = cursor.Decode(&row); err != nil {
	//		log.Fatal(err)
	//	}
	//	//fmt.Println(row)
	//
	//}

	err = cursor.All(ctx, &diabetesList)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	for _, row := range diabetesList {
		diabetesChart.SugarValues = append(diabetesChart.SugarValues, row.SugarValue)
		diabetesChart.Dates = append(diabetesChart.Dates, row.SugarDate.UnixNano()/1000000)
		//log.Println(row)
	}

	min, max := FindMaxAndMin(diabetesChart.SugarValues)
	diabetesChart.MinSugarValue = min
	diabetesChart.MaxSugarValue = max
	//log.Println(time.Now().UnixNano() / 1000000)
	return &diabetesChart, nil
}

func FindMaxAndMin(list []int16) (min int16, max int16) {
	min = list[0]
	max = list[0]
	for _, el := range list {
		if el > max {
			max = el
		}
		if el < min {
			min = el
		}
	}
	return min, max
}
