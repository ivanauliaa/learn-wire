package user

import (
	"context"

	"github.com/ivanauliaa/learn-wire/domain"
)

type service struct {
	repo domain.UserRepositoryInterface
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	result, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       result.ID,
		Username: result.Username,
	}

	return &user, nil
}
