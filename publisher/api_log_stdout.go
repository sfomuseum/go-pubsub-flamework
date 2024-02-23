package publisher

// ./bin/subscribe -subscriber-uri redis://api_log -publisher-uri apilogstdout://
// 2024/02/23 02:29:01 INFO Listening for messages
// {"created":1708655348,"method":"GET","hostname":"ip-10-26-152-44.us-west-2.compute.internal","pid":409857,"remote_addr":"64.252.73.88","access_token_hash":"269a2b4df5bd415c66bc3f2ed92df9f7dd0e147d","auth_token_id":1367947,"api_key_id":1367945,"stat":"ok"}
//
// and so on...

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
	ps_publisher.RegisterPublisher(ctx, "apilogstdout", NewApiLogStdoutPublisher)
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
