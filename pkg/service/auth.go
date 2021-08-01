package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	solt       = "vkodh73823gd2uq"
	signingKey = "idhdhh7#$knsdnjn"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.GeneratePassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//Получить пользователя из БД
	user, err := s.repo.GetUSer(username, s.GeneratePassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GeneratePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}

func (s *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type *tokenClaims")
	}

	return claims.UserId, nil
}
