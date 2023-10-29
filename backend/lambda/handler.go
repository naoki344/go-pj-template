package main

import (
	"os"
	"log/slog"
	"context"
	"cdk-lambda-go/ogen"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	_ "github.com/go-sql-driver/mysql"
)

type Event struct {
	origin any
}


type Callback[T any] func(ctx context.Context, event *Event) (T, error)


func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		logger = logger.With(slog.String("aws_request_id", lc.AwsRequestID))
	}
	logger.Info("hello", "count", 3)
	cfg := &Config{
			USERNAME: os.Getenv("DB_USERNAME"),
			PASSWORD: os.Getenv("DB_PASSWORD"),
			HOST: os.Getenv("DB_HOST"),
			PORT: os.Getenv("DB_PORT"),
			NAME: os.Getenv("DB_NAME"),
		}
	service := InitializeEnAPIService(cfg, logger)
	defer service.dbRepository.db.conn.Close()
	s, _ := ogen.NewServer(service)
	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16

	return httpadapter.New(s).ProxyWithContext(ctx, event)
}


func main() {
	lambda.Start(wrapper)
}
