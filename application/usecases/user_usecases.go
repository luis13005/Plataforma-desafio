package usecases

import (
	"PLATAFORMA-DESAFIO/application/repositories"
	"PLATAFORMA-DESAFIO/domain"
	"log"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Insert(user)
	if err != nil {
		log.Fatalf("Error to persist new user: %v", err)
		return user, err
	}

	return user, nil
}
