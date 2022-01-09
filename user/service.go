package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsUsernameAvailable(input CheckUsernameInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.FullName = input.FullName
	user.Email = input.Email
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.IsActive = true
	user.IsLock = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.CreatedBy = 1

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("password is incorrect")
	}

	return user, nil
}

func (s *service) IsUsernameAvailable(input CheckUsernameInput) (bool, error) {
	username := input.Username
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
