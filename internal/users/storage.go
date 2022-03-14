package users

import (
	"context"
	"users/internal/serviceerror"
)

type Storage interface {
	Create(ctx context.Context, user User) (string, *serviceerror.ErrorResponse)
	FindOne(ctx context.Context, id string) (User, *serviceerror.ErrorResponse)
	FindAll(ctx context.Context) ([]User, *serviceerror.ErrorResponse)
	Update(ctx context.Context, user User) *serviceerror.ErrorResponse
	Delete(ctx context.Context, id string) *serviceerror.ErrorResponse
}
