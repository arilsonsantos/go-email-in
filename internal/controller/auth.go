package controller

import (
	"context"
	"emailn/internal/controller/utils"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			authorizationFailed("request does not contain an authorization token", w)
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		err, token := decodeToken(tokenStr)
		if !token.Valid || err != nil {
			authorizationFailed("invalid token", w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "couldn't parse token claims"})
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

func decodeToken(tokenStr string) (error, *jwt.Token) {
	jwksUrl, exists := os.LookupEnv("KEYCLOAK_URL_CERTS")
	if !exists {
		jwksUrl = "http://localhost:8080/realms/provider/protocol/openid-connect/certs"
	}

	jwks, err := utils.FetchKeycloakJWKS(jwksUrl)

	var jwtKeys utils.JWKS
	for _, key := range jwks.Keys {
		if key.Alg == "RS256" {
			jwtKeys.Keys = append(jwtKeys.Keys, key)
		}
	}

	keycloakPublicKey := jwtKeys.Keys[0].X5c[0]
	decoded, err := base64.StdEncoding.DecodeString(keycloakPublicKey)
	if err != nil {
		panic(err)
	}

	pemKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: decoded,
	})

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwt.ParseRSAPublicKeyFromPEM(pemKey)
	})
	return err, token
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
