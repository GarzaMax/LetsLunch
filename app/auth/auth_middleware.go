package auth

import (
	users_repository "cmd/app/entities/user/repository"
	"context"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strings"
)

type AuthConfig struct {
	users     users_repository.UsersRepository
	tokenAuth *jwtauth.JWTAuth
}

func NewAuthConfig(
	users users_repository.UsersRepository,
	tokenAuth *jwtauth.JWTAuth,
) *AuthConfig {
	return &AuthConfig{
		users:     users,
		tokenAuth: tokenAuth,
	}
}

func (a *AuthConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "user", "123")
		fmt.Println("received with token: ", r.Header.Get("Authorization"))

		tockenStr := r.Header.Get("Authorization")
		tockenStr, found := strings.CutPrefix(tockenStr, "Bearer ")
		if found == false {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err := a.tokenAuth.Decode(tockenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		//fmt.Println(token.AsMap(ctx))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
