package service

import (
	"log"
	"time"

	"github.com/agusbasari29/simlap.git/entity"
	"github.com/agusbasari29/simlap.git/repository"
	"github.com/agusbasari29/simlap.git/request"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(request request.UserRegisterRequest) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error)
	}
	user.Username = request.Username
	user.Fullname = request.Fullname
	user.Password = string(hashedPassword)
	user.Email = request.Email
	user.Phone = request.Phone
	user.Address = request.Address
	user.Role = request.Role
	user.CreatedAt = time.Now()
	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) UpdateUser(request request.UserRegisterRequest) (entity.Users, error) {
	user := entity.Users{}
	newUser, err := s.userRepository.UpdateUser(user)
	return newUser, err
}
