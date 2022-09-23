package domain

import (
	"context"
	"net/http"
)

type (
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}

	UserEntity struct {
		ID       string
		Username string
		Password string
	}

	UserRepositoryInterface interface {
		FetchByUsername(ctx context.Context, username string) (*UserEntity, error)
	}

	UserServiceInterface interface {
		FetchByUsername(ctx context.Context, username string) (*User, error)
	}

	UserHandlerInterface interface {
		FetchByUsername() http.HandlerFunc
	}
)
