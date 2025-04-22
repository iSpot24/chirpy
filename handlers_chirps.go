package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/iSpot24/chirpy/internal/auth"
	"github.com/iSpot24/chirpy/internal/database"
)

func (cfg *apiConfig) getChirpById(writer http.ResponseWriter, request *http.Request) {
	chirpID, err := uuid.Parse(request.PathValue("chirpID"))

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	chirp, err := cfg.db.GetChirpById(context.Background(), chirpID)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("content-type", "application/json")
	err = respondWithJSON(writer, http.StatusOK, chirp)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func (cfg *apiConfig) deleteChirpById(writer http.ResponseWriter, request *http.Request) {
	token, err := auth.GetHeaderToken(request.Header, "Authorization", "Bearer")

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	chirpID, err := uuid.Parse(request.PathValue("chirpID"))

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	chirp, err := cfg.db.GetChirpById(context.Background(), chirpID)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if chirp.UserID != userID {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	err = cfg.db.DeleteChirpById(context.Background(), chirp.ID)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (cfg *apiConfig) getChirps(writer http.ResponseWriter, request *http.Request) {
	var err error
	var userID uuid.NullUUID

	if authorID := request.URL.Query().Get("author_id"); authorID != "" {
		authorUUID, err := uuid.Parse(authorID)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		userID = uuid.NullUUID{
			UUID:  authorUUID,
			Valid: true,
		}
	}

	chirps, err := cfg.db.GetChirps(context.Background(), userID)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if orderDirectionQuery := request.URL.Query().Get("sort"); orderDirectionQuery == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		})
	}

	writer.Header().Set("content-type", "application/json")
	err = respondWithJSON(writer, http.StatusOK, chirps)

	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *apiConfig) createChirp(writer http.ResponseWriter, request *http.Request) {
	token, err := auth.GetHeaderToken(request.Header, "Authorization", "Bearer")

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var params database.Chirp
	err = json.NewDecoder(request.Body).Decode(&params)

	writer.Header().Set("content-type", "application/json")

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Something went wrong")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if len(params.Body) > 140 {
		err = respondWithError(writer, http.StatusBadRequest, "Chirp is too long")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	words := strings.Split(params.Body, " ")
	for i, word := range words {
		lcWord := strings.ToLower(word)
		if lcWord == "kerfuffle" || lcWord == "sharbert" || lcWord == "fornax" {
			words[i] = "****"
		}
	}

	chirp, err := cfg.db.CreateChirp(context.Background(), database.CreateChirpParams{
		ID:        uuid.New(),
		Body:      strings.Join(words, " "),
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		err = respondWithError(writer, http.StatusBadRequest, "Chirp cannot be created")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = respondWithJSON(writer, http.StatusCreated, chirp)

	if err != nil {
		log.Fatal(err)
	}
}
