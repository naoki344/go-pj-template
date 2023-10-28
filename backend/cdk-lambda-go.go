package main

import (
	"os"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	awslambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkLambdaGoStackProps struct {
	awscdk.StackProps
}

func NewCdkLambdaGoStack(scope constructs.Construct, id string, props *CdkLambdaGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	function := awslambdago.NewGoFunction(stack, jsii.String("handler"), &awslambdago.GoFunctionProps{
		Entry: jsii.String("lambda"),
		Description:  jsii.String("A function written in Go"),
		MemorySize:   jsii.Number(512),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
		Architecture: awslambda.Architecture_ARM_64(),
		FunctionName: jsii.String("test-todos-handler"),
		Environment: &map[string]*string{
			"LOG_LEVEL": jsii.String(os.Getenv("LOG_LEVEL")),
			"ENV":       jsii.String(os.Getenv("ENV")),
		},
	})

	restapi := awsapigateway.NewRestApi(stack, jsii.String("TestAPI"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("test-api-gateway"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
			StatusCode:   jsii.Number(200),
		},
	})
	apiResource := restapi.
		Root().
		AddResource(jsii.String("notes"), nil)
	apiResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	apiResource.AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todoIdResource := apiResource.AddResource(jsii.String("{noteID}"), nil)
	todoIdResource.AddMethod(jsii.String("get"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todoIdResource.AddMethod(jsii.String("PUT"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todoIdResource.AddMethod(jsii.String("DELETE"), awsapigateway.NewLambdaIntegration(function, nil), nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkLambdaGoStack(app, "CdkLambdaGoStack", &CdkLambdaGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_ACCOUNT_ID")),
		Region:  jsii.String(os.Getenv("AWS_REGION")),
	}
}
