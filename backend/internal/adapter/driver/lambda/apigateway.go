package lambdaadapter

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

type APIGatewayHandler struct {
	server http.Handler
}

func NewAPIGatewayHandler(
	server http.Handler,
) *APIGatewayHandler {
	return &APIGatewayHandler{
		server: server,
	}
}

func (handler *APIGatewayHandler) SetLogger(ctx context.Context) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		logger.With(slog.String("aws_request_id", lc.AwsRequestID))
	}
}

func (handler *APIGatewayHandler) Run(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	handler.SetLogger(ctx)

	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	return httpadapter.New(handler.server).ProxyWithContext(ctx, event) //nolint:wrapcheck
}
