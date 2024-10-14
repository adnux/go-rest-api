package middlewares

import (
	"net/http"

	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized.",
		})
		return
	}

	authUserId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized.",
			"error":   err.Error(),
		})
		return
	}

	context.Set("authUserId", authUserId)
	context.Next()
}
