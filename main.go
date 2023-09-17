package main

import (
	"github.com/etilite/xlsx-builder/delivery/http"
	"github.com/etilite/xlsx-builder/internal/config"
)

func main() {
	cfg := config.Read()
	mux := http.NewRouter()
	srv := http.NewServer(cfg.HTTPAddr, mux)

	srv.Start()
}
