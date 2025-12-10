package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repos repository.DiaryRepository
	secret string
}

func (s AuthService) Register(user model.User) (string, error) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	newUser := model.User{Name: user.Name, Password: hashedPass}
	create_err := s.repos.UserCreate(newUser)

	if create_err != nil {
		return "", create_err
	}

	token, err := s.AuthorizeUser(user)
	return token, err
}

func (s AuthService) AuthorizeUser(user model.User) (string, error) {
	registerdUser, db_err := s.repos.UserGet(user.Name)

	if db_err != nil {
		return "", db_err
	}

	pw_err := bcrypt.CompareHashAndPassword(registerdUser.Password, user.Password)

	if pw_err != nil {
		return "", pw_err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenStr, jwt_err := token.SignedString([]byte(s.secret))

	return tokenStr, jwt_err
}

func CreateAuthService(repos repository.DiaryRepository) AuthService {
	secret := os.Getenv("AUTH_SECRET");
	s := AuthService{repos: repos, secret: secret}
	return s
}