package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/iSpot24/chirpy/internal/auth"
	"github.com/iSpot24/chirpy/internal/database"
)

func (cfg *apiConfig) makeJWTWithRefresh(userID uuid.UUID, withRefreshToken bool) (token string, refreshToken string, err error) {
	token, err = auth.MakeJWT(userID, cfg.jwtSecret, time.Duration(time.Second*time.Duration(3600)))

	if err != nil {
		return "", "", err
	}

	refreshToken, err = auth.MakeRefreshToken()

	if err != nil {
		return "", "", err
	}

	if !withRefreshToken {
		return token, "", nil
	}

	err = cfg.db.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(24) * 60),
	})

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
