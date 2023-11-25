package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cockroachdb/errors"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	secretsadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/secretsmanager"
	lambdaadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ErrorStatusCode = 500
	sentryTimeout   = 2 * time.Second
)

func createErrorRes(err error) events.APIGatewayProxyResponse {
	logging(err)
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

func setupSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://b3576edb04c135b90f90c0d5601dc361@o4506284127027200.ingest.sentry.io/4506284129386496",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		slog.Error("sentry.Init: %s", err)
	}
}

func logging(err error) {
	slog.Error(
		err.Error(),
		slog.Any("trace", errors.GetAllSafeDetails(err)),
		slog.Any("hints", errors.GetAllHints(err)),
	)
	defer sentry.Flush(sentryTimeout)
	sentry.CaptureException(err)
}

func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	setupSentry()
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
