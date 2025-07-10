package usecases

import (
	"PLATAFORMA-DESAFIO/application/repositories"
	"PLATAFORMA-DESAFIO/domain"
	"fmt"
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

func (u *UserUseCase) Update(user *domain.User, oldPassword string) (*domain.User, error) {
	fmt.Println("USER UPDATE: ", user.ID, user.Password, user.Email)
	oldUser := u.UserRepository.GetUserForUpdate(user, oldPassword)
	fmt.Println("Old user: ", oldUser.Name, oldUser.Email, oldUser.Password, oldUser.Token)

	user, err := u.UserRepository.Updating(user)
	if err != nil {
		log.Fatalf("Error to update user: %v", err)
		return user, err
	}

	return user, nil
}
