package routes

import (
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/models"
	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}

	user, err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user.",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse request data.",
		})
		return
	}
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Could not authenticate user.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not authenticate user.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not parse user id.",
		})
		return
	}

	userToDelete := models.User{ID: userId}
	err = userToDelete.DeleteUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not delete user.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted!",
	})
}
