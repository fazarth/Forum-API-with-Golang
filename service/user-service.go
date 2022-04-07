package service

import (
	"log"

	"buddyku/model"
	"buddyku/repository"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user model.User) model.User
	Profile(userID string) model.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user model.User) model.User {
	userToUpdate := model.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) model.User {
	return service.userRepository.ProfileUser(userID)
}
