package controller

import (
	"context"
	"emailn/internal/infrastructure/credential"
	"github.com/go-chi/render"
	"net/http"
)

type ValidationTokenFunc func(tokenStr string) (string, error)

var ValidationToken ValidationTokenFunc = credential.ValidateToken

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization token"})
			return
		}

		email, err := ValidationToken(tokenStr)

		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
