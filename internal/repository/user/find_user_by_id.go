package userRepository

import (
	"go-dynamodb/internal/database"
	"go-dynamodb/internal/exceptions"
	"go-dynamodb/internal/models"

	"github.com/aws/aws-sdk-go/aws"
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

	params := &dynamodb.GetItemInput{
		TableName: aws.String("Auth.User"),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(id),
			},
		},
	}

	result, err := db.GetItem(params)

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, exceptions.ErrUserNotFound
	}

	if err = dynamodbattribute.UnmarshalMap(result.Item, user); err != nil {
		return nil, err
	}

	return user, nil
}
