package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adnux/go-rest-api/utils"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request)
type contextKey string
type EnsureAuth struct {
	handler AuthenticatedHandler
}

const authUserIDKey contextKey = "authUserId"

func (ensureAuth *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Ensuring authentication...")
	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Not authorized."}`))
		return
	}

	authUserId, err := utils.VerifyToken(token)

	if err != nil {
		http.Error(w, "Not authorized.", http.StatusUnauthorized)
	}

	r.Header.Set("authUserId", fmt.Sprintf("%d", authUserId))
	ctx := context.WithValue(r.Context(), authUserIDKey, authUserId)
	ensureAuth.handler(w, r.WithContext(ctx))
}

func EnsureAuthHandler(handlerToWrap AuthenticatedHandler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}
