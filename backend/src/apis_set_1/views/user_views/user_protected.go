package user_views

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"learn_go/src/database"
	"learn_go/src/my_modules"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @BasePath /api
// @Summary url to view user data
// @Schemes
// @Description allow people to view their user profile data
// @Tags View user data
// @Accept json
// @Produce json
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/ [get]
func GetUserData(c *gin.Context) {
	ctx := context.Background()
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}
	var id string = payload.Data.ID
	_id, err_id := primitive.ObjectIDFromHex(payload.Data.ID)

	var _limit int64 = 100
	var _page int64 = 0

	if c.Query("page") != "" {
		_page, _ = strconv.ParseInt(c.Query("page"), 10, 64)
		if c.Query("limit") != "" {
			_limit, _ = strconv.ParseInt(c.Query("limit"), 10, 64)
		}
	}

	if id != "" && err_id == nil {
		var err error
		var cursor *mongo.Cursor
		if _page > 0 {
			var _offset = _limit * (_page - 1)
			log.WithFields(log.Fields{
				"Skip":  _offset,
				"Limit": _limit,
			}).Info("Pagination...")
			cursor, err = database.MONGO_COLLECTIONS.Users.Find(c.Request.Context(), bson.M{},
				&options.FindOptions{
					Sort: bson.M{
						"_id": 1,
					},
					Skip:  &_offset,
					Limit: &_limit,
				})
		} else {
			cursor, err = database.MONGO_COLLECTIONS.Users.Find(ctx, bson.M{
				"_id": _id,
			})
		}

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
			var usersData []database.UsersModel = []database.UsersModel{}
			for cursor.Next(c.Request.Context()) {
				var userData database.UsersModel
				if err = cursor.Decode(&userData); err != nil {
					// log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
					// continue
					my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
					return
				}
				usersData = append(usersData, userData)
			}

			if _page > 0 {
				total_users, _ := database.REDIS_DB_CONNECTION.Get(context.Background(), "users_count").Result()
				total_users_int, _ := strconv.ParseInt(total_users, 10, 64)
				my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", map[string]interface{}{
					"users":       usersData,
					"cur_page":    _page,
					"total_users": total_users_int,
				})
				return
			} else {
				my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", usersData[0])
				return
			}
		}
	} else {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Didn't got UUID", nil)
		return
	}
}


// @BasePath /api
// @Summary url to update user data
// @Schemes
// @Description allow people to update their user profile data
// @Tags Update user data
// @Accept json
// @Produce json
// @Param new_user body database.UsersModel true "Add user"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/ [put]
func UpdateUserData(c *gin.Context) {
	ctx := context.Background()
	// db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var updateWithData database.UsersModel

	if err := c.ShouldBindJSON(&updateWithData); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	_time := time.Now()
	// overide any uuid with user uuid
	_id, _ := primitive.ObjectIDFromHex(payload.Data.ID)

	res, err := database.MONGO_COLLECTIONS.Users.UpdateOne(
		ctx,
		bson.M{
			"_id": _id,
		},
		bson.M{
			"$set": bson.M{
				"email":       updateWithData.Email,
				"name":        updateWithData.Name,
				"description": updateWithData.Description,
				"updatedAt":   _time}},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorln("Failed to update user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to update data", nil)
		return
	}
	var response_data = make(map[string]interface{})
	response_data["updated_with_data"] = updateWithData
	response_data["updated_count"] = res.ModifiedCount
	response_data["match_count"] = res.MatchedCount
	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", response_data)
}

// @BasePath /api
// @Summary get active user login session
// @Schemes
// @Description return the active user session/token across all browser
// @Tags Get Active sessions
// @Accept json
// @Produce json
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/active_sessions/ [get]
func GetActiveSession(c *gin.Context) {
	ctx := context.Background()
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}
	_id, _ := primitive.ObjectIDFromHex(payload.Data.ID)

	cursor, err := database.MONGO_COLLECTIONS.ActiveSessions.Find(ctx, bson.M{
		"user_id":  _id,
		"token_id": bson.M{"$ne": payload.Token_id},
	})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.WithFields(log.Fields{
				"err": err,
			}).Errorln("Failed to load session data")
		}
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
		return
	} else {
		defer cursor.Close(ctx)
		var sessionsData []database.ActiveSessionsModel = []database.ActiveSessionsModel{}
		for cursor.Next(c.Request.Context()) {
			var sessionData database.ActiveSessionsModel
			if err = cursor.Decode(&sessionData); err != nil {
				my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving session data", nil)
				return
			}
			sessionsData = append(sessionsData, sessionData)
		}
		my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", sessionsData)
		return
	}

}

