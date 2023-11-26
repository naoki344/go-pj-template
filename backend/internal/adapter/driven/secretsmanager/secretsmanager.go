package secretsmanager

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

type SecretsManagerAdapter struct {
	Client SecretsManagerClientInterface
}

func NewSecretsManagerAdapter(client SecretsManagerClientInterface) *SecretsManagerAdapter {
	return &SecretsManagerAdapter{
		Client: client,
	}
}

func NewClient() (*secretcache.Cache, error) {
	config := secretcache.CacheConfig{
		MaxCacheSize: secretcache.DefaultMaxCacheSize,
		VersionStage: secretcache.DefaultVersionStage,
		CacheItemTTL: secretcache.DefaultCacheItemTTL,
	}
	client, err := secretcache.New(
		func(c *secretcache.Cache) { c.CacheConfig = config },
	)
	if err != nil {
		slog.Error("Fetch db account secrets error", slog.Any("error", err))
		return nil, ErrDBAccountFetchFailure
	}
	return client, nil
}

func (adapter *SecretsManagerAdapter) GetDBAccount(ctx context.Context) (*DBAccount, error) {
	DBSecretID := os.Getenv("EN_DB_SECRET_NAME")
	secretString, err := adapter.Client.GetSecretStringWithContext(ctx, DBSecretID)
	if err != nil {
		slog.Error("Fetch db account secrets error", slog.Any("error", err))
		return nil, ErrDBAccountFetchFailure
	}
	slog.Info("secrets string", slog.Any("result_string", secretString))

	account := DBAccount{}
	jsonErr2 := json.Unmarshal([]byte(secretString), &account)
	if jsonErr2 != nil {
		slog.Error(err.Error())
		return nil, ErrDBAccountFetchFailure
	}
	return &account, nil
}
