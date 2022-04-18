package user_views

import (
	"context"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

// @BasePath /api
// @Summary url to signup
// @Schemes
// @Description allow people to create new to user account
// @Tags SignUp
// @Accept json
// @Produce json
// @Param new_user body database.UsersModel true "Add user"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Router /sign_up [post]
func SignUp(c *gin.Context) {
	ctx := context.Background()
	var newUserRow database.UsersModel
	// ShouldBindJSON will validate json body & convert it to structure object
	if err := c.ShouldBindJSON(&newUserRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	var newUserData database.UsersModel
	{
		_time := time.Now()
		newUserData = database.UsersModel{
			Email:       newUserRow.Email,
			Name:        newUserRow.Name,
			Description: newUserRow.Description,
			CreatedAt:   _time,
			UpdatedAt:   _time,
		}
		ins_res, ins_err := database.MONGO_COLLECTIONS.Users.InsertOne(ctx, newUserData)
		if ins_err != nil {
			log.WithFields(log.Fields{
				"ins_err": ins_err,
			}).Errorln("Error in inserting data to mongo users")
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user", nil)
			return
		} else {
			newUserData.ID = ins_res.InsertedID.(primitive.ObjectID)
			// log.WithFields(log.Fields{
			// 	"ins_res": ins_res.InsertedID,
			// 	"type":    fmt.Sprintf("%T", ins_res.InsertedID),
			// }).Info("successfully inserted data into mongo users")
		}
	}
	access_token_payload := my_modules.Authenticate(c, newUserData)
	{
		_time := time.Now()
		_, ins_err := database.MONGO_COLLECTIONS.ActiveSessions.InsertOne(ctx, database.ActiveSessionsModel{
			UserID:    newUserData.ID,
			TokenID:   access_token_payload.Token_id,
			IP:        c.ClientIP(),
			UA:        c.GetHeader("User-Agent"),
			Exp:       access_token_payload.Exp,
			Status:    "active",
			CreatedAt: _time,
			UpdatedAt: _time,
		})
		if ins_err != nil {
			log.WithFields(log.Fields{
				"ins_err": ins_err,
			}).Errorln("Error in inserting data to mongo users")
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user while marking active", nil)
			return
		}
	}
	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Regesteration successfully", newUserData)
}

type UserEmailID struct {
	Email string `json:"email" binding:"required"`
}

// @BasePath /api
// @Summary url to login
// @Schemes
// @Description allow people to login into their account
// @Tags Login
// @Accept json
// @Produce json
// @Param new_user body UserEmailID true "Add user"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /login [post]
func Login(c *gin.Context) {
	ctx := context.Background()
	var userEmailID UserEmailID
	if err := c.ShouldBindJSON(&userEmailID); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	var userData database.UsersModel
	{
		err := database.MONGO_COLLECTIONS.Users.FindOne(ctx, bson.M{
			"email": userEmailID.Email,
		}).Decode(&userData)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.WithFields(log.Fields{
					"Error": err,
					"Email": userData.Email,
				}).Warning("Error in finding user email")
				my_modules.CreateAndSendResponse(c, http.StatusForbidden, "error", "Invalid credential", nil)
				return
			}
			log.WithFields(log.Fields{
				"Error": err,
				"Email": userData.Email,
			}).Error("Error in finding user email")
			my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Error in finding user", nil)
			return
		}
	}
	access_token_payload := my_modules.Authenticate(c, userData)
	{
		_time := time.Now()
		_, ins_err := database.MONGO_COLLECTIONS.ActiveSessions.InsertOne(ctx, database.ActiveSessionsModel{
			UserID:    userData.ID,
			TokenID:   access_token_payload.Token_id,
			IP:        c.ClientIP(),
			UA:        c.GetHeader("User-Agent"),
			Exp:       access_token_payload.Exp,
			Status:    "active",
			CreatedAt: _time,
			UpdatedAt: _time,
		})
		if ins_err != nil {
			log.WithFields(log.Fields{
				"ins_err": ins_err,
			}).Errorln("Error in inserting data to mongo users")
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user while marking active", nil)
			return
		}
	}
	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Authorization success", userData)
}

// @BasePath /api
// @Summary
// @Schemes
// @Description api used to validate user login session
// @Tags Login status
// @Accept json
// @Produce json
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /login_status [get]
func LoginStatus(c *gin.Context) {
	decoded_token, err, http_status, ok := my_modules.LoginStatus(c,false)
	if err != "" {
		my_modules.CreateAndSendResponse(c, http_status, "error", err, nil)
		return
	}
	if ok {
		my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "active", decoded_token.Data)
	}
}
