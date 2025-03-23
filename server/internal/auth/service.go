package auth

import (
	"errors"
	"v1/familyManager/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	exsistedUser, _ := service.UserRepository.GetByEmail(email)
	if exsistedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(exsistedUser.HashPass), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return exsistedUser.Email, nil
}

func (service *AuthService) Register(email, password, first_name, last_name string) (string, error) {
	exsistedUser, _ := service.UserRepository.GetByEmail(email)
	if exsistedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	newUser := &user.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		HashPass:  string(hashedPassword),
	}
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", err
	}
	return newUser.Email, nil
}
