package publisher

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/sfomuseum/go-pubsub-flamework"
	ps_publisher "github.com/sfomuseum/go-pubsub/publisher"
)

const API_LOGS_TABLE_NAME string = "ApiLogs"

type ApiLogSQLPublisher struct {
	ps_publisher.Publisher
	db *sql.DB
}

func init() {
	ctx := context.Background()
	ps_publisher.RegisterPublisher(ctx, "apilogsql", NewApiLogSQLPublisher)
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

	logger := slog.Default()

	var api_log *flamework.ApiLog

	err := json.Unmarshal([]byte(msg), &api_log)

	if err != nil {
		logger.Error("Failed to parse log message", "message", msg, "error", err)
		return fmt.Errorf("Failed to parse log, %w", err)
	}

	logger = logger.With("id", api_log.Id)

	enc_params, err := json.Marshal(api_log.Params)

	if err != nil {
		logger.Error("Failed to marshal log parameters", "error", err)
		return fmt.Errorf("Failed to marshal parameters, %w", err)
	}

	enc_error, err := json.Marshal(api_log.Error)

	if err != nil {
		logger.Error("Failed to marshal log error", "error", err)
		return fmt.Errorf("Failed to marshal error, %w", err)
	}

	q := fmt.Sprintf("INSERT INTO %s (id, api_key_id, api_key_user_id, access_token_id, access_token_user_id, access_token_hash, remote_addr, hostname, method, params, stat, error, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", API_LOGS_TABLE_NAME)

	_, err = p.db.ExecContext(ctx, q, api_log.Id, api_log.ApiKeyId, api_log.ApiKeyUserId, api_log.AccessTokenId, api_log.AccessTokenUserId, api_log.AccessTokenHash, api_log.RemoteAddr, api_log.Hostname, api_log.Method, string(enc_params), api_log.Stat, string(enc_error), api_log.Created)

	if err != nil {
		logger.Error("Failed to record log message", "error", err)
		return fmt.Errorf("Failed to insert API log, %w", err)
	}

	logger.Debug("Wrote API log message")
	return nil
}

func (p *ApiLogSQLPublisher) Close() error {
	return p.db.Close()
}
