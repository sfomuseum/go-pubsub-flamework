package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/sfomuseum/go-pubsub/app/subscribe"
	_ "github.com/sfomuseum/go-pubsub-flamework/publisher"
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
