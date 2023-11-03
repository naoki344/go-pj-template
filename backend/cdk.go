package main

import (
	"os"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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

	myRole := awsiam.NewRole(stack, jsii.String("MyLambdaRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})

	function := awslambdago.NewGoFunction(stack, jsii.String("handler"), &awslambdago.GoFunctionProps{
		Entry:        jsii.String("cmd/api"),
		Description:  jsii.String("A function written in Go"),
		MemorySize:   jsii.Number(512),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
		Architecture: awslambda.Architecture_ARM_64(),
		FunctionName: jsii.String("test-todos-handler"),
		Environment: &map[string]*string{
			"LOG_LEVEL":   jsii.String(os.Getenv("LOG_LEVEL")),
			"ENV":         jsii.String(os.Getenv("ENV")),
			"DB_USERNAME": jsii.String(os.Getenv("DB_USERNAME")),
			"DB_PASSWORD": jsii.String(os.Getenv("DB_PASSWORD")),
			"DB_HOST":     jsii.String(os.Getenv("DB_HOST")),
			"DB_PORT":     jsii.String(os.Getenv("DB_PORT")),
			"DB_NAME":     jsii.String(os.Getenv("DB_NAME")),
		},
		Role: myRole,
	})
	dbPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("rds:*"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	})
	myRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))
	myRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaVPCAccessExecutionRole")))
	myRole.AddToPolicy(dbPolicy)


	restApi := awsapigateway.NewRestApi(stack, jsii.String("TestAPI"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("test-api-gateway"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
			StatusCode:   jsii.Number(200),
		},
	})



	cognitoUserPool := awscognito.NewUserPool(stack, jsii.String("en-userpool"), &awscognito.UserPoolProps{
		UserPoolName: jsii.String("myawesomeapp-userpool"),
		SignInCaseSensitive: jsii.Bool(false),
	})
	userPoolClient := cognitoUserPool.AddClient(jsii.String("en-userpool-client"), &awscognito.UserPoolClientOptions{
		GenerateSecret: jsii.Bool(false),
		AuthFlows: &awscognito.AuthFlow{
			AdminUserPassword: jsii.Bool(true),
			UserPassword: jsii.Bool(true),
			UserSrp: jsii.Bool(true),
		},
		SupportedIdentityProviders: &[]awscognito.UserPoolClientIdentityProvider{
			awscognito.UserPoolClientIdentityProvider_COGNITO(),
		},
	})
	identityPool := awscognito.NewCfnIdentityPool(stack, jsii.String("en-cognito-id-pool"), &awscognito.CfnIdentityPoolProps{
		IdentityPoolName: jsii.String("en-cognito-id-pool"),
		AllowUnauthenticatedIdentities: jsii.Bool(false),
		CognitoIdentityProviders: &[]*awscognito.CfnIdentityPool_CognitoIdentityProviderProperty{
			{
				ClientId: userPoolClient.UserPoolClientId(),
				ProviderName: cognitoUserPool.UserPoolProviderName(),
			},
		},
	})

	// IAM Role for Cognito Identity Pool
	identityPoolRole := awsiam.NewRole(stack, jsii.String("CognitoUserIdentityPoolRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewFederatedPrincipal(
			jsii.String("cognito-identity.amazonaws.com"),
			&map[string]interface{}{
				"StringEquals": map[string]interface{}{
					"cognito-identity.amazonaws.com:aud": identityPool.Ref(),
				},
				"ForAnyValue:StringLike": map[string]interface{}{
					"cognito-identity.amazonaws.com:amr": "authenticated",
				},
			},
			jsii.String("sts:AssumeRoleWithWebIdentity")),
	})

	// Cognito Identity Pool Role Attachment
	awscognito.NewCfnIdentityPoolRoleAttachment(
		stack, jsii.String("CognitoUserIdentityPoolRoleAttachment"), &awscognito.CfnIdentityPoolRoleAttachmentProps{
			IdentityPoolId: identityPool.Ref(),
			Roles: &map[string]interface{}{
				"authenticated": identityPoolRole.RoleArn(),
			},
	})

	// Custom policy statements
	identityPoolRole.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("mobileanalytics:PutEvents"),
			jsii.String("cognito-sync:*"),
			jsii.String("cognito-identity:*"),
		},
		Resources: &[]*string{jsii.String("*")},
	}))

	identityPoolRole.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("execute-api:Invoke"),
		},
		Resources: &[]*string{
			restApi.ArnForExecuteApi(
				jsii.String("*"),
				jsii.String("/*"),
				jsii.String("*"),
			),
		},
	}))


	apiResource := restApi.
		Root().
		AddResource(jsii.String("customers"), nil)
	apiResource.AddMethod(
		jsii.String("GET"), awsapigateway.NewLambdaIntegration(function, nil),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
	})
	apiResource.AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(function, nil), nil)

	customerIdResource := apiResource.AddResource(jsii.String("{customerID}"), nil)
	customerIdResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(function, nil), nil)

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
