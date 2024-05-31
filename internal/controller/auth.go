package controller

import (
	"context"
	"encoding/json"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			authorizationFailed("request does not contain an authorization token", w)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "reflect_it"})
		_, err = verifier.Verify(r.Context(), token)
		if err != nil {
			authorizationFailed("invalid token", w)
			return
		}

		idToken, err := verifier.Verify(r.Context(), token)

		claims := map[string]interface{}{}
		if err := idToken.Claims(&claims); err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error parsing claims"})
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "email claim not found or invalid"})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type Res401Struct struct {
	Status   string
	HTTPCode int
	Message  string
}

func authorizationFailed(message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	data := Res401Struct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
	res, _ := json.Marshal(data)
	w.Write(res)
}
