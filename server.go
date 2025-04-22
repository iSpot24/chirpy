package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/iSpot24/chirpy/internal/database"
	"github.com/joho/godotenv"
)

func getConfig() *apiConfig {
	godotenv.Load()

	return &apiConfig{
		platform:  os.Getenv("PLATFORM"),
		jwtSecret: os.Getenv("JWT_SECRET"),
		polkaKey:  os.Getenv("POLKA_KEY"),
	}
}

func initDB() *database.Queries {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal(err)
	}
	return database.New(db)
}

func startServer(cfg *apiConfig) {
	mux := http.NewServeMux()

	mux.Handle(baseURL, cfg.middlewareMetricsInc(http.StripPrefix(baseURL, http.FileServer(http.Dir(".")))))
	mux.Handle(baseURL+"/assets/", cfg.middlewareMetricsInc(http.StripPrefix(baseURL+"/assets/", http.FileServer(http.Dir("assets")))))
	mux.HandleFunc("GET /api/healthz", readiness)
	mux.HandleFunc("GET /admin/metrics", cfg.metrics)

	mux.HandleFunc("POST /api/login", cfg.login)
	mux.HandleFunc("POST /api/refresh", cfg.refresh)
	mux.HandleFunc("POST /api/revoke", cfg.revoke)

	mux.HandleFunc("POST /admin/reset", cfg.deleteUsers)
	mux.HandleFunc("POST /api/users", cfg.createUser)
	mux.HandleFunc("PUT /api/users", cfg.updateUser)

	mux.HandleFunc("POST /api/chirps", cfg.createChirp)
	mux.HandleFunc("GET /api/chirps", cfg.getChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.getChirpById)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", cfg.deleteChirpById)

	mux.HandleFunc("POST /api/polka/webhooks", cfg.markUserAsChirpyRed)

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + os.Getenv("DB_PORT"),
	}

	log.Fatal(server.ListenAndServe())
}
