package usermanager

import "context"

type UserManager interface {
	CreateUser(ctx context.Context, user interface{}) (id string, err error)
	GetUser(ctx context.Context, id string) (user interface{}, err error)
	UpdateUser(ctx context.Context, id string, user interface{}) error
	DeleteUser(ctx context.Context, id string) error
}
