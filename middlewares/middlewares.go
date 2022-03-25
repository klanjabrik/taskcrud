package middlewares

import (
	"net/http"

	api "taskcrud/api"
	"taskcrud/models"
	"taskcrud/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func TaskOwnershipValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		loggedinUserId, _ := token.ExtractTokenID(c)
		m := map[string]interface{}{"user_id": loggedinUserId, "id": c.Param("id")}

		if err := models.GetOneTask(&task, m); err != nil {
			api.RespondError(c, http.StatusBadRequest, api.WithMessageError(err.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}
