package userRepository

import (
	"fmt"
	"go-dynamodb/internal/dto"
	"go-dynamodb/internal/models"
)

func CreateUser(user *dto.CreateUserDTO) (*models.User, error) {
	userFound, err := FindUserByID(user.ID)

	fmt.Println("userFound", userFound)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
