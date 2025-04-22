package main

import "net/http"

func readiness(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}
