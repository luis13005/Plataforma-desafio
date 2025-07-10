package repositories

import (
	"PLATAFORMA-DESAFIO/domain"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Updating(user *domain.User) (*domain.User, error)
	GetUserForUpdate(user *domain.User, oldPassword string) *domain.User
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

func (repo UserRepositoryDb) Updating(user *domain.User) (*domain.User, error) {
	result := repo.DB.Save(&user)
	if result.Error != nil {
		log.Fatalf("Error updating user! %v", result.Error.Error())
		return user, nil
	}
	return user, nil
}

func (repo UserRepositoryDb) GetUserForUpdate(user *domain.User, oldpassword string) *domain.User {
	fmt.Println("oldUser.ID:", user.ID)
	var oldUser domain.User
	result := repo.DB.First(&oldUser, "id=?", user.ID).Error
	if result != nil {
		log.Fatalf("Old User not found! %v", result.Error())
	}

	user.PrepareUpdate(&oldUser, oldpassword)
	return &oldUser
}
