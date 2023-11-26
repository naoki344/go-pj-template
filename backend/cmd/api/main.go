package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	rdbadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/rdb"
	secretsadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/secretsmanager"
	lambdaadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driver/lambda"
	ogen "github.com/naoki344/go-pj-template/backend/internal/adapter/driver/ogenlib"
	"github.com/naoki344/go-pj-template/backend/pkg/monitoring"
)

const (
	ErrorStatusCode = 500
	sentryTimeout   = 2 * time.Second
)

var Client, _ = secretsadapter.NewClient() //nolint:gochecknoglobals

func createErrorResWithLogging(err error) events.APIGatewayProxyResponse {
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

func createDBAdapter(ctx context.Context) (*rdbadapter.MySQL, error) {
	secretsAdapter := secretsadapter.NewSecretsManagerAdapter(Client)
	account, err := secretsAdapter.GetDBAccount(ctx)
	if err != nil {
		return nil, err
	}
	dbadapter, err := rdbadapter.NewMySQL(&rdbadapter.MySQLConfig{
		USERNAME: account.UserName,
		PASSWORD: account.Password,
		HOST:     os.Getenv("EN_DB_HOST"),
		PORT:     os.Getenv("EN_DB_PORT"),
		NAME:     os.Getenv("EN_DB_NAME"),
	})
	if err != nil {
		return nil, err
	}
	return dbadapter, nil
}

func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	monitoring.SetupSentry()
	dbadapter, err := createDBAdapter(ctx)
	if err != nil {
		return createErrorResWithLogging(err), nil
	}
	defer dbadapter.Conn.Close() //nolint:errcheck
	service := InitializeAPIService(dbadapter)
	server, err := ogen.NewServer(service)
	if err != nil {
		return createErrorResWithLogging(err), nil
	}
	return lambdaadapter.NewAPIGatewayHandler(server).Run(ctx, event) //nolint:wrapcheck
}

func main() {
	lambda.Start(wrapper)
}
