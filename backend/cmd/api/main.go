package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	secretsadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/secretsmanager"
	lambdaadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	"github.com/g-stayfresh/en/backend/pkg/monitoring"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ErrorStatusCode = 500
	sentryTimeout   = 2 * time.Second
)

func createErrorRes(err error) events.APIGatewayProxyResponse {
	monitoring.ErrLoggingWithSentry(err)
	return events.APIGatewayProxyResponse{
		StatusCode: ErrorStatusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body: `{"type": "InternalServerError", "message": "unexpected error has occurred."}`,
	}
}

func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	monitoring.SetupSentry()
	client := &secretsadapter.SecretsManagerClient{}
	secretsAdapter := secretsadapter.NewSecretsManagerAdapter(client)
	account, err := secretsAdapter.GetDBAccount(ctx)
	if err != nil {
		return createErrorRes(err), nil
	}
	dbadapter, err := rdbadapter.NewMySQL(&rdbadapter.MySQLConfig{
		USERNAME: account.UserName,
		PASSWORD: account.Password,
		HOST:     os.Getenv("EN_DB_HOST"),
		PORT:     os.Getenv("EN_DB_PORT"),
		NAME:     os.Getenv("EN_DB_NAME"),
	})
	if err != nil {
		return createErrorRes(err), nil
	}
	defer dbadapter.Conn.Close() //nolint:errcheck
	service := InitializeEnAPIService(dbadapter)
	server, err := ogen.NewServer(service)
	if err != nil {
		return createErrorRes(err), nil
	}
	return lambdaadapter.NewAPIGatewayHandler(server).Run(ctx, event) //nolint:wrapcheck
}

func main() {
	lambda.Start(wrapper)
}
