package api_modules

import (
	"learn_go/src/my_modules"
	"time"

	"github.com/gin-gonic/gin"
)

func ForUserPagination(c *gin.Context) string {
	if c.Query("page") != "" {
		// admin uuid could have been return,if access level is implemented
		return "paginated"
	}
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return "id=" + time.Now().String()
	}
	return "id=" + payload.Data.ID
}
