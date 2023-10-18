package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Swejal08/go-ggqlen/utils"
)

// var userCtxKey = &contextKey{"userId"}

var CurrentUserId string

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			parts := strings.Split(auth, "Bearer")
			token := strings.Trim(parts[1], " ")

			secret := os.Getenv("JWT_SECRET")
			user, err := utils.VerifyToken(token, secret)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			CurrentUserId = "currentUserId"

			ctx := context.WithValue(r.Context(), CurrentUserId, user.ID)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
