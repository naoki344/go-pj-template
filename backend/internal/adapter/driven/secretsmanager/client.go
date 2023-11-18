package secretsmanager

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type SecretsManagerClient struct {
}

func (client *SecretsManagerClient) GetSecret(ctx context.Context, secretsID string) (*string, error) {
	url := "http://localhost:2773/secretsmanager/get?secretId=" + secretsID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.Error(
			"can not create http-get request.", slog.Any("err", err))
		return nil, ErrDBAccountFetchFailure
	}

	secretToken := os.Getenv("AWS_SESSION_TOKEN")
	req.Header.Set("X-Aws-Parameters-Secrets-Token", secretToken)
	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		slog.Error("Fetch secretsmanager error", slog.Any("error", err))
		return nil, ErrDBAccountFetchFailure
	}
	defer resp.Body.Close() //nolint:errcheck

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Fetch secretsmanager error", slog.Any("error", err))
		return nil, ErrDBAccountFetchFailure
	}
	secretString := string(body)
	return &secretString, nil

}
