package secretsmanager

import (
	"context"

	"github.com/cockroachdb/errors"
)

var ErrDBAccountFetchFailure = errors.New("db account fetch failure")

type SecretsManagerClientInterface interface {
	GetSecretStringWithContext(ctx context.Context, secretsID string) (string, error)
}

type SecretsManagerAdapterInterface interface {
	GetPrimaryDBAccount(ctx context.Context) DBAccount
	GetSecondaryDBAccount(ctx context.Context) DBAccount
}

type DBAccount struct {
	UserName string `json:"username"` //nolint: tagliatelle
	Password string `json:"password"` //nolint: tagliatelle
}
