package handlers

import (
	"net/http"

	"github.com/sajal1123/serverless-stack/pkg/users"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := users.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return api_response(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return api_response(http.StatusOK, result)
	}

	result, err := users.FetchUsers(tableName, dynaClient)
	if err != nil {
		return api_response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return api_response(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := users.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return api_response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return api_response(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := users.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return api_response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return api_response(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := users.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return api_response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return api_response(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return api_response(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
