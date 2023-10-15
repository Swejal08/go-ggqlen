package middleware

import (
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"userId"}

type contextKey struct {
	userId string
}

func UserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userId := r.Header.Get("UserId")

			if userId == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				errMsg := `{"error": "Header not found"}`
				w.Write([]byte(errMsg))
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
