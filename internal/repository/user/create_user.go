package userRepository

import (
	"go-dynamodb/internal/dto"
	"go-dynamodb/internal/models"
)

func CreateUser(user *dto.CreateUserDTO) (*models.User, error) {
	_, err := FindUserByID(user.ID)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
