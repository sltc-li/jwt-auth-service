package repositories

import (
	"errors"

	"github.com/li-go/jwt-auth-service/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// database
var users = []models.User{
	{ID: 1, Name: "taro", Password: "ptaro"},
	{ID: 2, Name: "jiro", Password: "pjiro"},
	{ID: 3, Name: "saburo", Password: "psaburo"},
}

type User struct{}

func (u *User) FindByName(name string) (*models.User, error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}
