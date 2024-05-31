package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JWKS struct {
	Keys []struct {
		Alg string   `json:"alg"`
		E   string   `json:"e"`
		Kid string   `json:"kid"`
		Kty string   `json:"kty"`
		N   string   `json:"n"`
		X5c []string `json:"x5c"`
	} `json:"keys"`
}

func FetchKeycloakJWKS(url string) (*JWKS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchKeycloakJWKS: failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetchKeycloakJWKS: failed to read the response body: %v", err)
	}

	var jwks JWKS
	if err = json.Unmarshal(body, &jwks); err != nil {
		return nil, fmt.Errorf("fetchKeycloakJWKS: failed to unmarshal the response body: %v", err)
	}

	return &jwks, nil
}
