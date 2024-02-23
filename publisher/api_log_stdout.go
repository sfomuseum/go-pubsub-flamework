package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sfomuseum/go-pubsub-flamework"
	ps_publisher "github.com/sfomuseum/go-pubsub/publisher"
)

type ApiLogStdoutPublisher struct {
	ps_publisher.Publisher
}

func init(){
	ctx := context.Background()
	ps_publisher.RegisterPublisher(ctx, "apilog_stdout", NewApiLogStdoutPublisher)
}

func NewApiLogStdoutPublisher(ctx context.Context, uri string) (ps_publisher.Publisher, error) {
	p := &ApiLogStdoutPublisher{}
	return p, nil
}

func (p *ApiLogStdoutPublisher) Publish(ctx context.Context, msg string) error {

	l, err := flamework.UnmarshalApiLog(msg)

	if err != nil {
		return fmt.Errorf("Failed to decode log message, %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(l)

	if err != nil {
		return fmt.Errorf("Failed to encode log message, %w", err)
	}

	return nil
}

func (p *ApiLogStdoutPublisher) Close() error {
	return nil
}
