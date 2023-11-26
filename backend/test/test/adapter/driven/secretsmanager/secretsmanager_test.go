package secretsmanager_test

import (
	"context"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	secretsadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/secretsmanager"
	secretsMock "github.com/naoki344/go-pj-template/backend/test/mock/adapter/driven/secretsmanager"
	"github.com/stretchr/testify/assert"
)

func TestSecretsManagerAdapter_GetDBAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	DBSecretID := os.Getenv("EN_DB_SECRET_NAME")

	mock := secretsMock.NewMockSecretsManagerClientInterface(ctrl)
	secretString := "{\"username\":\"account_test\",\"password\":\"testpassword\"}" //nolint:gosec
	mock.EXPECT().GetSecretStringWithContext(ctx, DBSecretID).Return(secretString, nil)

	type fields struct {
		Client secretsadapter.SecretsManagerClientInterface
	}
	type args struct {
		ctx context.Context
	}
	expectAccount := secretsadapter.DBAccount{
		UserName: "account_test",
		Password: "testpassword",
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *secretsadapter.DBAccount
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "adapter/secretsmanager GetDBAccount Test - success",
			fields: fields{
				Client: mock,
			},
			args: args{
				ctx: ctx,
			},
			want:      &expectAccount,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &secretsadapter.SecretsManagerAdapter{
				Client: tt.fields.Client,
			}
			got, err := adapter.GetDBAccount(tt.args.ctx)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
