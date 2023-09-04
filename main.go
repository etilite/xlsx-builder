package main

import (
	"xlsx-builder/delivery/http"
	"xlsx-builder/internal/config"
)

func main() {
	cfg := config.Read()
	mux := http.NewRouter()
	srv := http.NewServer(cfg.HTTPAddr, mux)

	srv.Start()
}
