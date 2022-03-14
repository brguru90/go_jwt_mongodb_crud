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

type ActiveSessionsModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TokenID   string             `json:"token_id,omitempty" bson:"token_id,omitempty"`
	UA        string             `json:"ua,omitempty" bson:"ua,omitempty"`
	IP        string             `json:"ip,omitempty" bson:"ip,omitempty"`
	Exp       int64              `json:"exp,omitempty" bson:"exp,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty"`
}

func InitActiveSessionCollection() {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{
				"token_id": 1,
			},
		},
		{
			Keys: bson.M{
				"user_id": 1,
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "_id",
					Value: 1,
				},
				{
					Key:   "token_id",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := MONGO_COLLECTIONS.ActiveSessions.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorln("error in creating index on ActiveSessionsModel")
	}
}
