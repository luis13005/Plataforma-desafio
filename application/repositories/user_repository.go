package repositories

import (
	"PLATAFORMA-DESAFIO/domain"
	"log"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	DB *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	err = repo.DB.Create(user).Error
	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}
