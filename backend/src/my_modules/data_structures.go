package my_modules

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// structure member first letter should be capital indorser to export it as json
// binding encforce required condition
// empty interface data type allow unknown json format for which we don't have structure defined
// empty interface converted to map object

type ResponseFormat struct {
	Status string      `json:"status" binding:"required"`
	Msg    string      `json:"msg" binding:"required"`
	Data   interface{} `json:"data" binding:"required"`
}



// type ActiveSessionsRow struct {
// 	Column_id        primitive.ObjectID       `json:"_id"`
// 	Column_user_id primitive.ObjectID      `json:"user_id"`
// 	Column_token_id  string      `json:"token_id" binding:"required"`
// 	Column_ua        string      `json:"ua"`
// 	Column_ip        string      `json:"ip"`
// 	Column_exp       int64       `json:"exp" binding:"required"`
// 	Column_status    string      `json:"status"`
// 	Column_createdAt time.Time  `json:"createdAt"`
// 	Column_updatedAt time.Time  `json:"updatedAt"`
// }

type ResponseCacheStruct struct {
	ResponseData   string    `json:"response_data" binding:"required"`
	ContentType    string    `json:"content_type" binding:"required"`
	HTTPStatusCode int       `json:"http_status_code" binding:"required"`
	LastModified   time.Time `json:"last_modified" binding:"required"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
