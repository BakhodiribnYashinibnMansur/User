package users

import "context"

type Service struct {
	storage Storage
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO for next one
	return
}
