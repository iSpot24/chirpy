package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(writer http.ResponseWriter, code int, payload interface{}) error {
	resp, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	writer.WriteHeader(code)
	_, err = writer.Write(resp)

	return err
}

func respondWithError(writer http.ResponseWriter, code int, msg string) error {
	writer.WriteHeader(code)

	errResp := errorResponse{Error: msg}
	errorData, err := json.Marshal(errResp)

	if err != nil {
		return err
	}

	_, err = writer.Write(errorData)

	return err
}
