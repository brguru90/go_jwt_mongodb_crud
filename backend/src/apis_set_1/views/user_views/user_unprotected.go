package user_views

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

func GetAllUserData(c *gin.Context) {
	ctx := c.Request.Context()

	var err error
	var cursor *mongo.Cursor
	cursor, err = database.MONGO_COLLECTIONS.Users.Find(ctx, bson.M{})
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
				log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				// continue
				my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
				return
			}
			usersData = append(usersData, userData)
		}

		my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", map[string]interface{}{
			"users":       usersData,
		})
	}
}
