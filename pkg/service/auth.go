package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isido5ik/todo-app"
	"github.com/isido5ik/todo-app/pkg/repository"
)

// const (
// 	salt       = "fefuif19ojf98h8492hfnf87isbg97"
// 	signingKey = "fh2fiu487fmfoeornjirnijgrnjkrbu487ubi"
// 	tokenTTL   = 12 * time.Hour
// )

// type tokenClaims struct {
// 	jwt.StandardClaims
// 	UserId int `json:"user_id"`
// }

// type AuthService struct {
// 	repo repository.Authorization
// }

// func NewAuthService(repo repository.Authorization) *AuthService {
// 	return &AuthService{repo: repo}
// }

// func (s *AuthService) CreateUser(user todo.User) (int, error) {
// 	user.Password = s.generatePasswordHash(user.Password)
// 	return s.repo.CreateUser(user)
// }

// func (s *AuthService) generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }

// func (s *AuthService) GenerateToken(username, password string) (string, error) {
// 	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
// 	if err != nil {
// 		return "", err
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		&tokenClaims{
// 			jwt.StandardClaims{
// 				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 				IssuedAt:  time.Now().Unix(),
// 			},
// 			user.Id,
// 		})

// 	return token.SignedString([]byte(signingKey)) //signingKey = "fh2fiu487fmfoeornjirnijgrnjkrbu487ubi"
// }

// func (s *AuthService) ParseToken(accessToken string) (int, error) {
// 	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invaling signing method")
// 		}

// 		return []byte(signingKey), nil
// 	})
// 	if err != nil {
// 		return 0, nil
// 	}

// 	claims, ok := token.Claims.(*tokenClaims)
// 	if !ok {
// 		return 0, errors.New("token claims are not of type *tokenClaims")
// 	}
// 	return claims.UserId, nil
// }

// func (s *AuthService) GetUsers() ([]todo.User, error) {
// 	return s.repo.GetUsers()
// }

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
