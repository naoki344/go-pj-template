package secretsmanager_test

import (
	"context"
	"os"
	"testing"

	secretsadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/secretsmanager"
	secretsMock "github.com/g-stayfresh/en/backend/test/mock/adapter/driven/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSecretsManagerAdapter_GetDBAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	DBSecretID := os.Getenv("EN_DB_SECRET_NAME")

	mock := secretsMock.NewMockSecretsManagerClientInterface(ctrl)
	secretString := "{\"ARN\":\"arn:aws:secretsmanager:ap-northeast-1:376145842201:secret:rds!db-7849d76f-fdb2-4bdb-8dd3-c02f83c6c7d6-ck3gqS\",\"CreatedDate\":\"2023-11-11T10:10:38.495Z\",\"Name\":\"rds!db-7849d76f-fdb2-4bdb-8dd3-c02f83c6c7d6\",\"SecretBinary\":null,\"SecretString\":\"{\\\"username\\\":\\\"enstayfresh_test\\\",\\\"password\\\":\\\"testpassword\\\"}\",\"VersionId\":\"2e57d542-703d-4969-b5de-7553ddf09a53\",\"VersionStages\":[\"AWSCURRENT\",\"AWSPENDING\"],\"ResultMetadata\":{}}"
	mock.EXPECT().GetSecret(ctx, DBSecretID).Return(&secretString, nil)

	type fields struct {
		Client secretsadapter.SecretsManagerClientInterface
	}
	type args struct {
		ctx context.Context
	}
	expectAccount := secretsadapter.DBAccount{
		UserName: "enstayfresh_test",
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
