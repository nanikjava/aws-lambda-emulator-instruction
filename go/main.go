package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayCustomAuthorizerRequestTypeRequest
type Response events.APIGatewayCustomAuthorizerResponse

func handler(ctx context.Context, req Request) (Response, error) {
	return generatePolicy("Deny", req.MethodArn), nil
}

func generatePolicy(effect, resource string) Response {
	return Response{
		PrincipalID: "principalID",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Effect:   effect,
					Action:   []string{"execute-api:Invoke"},
					Resource: []string{resource},
				},
			},
		},
	}
}

func main() {
	lambda.Start(handler)
}