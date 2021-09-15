package token

import (
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	service := NewJWTService([]byte("test"))
	tokenString, err := service.GenerateToken(24, &Claims{
		Subject: "Sudo",
		Role:    "Admin",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestValidateToken(t *testing.T) {
	service := NewJWTService([]byte("test"))
	tokenString, err := service.GenerateToken(24, &Claims{
		Subject: "Sudo",
		Role:    "Admin",
	})
	assert.NoError(t, err)
	
	claims, err := service.ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, "Sudo", claims.Subject)
	assert.Equal(t, "Admin", claims.Role)
}

func TestInvalidMethod(t *testing.T) {
	tokenInvalid := []string{
		"eyJhbGciOiAibm9uZSIsInR5cCI6ICJKV1QifQ",
		"eyJleHAiOjE1OTAwMDAwMDB9",
		"eyJleHAiOjE1OTAwMDAwMDB9",
	}

	service := NewJWTService([]byte("test"))
	_, err := service.ValidateToken(strings.Join(tokenInvalid, "."))
	assert.Error(t, err)
}

func TestInvalidClaim(t *testing.T) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, err := token.SignedString([]byte("test"))
	assert.NoError(t, err)

	service := NewJWTService([]byte("test"))
	_, err = service.ValidateToken(tokenString )
	assert.Error(t, err)
}

func TestInvalidSecret(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, err := token.SignedString([]byte("invalid"))
	assert.NoError(t, err)

	service := NewJWTService([]byte("test"))
	_, err = service.ValidateToken(tokenString)
    assert.Error(t, err)
}

func TestExpiration(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * time.Duration(-1)).Unix(),
	})
	tokenString, err := token.SignedString([]byte("test"))
	assert.NoError(t, err)
	
	service := NewJWTService([]byte("test"))
	_, err = service.ValidateToken(tokenString)
    assert.Error(t, err)
}
