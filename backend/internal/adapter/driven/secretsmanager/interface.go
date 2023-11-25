package secretsmanager

import (
	"context"

	"github.com/cockroachdb/errors"
)

var ErrDBAccountFetchFailure = errors.New("db account fetch failure")

type SecretsManagerClientInterface interface {
	GetSecret(ctx context.Context, secretsID string) (*string, error)
}

type SecretsManagerAdapterInterface interface {
	GetDBAccount(ctx context.Context) DBAccount
}

type DBAccount struct {
	UserName string `json:"username"` //nolint: tagliatelle
	Password string `json:"password"` //nolint: tagliatelle
}

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
