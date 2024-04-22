package controller

import (
	"LoginAPI/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Pranjal:Pranjal%40123@cluster0.sc7ucqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "LoginBucket"
const colName = "LoginUser"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
}
func UpdateData(number string) int64 {
	filter := bson.M{"usernumber": number}
	update := bson.M{"$set": bson.M{"isalreadyexit": true}}
	updateValue, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateValue.ModifiedCount

}
func CheckAlreadyExit(number string) bool {
	filter := bson.D{{Key: "usernumber", Value: number}}
	var result bson.D
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return err != mongo.ErrNoDocuments
}
func InsertIntoDB(loginData model.LoginData) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.Background(), loginData)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
