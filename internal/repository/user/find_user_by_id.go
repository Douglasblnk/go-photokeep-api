package userRepository

import (
	"go-dynamodb/internal/database"
	"go-dynamodb/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Year  int
	Title string
}

func FindUserByID(id string) (*models.User, error) {
	// user := &models.User{}

	db := dynamodb.New(database.Session)

	db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Auth.User"),
	})

	return nil, nil
}
