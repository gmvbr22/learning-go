package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Subject string `json:"sub,omitempty"`
	Role    string `json:"role,omitempty"`
}

type ClaimsResult struct {
	jwt.StandardClaims
	Role string `json:"role,omitempty"`
}

type Service interface {
	GenerateToken(expiration int, claims *Claims) (string, error)
	ValidateToken(tokenString string) (*ClaimsResult, error)
}

type service struct {
	Secret []byte
}

func NewJWTService(secret []byte) Service {
	return &service{
		Secret: secret,
	}
}

func (service *service) GenerateToken(expiration int, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  claims.Subject,
		"role": claims.Role,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).Unix(),
	})
	tokenString, err := token.SignedString(service.Secret)
	return tokenString, err
}

func (service *service) ValidateToken(tokenString string) (*ClaimsResult, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsResult{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return service.Secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims := token.Claims.(*ClaimsResult)
	m_errors := []string{}
	if len(claims.Subject) == 0 {
		m_errors = append(m_errors, "missing claims.Subject")
	}
	if len(claims.Role) == 0 {
		m_errors = append(m_errors, "missing claims.Role")
	}
	if len(m_errors) > 0 {
		return nil, errors.New(strings.Join(m_errors, "\n"))
	}
	return claims, nil
}
