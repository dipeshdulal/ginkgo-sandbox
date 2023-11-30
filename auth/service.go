package auth

import "ginkgo-tests/contracts"

type Service struct {
	repo contracts.AuthRepository
}

func NewService(repo contracts.AuthRepository) *Service {
	return &Service{repo}
}

func (s Service) GetItems() ([]string, error) {
	return s.repo.GetItems()
}
