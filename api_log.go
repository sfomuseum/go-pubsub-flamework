package flamework

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type ApiLog struct {
	Created         int64                  `json:"created"`
	Method          string                 `json:"method"`
	Hostname        string                 `json:"hostname"`
	PID             int                    `json:"pid"`
	RemoteAddr      string                 `json:"remote_addr"`
	AccessTokenHash string                 `json:"access_token_hash"`
	AuthTokenId     int64                  `json:"auth_token_id"`
	ApiKeyId        int64                  `json:"api_key_id"`
	Stat            string                 `json:"stat"`
	Error           *ApiError               `json:"error,omitempty"`
	Params          map[string]interface{} `json:"params,omitempty"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UnmarshalApiLog(msg string) (*ApiLog, error) {

	var log *ApiLog

	err := json.Unmarshal([]byte(msg), &log)

	if err != nil {
	   slog.Info("Failed to unmarshal message", "message", msg)
		return nil, fmt.Errorf("Failed to unmarshal message, %w", err)
	}

	return log, nil
}
