package repository

import (
	"graphql-sample/model"
	"graphql-sample/resolver/types"
	"graphql-sample/store"
)

type UserRepository struct{}

func (UserRepository) FindAll() ([]model.User, error) {
	var userList []model.User
	err := store.DB.Find(&userList).Error // Find의 두번째 파라미터로 id list를 입력하지 않으면 모든 record 조회
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (UserRepository) FindById(id int) (*model.User, error) {
	user := model.User{}
	err := store.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) CreateUser(input types.CreateUserInput) (*model.User, error) {
	user := model.User{}.NewGormUser(input)
	err := store.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
