package main

import (
	_ "github.com/lib/pq"
)

func main() {
	cfg := getConfig()
	cfg.db = initDB()

	startServer(cfg)
}
