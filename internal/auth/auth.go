package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetHeaderToken(headers http.Header, key string, prefix string) (string, error) {
	authHeader := headers.Get(key)

	if authHeader == "" {
		return "", fmt.Errorf("missing %s", strings.ToLower(key))
	}

	authHeader, found := strings.CutPrefix(authHeader, fmt.Sprintf("%s ", prefix))

	if !found {
		return "", fmt.Errorf("invalid %s header", strings.ToLower(key))
	}

	return authHeader, nil
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPasswordHash(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			Subject:   userID.String(),
		},
	)
	signed, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", err
	}

	return signed, err
}

func ValidateJWT(tokenString string, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(tokenSecret), nil
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	id, err := token.Claims.GetSubject()

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(id), err
}

func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(token), nil
}
