package service

import (
	"crypto/sha1"
	"fmt"

	todo_app "github.com/viktorov13/todo-app"
	"github.com/viktorov13/todo-app/pkg/repository"
)

const salt = "djfsjh2398hfiudsifnaopdjq0"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
