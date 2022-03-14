package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
)

type UsersModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Name        string             `json:"name,omitempty" binding:"required" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

func InitUserCollection() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_,err:=MONGO_COLLECTIONS.Users.Indexes().CreateMany(context.Background(), indexes, opts)
	if err!=nil{
		log.WithFields(log.Fields{
			"err":err,
		}).Errorln("error in creating index on UsersModel")
	}
}
