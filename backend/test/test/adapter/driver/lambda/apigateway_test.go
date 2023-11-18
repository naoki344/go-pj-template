package lambdaadapter

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	apigatewayadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/lambda"
	"github.com/stretchr/testify/assert"
)

type HandlerMock struct {
	resBody string
}

func (handler *HandlerMock) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(handler.resBody))
}

func TestNewAPIGatewayHandler(t *testing.T) {
	serverMock := &HandlerMock{"message"}
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

	type fields struct {
		server http.Handler
	}
	type args struct {
		ctx   context.Context
		event events.APIGatewayProxyRequest
	}
	body := "body message"
	serverMock := &HandlerMock{body}
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
				event: events.APIGatewayProxyRequest{},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				MultiValueHeaders: map[string][]string{
					"Content-Type": {"text/plain; charset=utf-8"},
				},
				Body:            body,
				IsBase64Encoded: false,
			},
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
