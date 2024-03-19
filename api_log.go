package flamework

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type ApiLog struct {
	Id                int64                  `json:"id"`
	ApiKeyId          int64                  `json:"api_key_id"`
	ApiKeyUserId      int64                  `json:"api_key_user_id"`
	ApiKeyRoleId      uint8                  `json:"api_key_role_id"`
	AccessTokenId     int64                  `json:"access_token_id"`
	AccessTokenUserId int64                  `json:"access_token_user_id"`
	AccessTokenHash   string                 `json:"access_token_hash"`
	Method            string                 `json:"method"`
	Hostname          string                 `json:"hostname"`
	RemoteAddr        int                    `json:"remote_addr"`
	Stat              uint8                  `json:"stat"`
	Error             *ApiError              `json:"error,omitempty"`
	Params            map[string]interface{} `json:"params,omitempty"`
	Created           int64                  `json:"created"`
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
