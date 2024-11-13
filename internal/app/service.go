package app

import "github.com/purisaurabh/car-rental/internal/repository"

type service struct {
	Repo repository.Repository
}

type Service interface {
}

func NewService(repo repository.Repository) *service {
	return &service{
		Repo: repo,
	}
}
