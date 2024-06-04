package credential

import (
	"emailn/internal/controller/utils"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func ValidateToken(tokenStr string, r *http.Request, w http.ResponseWriter) (string, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	err, token := decodeToken(tokenStr)
	if !token.Valid || err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("couldn't parse token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email claim not found or invalid")
	}
	return email, nil
}

func decodeToken(tokenStr string) (error, *jwt.Token) {
	jwksUrl := os.Getenv("KEYCLOAK_URL_CERTS")
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

func AuthorizationFailed(message string, w http.ResponseWriter) Res401Struct {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	data := Res401Struct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
	return data
}
