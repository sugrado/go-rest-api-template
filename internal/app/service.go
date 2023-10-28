package app

import (
	"github.com/sugrado/go-rest-api-template/internal/app/users"
	"github.com/sugrado/go-rest-api-template/internal/storage"
)

type Service struct {
	user users.Service
}

func RegisterServices(db *storage.Database) *Service {
	return &Service{
		user: users.NewService(db.Users()),
	}
}

func (s *Service) User() users.Service {
	return s.user
}
