package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Builder struct {
	TableName                 string
	KeyConditionExpression    string
	ExpressionAttributeValues string
}

func MountDynamodbQuery(expressions *Builder) (*dynamodb.QueryInput, error) {
	keyCond := expression.
		Key(*aws.String(expressions.KeyConditionExpression)).
		Equal(expression.Value(expressions.ExpressionAttributeValues))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	fmt.Println(expr)
	if err != nil {
		return nil, err
	}

	return &dynamodb.QueryInput{
		TableName:                 aws.String(expressions.TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
	}, nil
}
