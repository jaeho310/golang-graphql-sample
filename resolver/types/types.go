// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type CreateUserInput struct {
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	ID string `json:"id"`
}

type UserList struct {
	List []*User `json:"list"`
}
