package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) metrics(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "text/html")
	writer.WriteHeader(http.StatusOK)

	content := fmt.Sprintf(`
	<html>
	<body>
		<h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	</html>
	`, cfg.fileserverHits.Load())

	writer.Write([]byte(content))
}
