package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/etilite/xlsx-builder/internal/config"
	"github.com/etilite/xlsx-builder/internal/delivery/http"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	realMain(ctx)
}

func realMain(ctx context.Context) {
	cfg := config.Read()
	mux := http.NewRouter()
	srv := http.NewServer(cfg.HTTPAddr, mux)

	srv.Run(ctx)
}
