package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/iSpot24/chirpy/internal/auth"
)

func (cfg *apiConfig) markUserAsChirpyRed(writer http.ResponseWriter, request *http.Request) {
	polkaKey, err := auth.GetHeaderToken(request.Header, "Authorization", "ApiKey")

	if err != nil || polkaKey != cfg.polkaKey {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var webhook PolkaWebhook
	err = json.NewDecoder(request.Body).Decode(&webhook)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if webhook.Event != "user.upgraded" {
		writer.WriteHeader(http.StatusNoContent)
		return
	}

	userID, err := uuid.Parse(webhook.Data.UserID)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = cfg.db.MarkUserAsChirpyRed(context.Background(), userID)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
