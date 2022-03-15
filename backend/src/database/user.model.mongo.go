package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	Email       string             `json:"email,omitempty" binding:"required" bson:"email,omitempty"`
	Name        string             `json:"name,omitempty" binding:"required" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" binding:"required" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" swaggerignore:"true"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" swaggerignore:"true"`
}

func InitUserCollection() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{
					Key:   "_id",
					Value: 1,
				},
				{
					Key:   "email",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := MONGO_COLLECTIONS.Users.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorln("error in creating index on UsersModel")
	}
}
