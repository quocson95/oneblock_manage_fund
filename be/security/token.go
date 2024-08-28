package security

import (
	"errors"
	"oneblock_manage_fund/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type UsertClaims struct {
	Id   int
	Role models.RoleUser
	jwt.Claims
}

var secretKey = "secretKey"

type TokenResponse struct {
	Token      string `json:"token,omitempty"`
	ExpireUnix int64  `json:"expire_unix,omitempty"`
}

func CreateToken(user *models.User) (*TokenResponse, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	if user.ID < 0 || user.Role <= 0 {
		return nil, errors.New("user invalid")
	}
	now := time.Now()
	expireAfter := time.Duration(7 * 24 * time.Hour)
	expireAt := now.Add(expireAfter)
	claims := &UsertClaims{
		Id:   int(user.ID),
		Role: user.Role,
		Claims: jwt.StandardClaims{
			Issuer:    "system",
			NotBefore: now.Unix(),
			// Id:        models.Sha265Random(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		Token:      signedToken,
		ExpireUnix: expireAt.Unix(),
	}, nil
}
