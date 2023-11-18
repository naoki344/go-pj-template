package lambdaadapter

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	apigatewayadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	"github.com/stretchr/testify/assert"
)

type HandlerMock struct{}

func (handler *HandlerMock) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestNewAPIGatewayHandler(t *testing.T) {
	serverMock := &HandlerMock{}
	type args struct {
		Server http.Handler
	}
	tests := []struct {
		name string
		args args
		want *apigatewayadapter.APIGatewayHandler
	}{
		{
			name: "adapter/lambda NewAPIGatewayHandler Test - success",
			args: args{
				Server: serverMock,
			},
			want: &apigatewayadapter.APIGatewayHandler{
				Server: serverMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, apigatewayadapter.NewAPIGatewayHandler(tt.args.Server))
		})
	}
}

func TestAPIGatewayHandler_Run(t *testing.T) {
	reqContext := events.APIGatewayProxyRequestContext{
		AccountID:         "111122222",
		ResourceID:        "111122222",
		OperationName:     "111122222",
		Stage:             "111122222",
		DomainName:        "111122222",
		DomainPrefix:      "111122222",
		RequestID:         "111122222",
		ExtendedRequestID: "111122222",
		Protocol:          "111122222",
		Identity:          events.APIGatewayRequestIdentity{},
		ResourcePath:      "111122222",
		Path:              "111122222",
		HTTPMethod:        "GET",
		RequestTime:       "111122222",
		RequestTimeEpoch:  int64(11111111111),
		APIID:             "111122222",
	}
	event := events.APIGatewayProxyRequest{
		Resource:   "resource",
		Path:       "/path/string",
		HTTPMethod: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		MultiValueHeaders: map[string][]string{
			"Content-Type": {"application/json"},
		},
		QueryStringParameters: map[string]string{
			"Params1": "paramsValue",
		},
		MultiValueQueryStringParameters: map[string][]string{
			"Content-Type": {"application/json"},
		},
		PathParameters: map[string]string{
			"Params1": "paramsValue",
		},
		StageVariables: map[string]string{
			"Params1": "paramsValue",
		},
		RequestContext:  reqContext,
		Body:            "body value",
		IsBase64Encoded: false,
	}

	type fields struct {
		server http.Handler
	}
	type args struct {
		ctx   context.Context
		event events.APIGatewayProxyRequest
	}
	serverMock := &HandlerMock{}
	ctx := context.Background()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      events.APIGatewayProxyResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "adapter/lambda NewAPIGatewayHandler Test - success",
			fields: fields{
				server: serverMock,
			},
			args: args{
				ctx:   ctx,
				event: event,
			},
			want:      events.APIGatewayProxyResponse{StatusCode: 200},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &apigatewayadapter.APIGatewayHandler{
				Server: tt.fields.server,
			}
			got, err := handler.Run(tt.args.ctx, tt.args.event)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
