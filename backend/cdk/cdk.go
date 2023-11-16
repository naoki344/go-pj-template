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

const (
	lambdaMemorySize          = 512
	lambdaTimeout             = 28
	corsDefaultStatusCode     = 200
	paramsAndSecretsCacheSize = 500
	paramsAndSecretsLayerArn  = "arn:aws:lambda:ap-northeast-1:133490724326:layer:AWS-Parameters-and-Secrets-Lambda-Extension-Arm64:11" //nolint:gosec
)

type CdkLambdaGoStackProps struct {
	awscdk.StackProps
}

func NewCdkLambdaGoStack(
	scope constructs.Construct,
	stackName string,
	props *CdkLambdaGoStackProps,
) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &stackName, &sprops)

	myRole := awsiam.NewRole(stack, jsii.String("MyLambdaRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})
	layerArn := paramsAndSecretsLayerArn
	paramsAndSecrets := awslambda.ParamsAndSecretsLayerVersion_FromVersionArn(
		&layerArn, &awslambda.ParamsAndSecretsOptions{
			CacheSize: jsii.Number(paramsAndSecretsCacheSize),
			LogLevel:  awslambda.ParamsAndSecretsLogLevel(os.Getenv("LOG_LEVEL")),
		})

	function := awslambdago.NewGoFunction(stack, jsii.String("handler"), &awslambdago.GoFunctionProps{
		Entry:        jsii.String("../cmd/api"),
		Description:  jsii.String("A function written in Go"),
		MemorySize:   jsii.Number(lambdaMemorySize),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(lambdaTimeout)),
		Architecture: awslambda.Architecture_ARM_64(),
		FunctionName: jsii.String("test-todos-handler"),
		Environment: &map[string]*string{
			"LOG_LEVEL":         jsii.String(os.Getenv("LOG_LEVEL")),
			"ENV":               jsii.String(os.Getenv("ENV")),
			"EN_DB_HOST":        jsii.String(os.Getenv("EN_DB_HOST")),
			"EN_DB_PORT":        jsii.String(os.Getenv("EN_DB_PORT")),
			"EN_DB_NAME":        jsii.String(os.Getenv("EN_DB_NAME")),
			"EN_DB_SECRET_NAME": jsii.String(os.Getenv("EN_DB_SECRET_NAME")),
			"EN_AWS_REGION":     jsii.String(os.Getenv("EN_AWS_REGION")),
		},
		Role:             myRole,
		ParamsAndSecrets: paramsAndSecrets,
	})
	dbPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("rds:*"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	})
	secretsManagerPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("secretsmanager:GetSecretValue"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	})
	myRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))
	myRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaVPCAccessExecutionRole")))
	myRole.AddToPolicy(dbPolicy)
	myRole.AddToPolicy(secretsManagerPolicy)

	restAPI := awsapigateway.NewRestApi(stack, jsii.String("TestAPI"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("test-api-gateway"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
			StatusCode:   jsii.Number(corsDefaultStatusCode),
		},
	})

	cognitoUserPool := awscognito.NewUserPool(stack, jsii.String("en-userpool"), &awscognito.UserPoolProps{
		UserPoolName:        jsii.String("myawesomeapp-userpool"),
		SignInCaseSensitive: jsii.Bool(false),
	})
	userPoolClient := cognitoUserPool.AddClient(jsii.String("en-userpool-client"), &awscognito.UserPoolClientOptions{
		GenerateSecret: jsii.Bool(false),
		AuthFlows: &awscognito.AuthFlow{
			AdminUserPassword: jsii.Bool(true),
			UserPassword:      jsii.Bool(true),
			UserSrp:           jsii.Bool(true),
		},
		SupportedIdentityProviders: &[]awscognito.UserPoolClientIdentityProvider{
			awscognito.UserPoolClientIdentityProvider_COGNITO(),
		},
	})
	identityPool := awscognito.NewCfnIdentityPool(stack, jsii.String("en-cognito-id-pool"), &awscognito.CfnIdentityPoolProps{
		IdentityPoolName:               jsii.String("en-cognito-id-pool"),
		AllowUnauthenticatedIdentities: jsii.Bool(false),
		CognitoIdentityProviders: &[]*awscognito.CfnIdentityPool_CognitoIdentityProviderProperty{
			{
				ClientId:     userPoolClient.UserPoolClientId(),
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
			restAPI.ArnForExecuteApi(
				jsii.String("*"),
				jsii.String("/*"),
				jsii.String("*"),
			),
		},
	}))

	apiResource := restAPI.
		Root().
		AddResource(jsii.String("customers"), nil)
	apiResource.AddMethod(
		jsii.String("POST"), awsapigateway.NewLambdaIntegration(function, nil),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		})
	customerSearchResource := apiResource.AddResource(jsii.String("search"), nil)
	customerSearchResource.AddMethod(
		jsii.String("POST"), awsapigateway.NewLambdaIntegration(function, nil),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		})

	customerIDResource := apiResource.AddResource(jsii.String("{customerID}"), nil)
	customerIDResource.AddMethod(
		jsii.String("GET"), awsapigateway.NewLambdaIntegration(function, nil),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		})
	customerIDResource.AddMethod(
		jsii.String("PUT"), awsapigateway.NewLambdaIntegration(function, nil),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		})

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
