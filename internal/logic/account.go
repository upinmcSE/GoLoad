package logic

import "context"

type CreateAccountParams struct {
	Username string
	Password string
}

type User struct {
	ID uint64
	Username string
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (User, error)
	CreateSession(ctx context.Context, params CreateAccountParams) (string, error)
}

type account struct {

}

func NewAccount() Account {
	return &account{}
}


func (a *account) CreateAccount(ctx context.Context, params CreateAccountParams) (User, error) {
	return User{}, nil
}