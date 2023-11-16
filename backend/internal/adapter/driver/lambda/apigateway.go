package lambdaadapter

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"

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

type Lambda struct {
	AWSRequestID string
	FunctionName string
}

func (handler *APIGatewayHandler) SetLogger(ctx context.Context) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		slog.Error("error log settings failure.",
			slog.Any("lambdacontext", lc))
		return
	}
	ops := slog.HandlerOptions{AddSource: true}
	// DEBUG時にはログ出力レベルを変更する
	if strings.EqualFold(os.Getenv("LOG_LEVEL"), "debug") {
		ops.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &ops))
	slogger := logger.With(
		slog.Any("Lambda", Lambda{
			AWSRequestID: lc.AwsRequestID,
			FunctionName: lc.InvokedFunctionArn,
		}),
	)
	slog.SetDefault(slogger)
	slog.Info("log settings success.", slog.Any("context", lc))
}

func (handler *APIGatewayHandler) Run(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	handler.SetLogger(ctx)

	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	return httpadapter.New(handler.server).ProxyWithContext(ctx, event) //nolint:wrapcheck
}
