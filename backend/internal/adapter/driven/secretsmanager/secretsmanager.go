package secretsmanager

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
)

type SecretsManagerAdapter struct {
	Client SecretsManagerClientInterface
}

func NewSecretsManagerAdapter(client SecretsManagerClientInterface) *SecretsManagerAdapter {
	return &SecretsManagerAdapter{
		Client: client,
	}
}

func (adapter *SecretsManagerAdapter) GetDBAccount(ctx context.Context) (*DBAccount, error) {
	DBSecretID := os.Getenv("EN_DB_SECRET_NAME")
	secretString, err := adapter.Client.GetSecret(ctx, DBSecretID)
	if err != nil {
		slog.Error("Fetch db account secrets error", slog.Any("error", err))
		return nil, ErrDBAccountFetchFailure
	}

	secret := secretInfo{}
	jsonErr := json.Unmarshal([]byte(*secretString), &secret)
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
