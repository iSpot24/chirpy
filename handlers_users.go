package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iSpot24/chirpy/internal/auth"
	"github.com/iSpot24/chirpy/internal/database"
)

func (cfg *apiConfig) createUser(writer http.ResponseWriter, request *http.Request) {
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

	hashedPassword, err := auth.HashPassword(params.Password)

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	user, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:             uuid.New(),
		Email:          params.Email,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		HashedPassword: hashedPassword,
	})

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = respondWithJSON(writer, http.StatusCreated, user)

	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *apiConfig) updateUser(writer http.ResponseWriter, request *http.Request) {
	token, err := auth.GetHeaderToken(request.Header, "Authorization", "Bearer")

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)

	if err != nil {
		err = respondWithError(writer, http.StatusUnauthorized, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	var params User
	err = json.NewDecoder(request.Body).Decode(&params)

	writer.Header().Set("content-type", "application/json")

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	user, err := cfg.db.UpdateUser(context.Background(), database.UpdateUserParams{
		ID:             userID,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = respondWithJSON(writer, http.StatusOK, user)

	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *apiConfig) deleteUsers(writer http.ResponseWriter, request *http.Request) {
	if cfg.platform != "dev" {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	err := cfg.db.DeleteUsers(context.Background())

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
}
