package service

import (
	"errors"
	"os"
	"time"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		return "", errors.New("no signing key found")
	}

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		return 0, errors.New("no signing key found")
	}

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	idFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId claim is of wrong type")
	}

	return int(idFloat), nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}
