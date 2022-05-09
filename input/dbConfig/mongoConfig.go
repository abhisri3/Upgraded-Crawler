package dbconfig

import (
	"context"
	model "crawl/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://abhisri344:learning_mongo@cluster0.5zlzm.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "InsuranceCrawler"
const urlCol = "InputData"
const dataCol = "InsuranceData"

var urlCollection *mongo.Collection
var dataCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	urlCollection = client.Database(dbName).Collection(urlCol)
	dataCollection = client.Database(dbName).Collection(dataCol)

}

func Insert(searchdata model.Search) {
	if searchdata.Container ==""{
		searchdata.Container ="body"
	}

	if len(searchdata.Patterns) == 0{
		searchdata.Patterns = []string{"insurance","plan","bima"}
	}
	now := time.Now()

	searchdata.Id = strconv.FormatInt(now.UnixNano(), 16)
	inserted, err := urlCollection.InsertOne(context.Background(), searchdata)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ", inserted.InsertedID)
}

func GetAllSeraches() []model.Search {
	cursor, err := urlCollection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var searches []model.Search

	for cursor.Next(context.Background()) {
		var search model.Search

		err := cursor.Decode(&search)

		if err != nil {
			log.Fatal(err)
		}
		searches = append(searches, search)
	}
	defer cursor.Close(context.Background())

	return searches
}

func ChangeStatus(Id string, status bool) {
	// id, err := primitive.ObjectIDFromHex(Id)

	// if err != nil{
	// 	log.Fatal(err)
	// }

	filter := bson.M{"_id": Id}
	update := bson.M{"$set": bson.M{"completed": status}}

	urlCollection.UpdateOne(context.Background(), filter, update)
}

func InsertData(insuranceData model.InsuranceData) {
	now := time.Now()

	insuranceData.Id = strconv.FormatInt(now.UnixNano(), 16)
	_, err := dataCollection.InsertOne(context.Background(), insuranceData)

	if err != nil {
		log.Fatal(err)
	}
}
