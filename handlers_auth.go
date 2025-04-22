package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/iSpot24/chirpy/internal/auth"
	"github.com/iSpot24/chirpy/internal/database"
)

func (cfg *apiConfig) login(writer http.ResponseWriter, request *http.Request) {
	var params User
	err := json.NewDecoder(request.Body).Decode(&params)

	writer.Header().Set("content-type", "application/json")

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	user, err := cfg.db.GetUserByEmail(context.Background(), params.Email)

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, "Incorrect email or password")

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = auth.CheckPasswordHash([]byte(user.HashedPassword), params.Password)

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, "Incorrect email or password")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	token, refreshToken, err := cfg.makeJWTWithRefresh(user.ID, true)

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, "Incorrect email or password")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = respondWithJSON(writer, http.StatusOK, map[string]any{
		"id":            user.ID,
		"email":         user.Email,
		"is_chirpy_red": user.IsChirpyRed,
		"created_at":    user.CreatedAt,
		"updated_at":    user.UpdatedAt,
		"token":         token,
		"refresh_token": refreshToken,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *apiConfig) refresh(writer http.ResponseWriter, request *http.Request) {
	refreshTokenString, err := auth.GetHeaderToken(request.Header, "Authorization", "Bearer")

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	refreshToken, err := cfg.db.GetRefreshToken(context.Background(), refreshTokenString)

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if refreshToken.RevokedAt.Valid && time.Now().UTC().Compare(refreshToken.RevokedAt.Time) >= 0 ||
		time.Now().UTC().Compare(refreshToken.ExpiresAt) >= 0 {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, _, err := cfg.makeJWTWithRefresh(refreshToken.UserID, false)

	if err != nil {
		err = respondWithError(writer, http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = respondWithJSON(writer, http.StatusOK, map[string]any{
		"token": token,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *apiConfig) revoke(writer http.ResponseWriter, request *http.Request) {
	refreshTokenString, err := auth.GetHeaderToken(request.Header, "Authorization", "Bearer")

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = cfg.db.UpdateRevokedAt(context.Background(), database.UpdateRevokedAtParams{
		Token:     refreshTokenString,
		RevokedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		err = respondWithError(writer, http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
