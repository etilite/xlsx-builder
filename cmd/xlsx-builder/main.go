package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/etilite/xlsx-builder/internal/app"
	httpserver "github.com/etilite/xlsx-builder/internal/delivery/http"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	cfg := app.NewConfigFromEnv()

	server := httpserver.NewServer(cfg.HTTPAddr)

	if err := server.Run(ctx); err != nil {
		slog.Error("unable to start app", "error", err)
		panic(err)
	}
}
