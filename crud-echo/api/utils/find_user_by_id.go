package utils

import (
	"errors"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
)

func FindUserById(userId string) (*models.UserType, error) {

	for idx, user := range models.Users {
		if user.Id == userId {
			return &models.Users[idx], nil
		}
	}

	return nil, errors.New("User not found")
}