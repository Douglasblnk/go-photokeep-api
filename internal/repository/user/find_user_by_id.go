package userRepository

import (
	"fmt"
	"go-dynamodb/internal/database"
	"go-dynamodb/internal/exceptions"
	"go-dynamodb/internal/models"
	"go-dynamodb/internal/utils"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	ID   int
	Name string
}

func FindUserByID(id string) (*models.User, error) {
	user := &models.User{}
	db := dynamodb.New(database.Session)

	expression := &utils.Builder{
		TableName:                 "Auth.User",
		KeyConditionExpression:    "id",
		ExpressionAttributeValues: id,
	}

	params, err := utils.MountDynamodbQuery(expression)

	if err != nil {
		return nil, err
	}

	fmt.Println("params", params)

	result, err := db.Query(params)

	if err != nil {
		return nil, err
	}

	if result.Items == nil {
		return nil, exceptions.ErrUserNotFound
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, user); err != nil {
		return nil, err
	}

	return user, nil
}
