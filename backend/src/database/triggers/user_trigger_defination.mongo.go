package triggers

import (
	"context"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// func deleteUserCache(_id string, ctx context.Context) {
// 	// Deletes the cache for the specified user by his ID
// 	_users_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___id="+_id+"___/api/user/*").Result()
// 	database.REDIS_DB_CONNECTION.
// 	if err == nil {
// 		for _, key := range _users_keys {
// 			database.REDIS_DB_CONNECTION.Del(ctx, key)
// 			log.WithFields(log.Fields{
// 				"key": key,
// 			}).Debugln(">>>>>>>>>>>>>>>> Redis, " + key + " Removed")
// 		}
// 	}
// }

func deleteUserCache(_id string, ctx context.Context) {
	// Deletes the cache for the specified user by his ID
	database.REDIS_DB_CONNECTION.Del(ctx, "users___id="+_id)
}

// func eraseAllUserPaginationCache(ctx context.Context) {
// 	// erasing pagination caches
// 	_paginated_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___paginated*").Result()
// 	if err == nil {
// 		for _, key := range _paginated_keys {
// 			database.REDIS_DB_CONNECTION.Del(ctx, key)
// 			log.WithFields(log.Fields{
// 				"key": key,
// 			}).Debugln(">>>>>>>>>>>>>>>> Redis, users___paginated removed")
// 		}
// 	}
// }

func eraseAllUserPaginationCache(ctx context.Context) {
	database.REDIS_DB_CONNECTION.Del(ctx, "users___paginated")
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
	const max_users_update_in_progress_ttl = time.Minute * 15

	var redis_lock sync.Mutex
	redis_lock.Lock()
	users_update_in_progress, err := database.RedisPoolGet("users_update_in_progress")
	if err == nil {
		// if key users_update_in_progress present with some count
		users_update_in_progress_int, _ := strconv.ParseInt(users_update_in_progress, 10, 64)
		if operation == "delete" {
			users_update_in_progress_int--
			if users_update_in_progress_int == 0 {
				database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
				log.Debugln("deleted users_update_in_progress")
			}
		} else {
			users_update_in_progress_int++
		}

		// log.WithFields(log.Fields{
		// 	"users_update_in_progress_int":users_update_in_progress_int,
		// }).Debugln("modifyCacheProgressStatus")

		if users_update_in_progress_int != 0 {
			database.RedisPoolSet("users_update_in_progress", strconv.FormatInt(users_update_in_progress_int, 10), max_users_update_in_progress_ttl)
		}
	} else {
		// if key users_update_in_progress not exists
		if operation != "delete" {
			database.RedisPoolSet("users_update_in_progress", "1", max_users_update_in_progress_ttl)
		}
	}
	redis_lock.Unlock()
}

var invalidate_cache_timeout context.CancelFunc = nil

func invalidateCache(_id string) {
	// log.WithFields(log.Fields{
	// 	"_id": _id,
	// }).Debugln("invalidateCache....")
	const max_users_update_in_progress_ttl = time.Minute * 15

	if invalidate_cache_timeout != nil {
		invalidate_cache_timeout()
	}
	database.REDIS_DB_CONNECTION.Set(context.Background(), "users_update_in_progress", "1", max_users_update_in_progress_ttl)
	go deleteUserCache(_id, context.Background())
	cb := func() {
		ctx := context.Background()
		eraseAllUserPaginationCache(ctx)
		getUsersCount(ctx)
		database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
	}
	invalidate_cache_timeout = my_modules.SetTimeOut(cb, time.Millisecond*500)

	// modifyCacheProgressStatus("insert", ctx)
	// modifyCacheProgressStatus("delete", ctx)

	// getUsersCount(ctx)
}

func OnUserModification(_id string, userData database.UsersModel, operationType string) {
	invalidateCache(_id)
}
