package controller

import (
	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization token"})
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "reflect_it"})
		_, err = verifier.Verify(r.Context(), token)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
