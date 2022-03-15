package user_views

import (
	"context"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const api_secret = "1234"


func deleteUsercache(_id string, ctx context.Context) {
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

// @BasePath /api
// @Summary InvalidateUsercache
// @Schemes
// @Description will be used in postgres trigger to delete redis cache
// @Tags Delete user cache
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param secret header string true "trigger secret"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /del_user_cache/{id} [get]
func InvalidateUsercache(c *gin.Context) {
	if c.GetHeader("secret") != api_secret {
		my_modules.CreateAndSendResponse(c, http.StatusForbidden, "error", "Invalid secret", nil)
		return
	}
	ctx := context.Background()
	database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", "true", time.Second*0)

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

	// db_connection := database.POSTGRES_DB_CONNECTION

	cursor, err:=database.MONGO_COLLECTIONS.Users.Find(ctx,bson.M{
		"_id": c.Param("id"),
	})
	if err != nil {
		if err != context.Canceled {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorln("QueryRow failed ==>")
		}
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
		return
	} else { 
		defer cursor.Close(ctx)
		for cursor.Next(c.Request.Context()) {
			var userData database.UsersModel
			if err = cursor.Decode(&userData); err != nil {
				// log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				// continue
				my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
				return
			}
			deleteUsercache(userData.ID.Hex(), ctx)
		}
	}



	database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
	log.Infoln("deleted users_update_in_progress")

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
