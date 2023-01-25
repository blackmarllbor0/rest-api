package service

import (
	restgo "github.com/go-rest-api"
	"github.com/go-rest-api/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const salt string = "vqouwpiefewfj"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user restgo.User) (int, error) {
	hash, err := generatePassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hash
	return as.repo.CreateUser(user)
}

func generatePassword(password string) (string, error) {
	//
	//hash := sha1.New()
	//hash.Write([]byte(password))
	//return fmt.Sprint(hash.Sum([]byte(salt)))
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
