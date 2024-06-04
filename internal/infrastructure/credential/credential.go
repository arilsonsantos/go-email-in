package credential

import (
	"emailn/internal/controller/utils"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

const invalid_token = "invalid token"

func ValidateToken(tokenStr string) (string, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := decodeToken(tokenStr)

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		println("couldn't parse token claims")
		return "", errors.New(invalid_token)
	}

	email, ok := claims["email"].(string)
	if !ok {
		println("email claim not found or invalid")
		return "", errors.New(invalid_token)
	}

	return email, nil
}

func decodeToken(tokenStr string) (*jwt.Token, error) {
	jwksUrl := os.Getenv("KEYCLOAK_URL_CERTS")
	jwks, err := utils.FetchKeycloakJWKS(jwksUrl)

	var jwtKeys utils.JWKS
	for _, key := range jwks.Keys {
		if key.Alg == "RS256" {
			jwtKeys.Keys = append(jwtKeys.Keys, key)
		}
	}

	if len(jwtKeys.Keys) == 0 {
		fmt.Println("jwtKeys está vazio")
		return nil, errors.New(invalid_token)
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
			println("unexpected signing method")
			return fmt.Errorf(invalid_token), nil
		}
		return jwt.ParseRSAPublicKeyFromPEM(pemKey)
	})
	return token, err
}
