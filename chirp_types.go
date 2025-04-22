package main

import (
	"sync/atomic"

	"github.com/iSpot24/chirpy/internal/database"
)

const baseURL = "/app"

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	jwtSecret      string
	polkaKey       string
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PolkaWebhook struct {
	Event string `json:"event"`
	Data  struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

type errorResponse struct {
	Error string `json:"error"`
}
