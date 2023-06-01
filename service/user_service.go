package service

import (
	"graphql-sample/model"
	"graphql-sample/repository"
	"graphql-sample/resolver/types"
)

type UserService struct{}

func (UserService) GetUserList() ([]model.User, error) {
	return repository.UserRepository{}.FindAll()
}

func (UserService) GetUser(userId int) (*model.User, error) {
	return repository.UserRepository{}.FindById(userId)

}

func (UserService) CreateUser(input types.CreateUserInput) (*model.User, error) {
	return repository.UserRepository{}.CreateUser(input)

}
