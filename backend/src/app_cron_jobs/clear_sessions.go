package app_cron_jobs

import (
	"context"
	"learn_go/src/database"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func ClearExpiredToken() {
	// mongo connection is a concurrency-safe ignore --race errors

	ctx := context.Background()
	_time := time.Now()

	log.WithFields(log.Fields{
		"time": _time,
	}).Debug(" -- ClearExpiredToken Cron job started -- ")

	result,err:=database.MONGO_COLLECTIONS.ActiveSessions.DeleteMany(ctx,bson.M{
		"exp":bson.M{"$lte":_time.UnixMilli()},
	})

	if err!=nil{
		log.WithFields(log.Fields{
			"err": err,
		}).Errorln("Failed to delete user data")
		return
	}

	log.WithFields(log.Fields{
		"Total_cleared_session_entry": result.DeletedCount,
	}).Debug(" -- ClearExpiredToken Cron finished -- ")
}
