package spotify

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateCodeVerifier() (string, error) {
	verifier := make([]byte, 32) // base64 43 chars

	_, err := rand.Read(verifier)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(verifier), nil
}

func GenerateCodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func GeneratePKCE() (string, string) {
	verifier, err := GenerateCodeVerifier()
	if err != nil {
		panic(err)
	}

	challenge := GenerateCodeChallenge(verifier)

	return verifier, challenge
}
