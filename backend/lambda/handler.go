package main

import (
	"os"
	"cdk-lambda-go/ogen"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	cfg := &Config{
			USERNAME: os.Getenv("DB_USERNAME"),
			PASSWORD: os.Getenv("DB_PASSWORD"),
			HOST: os.Getenv("DB_HOST"),
			PORT: os.Getenv("DB_PORT"),
			NAME: os.Getenv("DB_NAME"),
		}
	service := InitializeNoteService(cfg)
	defer service.dbRepository.db.conn.Close()
	s, _ := ogen.NewServer(service)
	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	lambda.Start(httpadapter.New(s).ProxyWithContext)
}
