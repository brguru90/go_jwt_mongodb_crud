package triggers

import (
	"context"
	"learn_go/src/database"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)


func deleteUserCache(_id string, ctx context.Context) {
	// Deletes the cache for the specified user by his ID
	_users_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___id="+_id+"___/api/user/*").Result()
	if err == nil {
		for _, key := range _users_keys {
			database.REDIS_DB_CONNECTION.Del(ctx, key)
			log.WithFields(log.Fields{
				"key": key,
			}).Debugln(">>>>>>>>>>>>>>>> Redis, " + key + " Removed")
		}
	}
}

func eraseAllUserPaginationCache(ctx context.Context)  {
	// erasing pagination caches
	_paginated_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___paginated*").Result()
	if err == nil {
		for _, key := range _paginated_keys {
			database.REDIS_DB_CONNECTION.Del(ctx, key)
			log.WithFields(log.Fields{
				"key": key,
			}).Debugln(">>>>>>>>>>>>>>>> Redis, users___paginated removed")
		}
	}	
}


func getUsersCount(ctx context.Context)  {
	count,err:=database.MONGO_COLLECTIONS.Users.CountDocuments(ctx,bson.M{})
	if err==nil{
		err2:=database.REDIS_DB_CONNECTION.Set(ctx,"users_count",count,time.Second*0).Err()
		if err2!=nil{
			log.WithFields(log.Fields{
				"errors":err2,
			}).Errorln("Error in setting user count to redis")
		}
	} else {
		log.WithFields(log.Fields{
			"errors":err,
		}).Errorln("Error in getting user count")
	}
}

func invalidateCache(_id string)  {
	ctx := context.Background()
	database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", "true", time.Second*0)

	eraseAllUserPaginationCache(ctx)
	deleteUserCache(_id,ctx)
	getUsersCount(ctx)

	database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
	log.Infoln("deleted users_update_in_progress")
}


func OnUserModification(_id string,userData database.UsersModel,operationType string)  {
	invalidateCache(_id)
}
