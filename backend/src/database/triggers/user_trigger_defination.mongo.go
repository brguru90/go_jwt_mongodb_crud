package triggers

import (
	"context"
	"learn_go/src/database"
	"strconv"
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

func eraseAllUserPaginationCache(ctx context.Context) {
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

func getUsersCount(ctx context.Context) {
	count, err := database.MONGO_COLLECTIONS.Users.CountDocuments(ctx, bson.M{})
	if err == nil {
		err2 := database.REDIS_DB_CONNECTION.Set(ctx, "users_count", count, time.Second*0).Err()
		if err2 != nil {
			log.WithFields(log.Fields{
				"errors": err2,
			}).Errorln("Error in setting user count to redis")
		}
	} else {
		log.WithFields(log.Fields{
			"errors": err,
		}).Errorln("Error in getting user count")
	}
}

func modifyCacheProgressStatus(operation string, ctx context.Context) {
	const max_users_update_in_progress_ttl=time.Minute*5
	
	users_update_in_progress, err := database.REDIS_DB_CONNECTION.Get(ctx, "users_update_in_progress").Result()
	if err == nil {
		users_update_in_progress_int, _ := strconv.ParseInt(users_update_in_progress, 10, 64)
		if operation=="delete"{
			users_update_in_progress_int--
			if users_update_in_progress_int==0{
				database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
				log.Debugln("deleted users_update_in_progress")				
			}
		} else{
			users_update_in_progress_int++
		}

		// log.WithFields(log.Fields{
		// 	"users_update_in_progress_int":users_update_in_progress_int,
		// }).Debugln("modifyCacheProgressStatus")

		if users_update_in_progress_int!=0{
			database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", strconv.FormatInt(users_update_in_progress_int, 10),max_users_update_in_progress_ttl )
		}
	} else {
		if operation != "delete" {
			database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", "1", max_users_update_in_progress_ttl)
		}
	}
}

func invalidateCache(_id string) {
	log.WithFields(log.Fields{
		"_id":_id,
	}).Debugln("invalidateCache....")
	ctx := context.Background()

	modifyCacheProgressStatus("insert",ctx)
	eraseAllUserPaginationCache(ctx)
	deleteUserCache(_id, ctx)
	modifyCacheProgressStatus("delete",ctx)

	getUsersCount(ctx)
}

func OnUserModification(_id string, userData database.UsersModel, operationType string) {
	go invalidateCache(_id)
}
