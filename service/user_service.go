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
	CreateUser(request request.UserRegisterRequest) (entity.Users, error)
	GetUser(request request.UserGetRequest) (entity.Users, error)
	GetUsers() ([]entity.Users, error)
	UpdateUser(request request.UserGetRequest) (entity.Users, error)
	DeleteUser(request request.UserGetRequest) (entity.Users, error)
	VerifyCredentials(username string, password string) interface{}
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
	// user.Username = request.Username
	// user.Fullname = request.Fullname
	user.Password = string(hashedPassword)
	// user.Email = request.Email
	// user.Phone = request.Phone
	// user.Address = request.Address
	// user.Role = request.Role
	user.CreatedAt = time.Now()
	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) UpdateUser(request request.UserRegisterRequest) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		log.Fatalf("Failed to smapping. %v", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error)
	}
	user.Password = string(hashedPassword)
	user.UpdatedAt = time.Now()
	newUser, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) GetUser(request request.UserGetRequest) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		log.Fatalf("Failed to smapping.")
	}
	getUser, err := s.userRepository.GetUser(user)
	if err != nil {
		return getUser, err
	}
	return getUser, nil
}

func (s *userService) GetUsers() ([]entity.Users, error) {
	users := []entity.Users{}
	getUsers, err := s.userRepository.GetUsers(users)
	if err != nil {
		return getUsers, err
	}
	return getUsers, nil
}

func (s *userService) DeleteUser(userID uint64) bool {
	user := entity.Users{
		ID: userID,
	}
	return s.userRepository.DeleteUser(user)
}

func (s *userService) VerifyCredentials(username, password string) interface{} {
	result := s.userRepository.GetUserByUsername(username)
	if v, ok := result.(entity.Users); ok {
		comparedPassword := comparePassword(v.Password, password)
		if v.Username == username && comparedPassword {
			return result
		}
		return false
	}
	return false
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
