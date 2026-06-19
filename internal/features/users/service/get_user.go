package users_service

import (
	"context"
	"fmt"

	"github.com/i-dontneedaname/study-todoapp/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (domain.User, error) {
	domainUser, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return domainUser, nil
}
