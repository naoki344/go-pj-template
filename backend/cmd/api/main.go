package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	lambdaadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slog.Info("wrapper: context ", "ctx", ctx)
	dbadapter, _ := rdbadapter.NewMySQL(&rdbadapter.MySQLConfig{
		USERNAME: os.Getenv("DB_USERNAME"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		NAME:     os.Getenv("DB_NAME"),
	})
	defer dbadapter.Conn.Close() //nolint:errcheck
	service := InitializeEnAPIService(dbadapter)
	server, _ := ogen.NewServer(service)
	return lambdaadapter.NewAPIGatewayHandler(server).Run(ctx, event) //nolint:wrapcheck
}

func main() {
	lambda.Start(wrapper)
}
