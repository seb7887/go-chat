package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/challenge/pkg/config"
	"github.com/challenge/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(userId uint, username string) (*models.Login, error) {
	payload := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"expiry":   time.Now().Add(time.Minute * 120).Unix(), // 2 hours
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), payload)
	token, err := jwtToken.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		return nil, err
	}

	return &models.Login{
		Id:    userId,
		Token: token,
	}, nil
}

func VerifyJwt(header string) (jwt.MapClaims, error) {
	jwtToken := strings.Replace(header, "Bearer ", "", -1)
	claimData := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwtToken, claimData, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claimData, nil
}
