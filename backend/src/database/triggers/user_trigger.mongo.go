package triggers

import (
	"context"
	"learn_go/src/database"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TriggerForUsersModification() {

	database.REDIS_DB_CONNECTION.Del(context.TODO(), "users_update_in_progress")

	updateStream, err := database.MONGO_COLLECTIONS.Users.Watch(context.Background(), mongo.Pipeline{})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panicln("Users: Error in watching mongo trigger")
	}

	// map[_id:map[_data:82623058D5000000012B022C0100296E5A10040E329BFE5CDA40428B922AE46B800E5C46645F69640064623058D5B609E0D80012591A0004] clusterTime:{1647335637 1} documentKey:map[_id:ObjectID("623058d5b609e0d80012591a")] fullDocument:map[_id:ObjectID("623058d5b609e0d80012591a") createdat:1647335637535 description:string email:asacasa name:strsaacing updatedat:1647335637535] ns:map[coll:users db:jwt4] operationType:insert]

	// map[_id:map[_data:8262305C48000000012B022C0100296E5A10040E329BFE5CDA40428B922AE46B800E5C46645F6964006462305C30EC812ED347E2FE9F0004] clusterTime:{1647336520 1} documentKey:map[_id:ObjectID("62305c30ec812ed347e2fe9f")] ns:map[coll:users db:jwt4] operationType:update updateDescription:map[removedFields:[] truncatedArrays:[] updatedFields:map[description:dscs email:stsdcsdring name:sdc updatedAt:1647336520406]]]

	// map[_id:map[_data:8262305CED000000012B022C0100296E5A10040E329BFE5CDA40428B922AE46B800E5C46645F6964006462305C30EC812ED347E2FE9F0004] clusterTime:{1647336685 1} documentKey:map[_id:ObjectID("62305c30ec812ed347e2fe9f")] ns:map[coll:users db:jwt4] operationType:delete]

	defer updateStream.Close(context.TODO())
	for updateStream.Next(context.TODO()) {
		var data bson.M
		if err := updateStream.Decode(&data); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panicln("Users: Error in fetching data from mongo trigger stream")
		}
		// log.Warnln(data)
		// operationType: insert,delete,update
		if data["fullDocument"] != "" {
			var userData database.UsersModel
			document_key := data["documentKey"].(bson.M)["_id"].(primitive.ObjectID)

			switch data["operationType"] {
			case "insert":
				{
					bsonBytes, _ := bson.Marshal(data["fullDocument"])
					bson.Unmarshal(bsonBytes, &userData)
					OnUserModification(document_key.Hex(), userData, data["operationType"].(string))
				}
			case "update":
				{
					bsonBytes, _ := bson.Marshal(data["updatedFields"])
					bson.Unmarshal(bsonBytes, &userData)
					userData.ID = document_key
					OnUserModification(document_key.Hex(), userData, data["operationType"].(string))
				}
			case "delete":
				{

				}
			default:
				log.WithFields(log.Fields{
					"operationType": data["operationType"],
					"data":          data,
				}).Warnln("Unhandled operation")
			}

			// log.Infoln(userData)
		}

	}

	log.Error("TriggerForUsersModification finished...")
}
