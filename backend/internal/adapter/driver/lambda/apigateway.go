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
	Server http.Handler
}

func NewAPIGatewayHandler(
	server http.Handler,
) *APIGatewayHandler {
	return &APIGatewayHandler{
		Server: server,
	}
}

type APIGateway struct {
	Resource            string
	Path                string
	Method              string
	APIGatewayRequestID string
	APIID               string
	APIStage            string
}

func (handler *APIGatewayHandler) SetLogger(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) {
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
			FunctionType: FunctionTypeAPIGateway,
		}),
		slog.Any(string(FunctionTypeAPIGateway), APIGateway{
			Resource:            event.Resource,
			Path:                event.Path,
			Method:              event.HTTPMethod,
			APIGatewayRequestID: event.RequestContext.RequestID,
			APIID:               event.RequestContext.APIID,
			APIStage:            event.RequestContext.Stage,
		}),
	)
	slog.SetDefault(slogger)
	slog.Info("log settings success.", slog.Any("LambdaContext", lc))
}

func (handler *APIGatewayHandler) Run(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	handler.SetLogger(ctx, event)
	slog.Debug("apigatewayEvent", slog.Any("event", event))

	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	return httpadapter.New(handler.Server).ProxyWithContext(ctx, event) //nolint:wrapcheck
}
