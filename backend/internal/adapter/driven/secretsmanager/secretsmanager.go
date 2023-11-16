package secretsmanager

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBAccount struct {
	UserName string `json:"username"` //nolint: tagliatelle
	Password string `json:"password"` //nolint: tagliatelle
}

var ErrDBAccountFetchFailure = errors.New("db account fetch failure")

type secretInfo struct {
	ARN            string   `json:"ARN"`            //nolint: tagliatelle
	CreatedDate    string   `json:"CreatedDate"`    //nolint: tagliatelle
	Name           string   `json:"Name"`           //nolint: tagliatelle
	SecretBinary   []byte   `json:"SecretBinary"`   //nolint: tagliatelle
	SecretString   string   `json:"SecretString"`   //nolint: tagliatelle
	VersionID      string   `json:"VersionId"`      //nolint: tagliatelle
	VersionStages  []string `json:"VersionStages"`  //nolint: tagliatelle
	ResultMetadata struct{} `json:"ResultMetadata"` //nolint: tagliatelle
}

func FetchDBAccount(ctx context.Context) (*DBAccount, error) {
	secretID := os.Getenv("EN_DB_SECRET_NAME")
	url := "http://localhost:2773/secretsmanager/get?secretId=" + secretID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.Error(
			"can not create http-get request.", slog.Any("err", err))
		return nil, ErrDBAccountFetchFailure
	}

	secretToken := os.Getenv("AWS_SESSION_TOKEN")
	req.Header.Set("X-Aws-Parameters-Secrets-Token", secretToken)
	client := &http.Client{}

	resp, err := client.Do(req)
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
	secret := secretInfo{}
	jsonErr := json.Unmarshal([]byte(secretString), &secret)
	if jsonErr != nil {
		slog.Error(err.Error())
		return nil, ErrDBAccountFetchFailure
	}
	account := DBAccount{}
	jsonErr2 := json.Unmarshal([]byte(secret.SecretString), &account)
	if jsonErr2 != nil {
		slog.Error(err.Error())
		return nil, ErrDBAccountFetchFailure
	}
	return &account, nil
}
