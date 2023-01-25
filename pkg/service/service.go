package service

import (
	restgo "github.com/go-rest-api"
	"github.com/go-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user restgo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
