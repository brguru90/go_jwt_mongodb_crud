package database

import (
	"context"
	"fmt"
	"learn_go/src/configs"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_DB_CONNECTION *mongo.Client
var MONGO_DATABASE *mongo.Database

type MongoCollections struct {
	Users          *mongo.Collection
	ActiveSessions *mongo.Collection
}

var MONGO_COLLECTIONS MongoCollections

func InitMongoDB() {

	var DB_USER string = configs.EnvConfigs.MONGO_DB_USER
	var DB_PASSWORD string = configs.EnvConfigs.MONGO_DB_PASSWORD
	var DB_HOST string = configs.EnvConfigs.MONGO_DB_HOST
	var DATABASE string = configs.EnvConfigs.MONGO_DATABASE
	var DB_PORT int64 = configs.EnvConfigs.MONGO_DB_PORT
	
	mongo_url := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", DB_USER, DB_PASSWORD,DB_HOST, DB_PORT)

	log.Infoln(mongo_url)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error

	clientOptions := options.Client().ApplyURI(mongo_url)
	clientOptions = clientOptions.SetMaxPoolSize(100)
	MONGO_DB_CONNECTION, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"DB_URL": mongo_url,
		}).Error("Unable to connect to database ==>")
	}

	MONGO_DATABASE = MONGO_DB_CONNECTION.Database(DATABASE)
	MONGO_COLLECTIONS.Users = MONGO_DATABASE.Collection("users")
	MONGO_COLLECTIONS.ActiveSessions = MONGO_DATABASE.Collection("active_sessions")

	InitUserCollection()
	InitActiveSessionCollection()
	createSampleCollection()
}

func createSampleCollection() {
	collection_name := "sample_collection"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	command := bson.D{{"create", collection_name}}
	var result bson.M = bson.M{
		"bsonType": "object",
		"required": []string{"endpointID", "ip", "port", "lastHeartbeatDate"},
		"properties": bson.M{
			"endpointID": bson.M{
				"bsonType":    "double",
				"description": "the endpoint Hash",
			},
			"ip": bson.M{
				"bsonType":    "string",
				"description": "the endpoint IP address",
			},
			"port": bson.M{
				"bsonType":    "int",
				"maximum":     65535,
				"description": "the endpoint Port",
			},
			"lastHeartbeatDate": bson.M{
				"bsonType":    "date",
				"description": "the last time when the heartbeat has been received",
			},
		},
	}
	if err := MONGO_DATABASE.RunCommand(ctx, command).Decode(&result); err != nil {
		log.Errorln(err)
	}
}
