package lambdaadapter

type FunctionType string

const (
	FunctionTypeAPIGateway = FunctionType("APIGateway")
)

type Lambda struct {
	AWSRequestID string
	FunctionName string
	FunctionType FunctionType
}

type LambdaHandler interface {
	Run() (interface{}, error)
}
