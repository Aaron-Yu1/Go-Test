package services

import "github.com/Aaron-Yu1/Go-Test/bookstore_users-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
