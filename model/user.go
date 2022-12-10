package model

import (
	"graphql-sample/resolver/types"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) GetGqlResponse() *types.User {
	return &types.User{
		ID:   strconv.Itoa(int(user.ID)),
		Name: user.Name,
	}
}

func (User) NewGormUser(input types.CreateUserInput) *User {
	return &User{
		Name: input.Name,
	}
}
