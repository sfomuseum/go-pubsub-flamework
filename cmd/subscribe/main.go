package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/sfomuseum/go-pubsub-flamework/publisher"
	"github.com/sfomuseum/go-pubsub/app/subscribe"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := subscribe.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to run subscribe tool", "error", err)
		os.Exit(1)
	}
}
