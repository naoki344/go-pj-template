package main

import (
	"os"
	"context"
	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	"github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	"github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"log/slog"
	_ "github.com/go-sql-driver/mysql"
)


func wrapper(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dbadapter, _ := rdbadapter.NewMySQL(&rdbadapter.MySQLConfig{
			USERNAME: os.Getenv("DB_USERNAME"),
			PASSWORD: os.Getenv("DB_PASSWORD"),
			HOST: os.Getenv("DB_HOST"),
			PORT: os.Getenv("DB_PORT"),
			NAME: os.Getenv("DB_NAME"),
		})
	slog.Info("hello", "count", 3)
	defer dbadapter.Conn.Close()
	service := InitializeEnAPIService(dbadapter)
	server, _ := ogen.NewServer(service)
	return lambdaadapter.NewAPIGatewayHandler(ctx, event, server).Run()
}


func main() {
	lambda.Start(wrapper)
}
