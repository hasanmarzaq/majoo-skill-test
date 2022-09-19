package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	// SaveAvatar(ID uint32, fileLocation string) (User, error)
	GetUserByID(ID uint64) (User, error)
	UpdateUser(Uuid string, input FormUpdateUserInput) (User, error)
	DeleteUser(Uuid string) ([]User, error)
	GetAllUsers() ([]User, error)
	// IsPermission(RoleID uint32, actionName string, methods string, Path string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.UserName = input.UserName
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Uuid = uuid.New().String()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// soal nomor 1 bagian A
func (s *service) Login(input LoginInput) (User, error) {
	user_name := input.UserName
	password := input.Password

	user, err := s.repository.FindByUserName(user_name)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that user name")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID uint64) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(Uuid string, input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByUuid(Uuid)
	fmt.Println(user)
	if err != nil {
		return user, err
	}
	user.Name = input.Name

	user.UserName = input.UserName
	password := input.Password
	if password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}

		user.Password = string(passwordHash)
	}
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	fmt.Println(updatedUser)
	return updatedUser, nil
}

func (s *service) DeleteUser(Uuid string) ([]User, error) {

	user, err := s.repository.DeleteByUuid(Uuid)
	if err != nil {
		return user, err
	}

	return user, nil
}
