// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package events

// APIGatewayProxyRequest contains data coming from the API Gateway proxy
type APIGatewayProxyRequest struct {
	Resource              string                        `json:"resource"` // The resource path defined in API Gateway
	Path                  string                        `json:"path"`     // The url path for the caller
	HTTPMethod            string                        `json:"httpMethod"`
	Headers               map[string]string             `json:"headers"`
	QueryStringParameters map[string]string             `json:"queryStringParameters"`
	PathParameters        map[string]string             `json:"pathParameters"`
	StageVariables        map[string]string             `json:"stageVariables"`
	RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
	Body                  string                        `json:"body"`
	IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyResponse configures the response to be returned by API Gateway for the request
type APIGatewayProxyResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyRequestContext contains the information to identify the AWS account and resources invoking the
// Lambda function. It also includes Cognito identity information for the caller.
type APIGatewayProxyRequestContext struct {
	AccountID    string                    `json:"accountId"`
	ResourceID   string                    `json:"resourceId"`
	Stage        string                    `json:"stage"`
	RequestID    string                    `json:"requestId"`
	Identity     APIGatewayRequestIdentity `json:"identity"`
	ResourcePath string                    `json:"resourcePath"`
	Authorizer   map[string]interface{}    `json:"authorizer"`
	HTTPMethod   string                    `json:"httpMethod"`
	APIID        string                    `json:"apiId"` // The API Gateway rest API Id
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId"`
	AccountID                     string `json:"accountId"`
	CognitoIdentityID             string `json:"cognitoIdentityId"`
	Caller                        string `json:"caller"`
	APIKey                        string `json:"apiKey"`
	SourceIP                      string `json:"sourceIp"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider"`
	UserArn                       string `json:"userArn"`
	UserAgent                     string `json:"userAgent"`
	User                          string `json:"user"`
}

// APIGatewayCustomAuthorizerContext represents the expected format of an API Gateway custom authorizer response.
// Deprecated. Code should be updated to use the Authorizer map from APIGatewayRequestIdentity. Ex: Authorizer["principalId"]
type APIGatewayCustomAuthorizerContext struct {
	PrincipalID *string `json:"principalId"`
	StringKey   *string `json:"stringKey,omitempty"`
	NumKey      *int    `json:"numKey,omitempty"`
	BoolKey     *bool   `json:"boolKey,omitempty"`
}