// @BasePath /api
// @Summary url to delete user account
// @Schemes
// @Description allow people to delete their account
// @Tags Delete user account
// @Accept json
// @Produce json
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/ [delete]
func Deleteuser(c *gin.Context) {
	// db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var id string = payload.Data.ID
	_id, _id_err := primitive.ObjectIDFromHex(payload.Data.ID)

	if id == "" || _id_err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "UUID of user is not provided", _id_err)
		return
	}

	result, err := database.MONGO_COLLECTIONS.Users.DeleteOne(context.Background(), bson.M{
		"_id": _id,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorln("Failed to delete user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to delete user data", nil)
		return
	}
	rows_deleted := result.DeletedCount

	// const sql_stmt string = `DELETE FROM users WHERE uuid=$1`
	// res, err := db_connection.Exec(context.Background(), sql_stmt, uuid)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"err": err,
	// 		"sql": fmt.Sprintf(`DELETE FROM users WHERE uuid=%s`, uuid),
	// 	}).Errorln("Failed to delete user data")
	// 	my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to delete user data", nil)
	// 	return
	// }

	// rows_deleted := res.RowsAffected()

	var response_data = make(map[string]interface{})
	response_data["deleted_user_with_uuid"] = id
	response_data["deleted_count"] = rows_deleted
	if rows_deleted > 0 {
		my_modules.DeleteCookie(c, "access_token")
		my_modules.DeleteCookie(c, "user_data")
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", response_data)

}

// @BasePath /api
// @Summary allow user to logout
// @Schemes
// @Description API allow user to logout, which delete the cookie which stores token
// @Tags Logout
// @Accept json
// @Produce json
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/logout/ [get]
func Logout(c *gin.Context) {
	my_modules.DeleteCookie(c, "access_token")
	my_modules.DeleteCookie(c, "user_data")
	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Logged out", nil)
}

func updateActiveSession(activeSessionsRow database.ActiveSessionsModel, status string) (int64, error) {
	// db_connection := database.POSTGRES_DB_CONNECTION

	where_map:=map[string]interface{}{
		"user_id": activeSessionsRow.UserID,
		"token_id":activeSessionsRow.TokenID,
		"exp":activeSessionsRow.Exp,
	}
	if status == "blocked" {
		where_map["status"]="active"
	}

	where,err_bson:=bson.Marshal(where_map)
	if err_bson!=nil{
		return -1, err_bson
	}


	result,err:=database.MONGO_COLLECTIONS.ActiveSessions.UpdateOne(
		context.Background(),
		where,
		bson.M{
			"$set":bson.M{ "status":status},
		},
	)
	if err!=nil{
		log.WithFields(log.Fields{
			"err": err,			
		}).Errorln("Failed to update user data")
		return -1, err
	}
	rows_updated:=result.ModifiedCount
	return rows_updated, nil
}

// @BasePath /api
// @Summary block specified session
// @Schemes
// @Description Adds the token of user to block list based on provided token id
// @Tags Block sessions
// @Accept json
// @Produce json
// @Param block_active_session body database.ActiveSessionsModel true "block token"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/block_token/ [post]
func BlockSession(c *gin.Context) {
	redis_db_connection := database.REDIS_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var activeSessionsRow database.ActiveSessionsModel
	if err := c.ShouldBindJSON(&activeSessionsRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	// overide any existing uuid with uuid of current user
	_id, _ := primitive.ObjectIDFromHex(payload.Data.ID)
	activeSessionsRow.UserID = _id

	rows_updated, err := updateActiveSession(activeSessionsRow, "blocked")
	if err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to update data", nil)
		return
	}
	if rows_updated <= 0 {
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Token doesn't exists/Already blacklisted", nil)
		return
	}
	var exp_sec time.Duration = time.UnixMilli(activeSessionsRow.Exp).UTC().Sub(time.Now().UTC())
	r_err := redis_db_connection.SetEX(context.Background(), activeSessionsRow.TokenID, activeSessionsRow.UserID.Hex(),
		exp_sec).Err()
	if r_err != nil {
		rows_updated, err := updateActiveSession(activeSessionsRow, "active")
		log.WithFields(log.Fields{
			"redis_err":        r_err,
			"sql_err":          err,
			"sql_rows_updated": rows_updated,
		}).Errorln("Failed to insert data on redis")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to blacklist the session", nil)
		return
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Blocked", rows_updated)
}
