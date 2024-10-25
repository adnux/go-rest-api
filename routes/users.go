package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adnux/go-rest-api/models"
	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Could not parse request data.", http.StatusBadRequest)
		return
	}

	err = user.Save()

	if err != nil {
		http.Error(w, "Could not save user.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonData, err := json.Marshal(gin.H{
		"message": "User created successfully",
		"user":    user,
	})
	if err != nil {
		http.Error(w, "Could not marshal response data.", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Could not parse request data.", http.StatusBadRequest)
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		http.Error(w, "Could not authenticate user.", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		http.Error(w, "Could not authenticate user.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(gin.H{
		"message": "Login successful!",
		"token":   token,
	})
	if err != nil {
		http.Error(w, "Could not marshal response data.", http.StatusInternalServerError)
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userToDelete := models.User{ID: userId}
	err = userToDelete.DeleteUser()

	if err != nil {
		http.Error(w, "Could not delete user.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(gin.H{
		"message": "User deleted!",
	})

	if err != nil {
		http.Error(w, "Could not marshal response data.", http.StatusInternalServerError)
		return
	}

}
