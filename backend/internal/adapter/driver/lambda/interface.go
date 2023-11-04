package lambdaadapter


type LambdaHandler interface {
	Run() (interface{}, error)
}
