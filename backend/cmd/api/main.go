package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	secretsadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/secretsmanager"
	lambdaadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	_ "github.com/go-sql-driver/mysql"
)

const ErrorStatusCode = 500

func createErrorRes(err error) events.APIGatewayProxyResponse {
	slog.Error(err.Error())
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
	account, err := secretsadapter.FetchDBAccount(ctx)
	if err != nil {
		return createErrorRes(err), nil
	}
	slog.Info("wrapper: context ", "ctx", ctx)
	dbadapter, _ := rdbadapter.NewMySQL(&rdbadapter.MySQLConfig{
		USERNAME: account.UserName,
		PASSWORD: account.Password,
		HOST:     os.Getenv("EN_DB_HOST"),
		PORT:     os.Getenv("EN_DB_PORT"),
		NAME:     os.Getenv("EN_DB_NAME"),
	})
	defer dbadapter.Conn.Close() //nolint:errcheck
	service := InitializeEnAPIService(dbadapter)
	server, _ := ogen.NewServer(service)
	return lambdaadapter.NewAPIGatewayHandler(server).Run(ctx, event) //nolint:wrapcheck
}

func main() {
	lambda.Start(wrapper)
}
