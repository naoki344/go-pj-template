package lambdaadapter

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"os"
	"log/slog"
	"dario.cat/mergo"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)


type ApiGatewayHandler struct {
	ctx context.Context
	event events.APIGatewayProxyRequest
	server http.Handler
}

func NewAPIGatewayHandler(ctx context.Context, event events.APIGatewayProxyRequest, server http.Handler) *ApiGatewayHandler{
	return &ApiGatewayHandler{
		ctx: ctx,
		event: event,
		server: server,
	}
}

func (handler *ApiGatewayHandler) SetLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	lc, ok := lambdacontext.FromContext(handler.ctx)
	if ok {
		logger = logger.With(slog.String("aws_request_id", lc.AwsRequestID))
	}
}


func (handler *ApiGatewayHandler) Run() (events.APIGatewayProxyResponse, error){
	handler.SetLogger()

	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	res, err := httpadapter.New(handler.server).ProxyWithContext(handler.ctx, handler.event)
	if err != nil {
		return res, err
	}
	newHeader := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Headers": "*",
	}
	mergo.Merge(&res.Headers, newHeader)
	return res, err
}
