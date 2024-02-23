package publisher

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sfomuseum/go-pubsub-flamework"
	ps_publisher "github.com/sfomuseum/go-pubsub/publisher"
)

type ApiLogSQLPublisher struct {
	ps_publisher.Publisher
	db *sql.DB
}

func init(){
	ctx := context.Background()
	ps_publisher.RegisterPublisher(ctx, "apilog_sql", NewApiLogSQLPublisher)
}

func NewApiLogSQLPublisher(ctx context.Context, uri string) (ps_publisher.Publisher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	engine := u.Host
	dsn := q.Get("dsn")

	db, err := sql.Open(engine, dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to open database, %w", err)
	}

	p := &ApiLogSQLPublisher{
		db: db,
	}

	return p, nil
}

func (p *ApiLogSQLPublisher) Publish(ctx context.Context, msg string) error {

	var api_log *flamework.ApiLog

	err := json.Unmarshal([]byte(msg), &api_log)

	if err != nil {
		return fmt.Errorf("Failed to parse log, %w", err)
	}

	return nil
}

func (p *ApiLogSQLPublisher) Close() error {
	return p.db.Close()
}
