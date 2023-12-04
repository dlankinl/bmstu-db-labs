package services

import (
	"context"
	"fmt"
	"lab_02/internal/domain/others"
)

type OthersService struct {
	repo others.Repository
}

func NewOthersService(repo others.Repository) *OthersService {
	return &OthersService{repo: repo}
}

func (s *OthersService) GetHostAddrPort() (host string, err error) {
	ctx := context.Background()

	host, err = s.repo.GetHostAddrPort(ctx)
	if err != nil {
		return "", fmt.Errorf("сервис для сторонних объектов: %w", err)
	}

	return host, nil
}
