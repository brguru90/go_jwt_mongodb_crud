package triggers

import (
	"context"
	"fmt"
	"learn_go/src/database"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TriggerForUsersModification() {

	updateStream, err := database.MONGO_COLLECTIONS.Users.Watch(context.Background(), mongo.Pipeline{})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panicln("Users: Error in watching mongo trigger")
	}

	defer updateStream.Close(context.TODO())
	for updateStream.Next(context.TODO()) {
		var data bson.M
		if err := updateStream.Decode(&data); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panicln("Users: Error in fetching data from mongo trigger stream")
		}
		fmt.Printf("%v\n", data)
	}
}
