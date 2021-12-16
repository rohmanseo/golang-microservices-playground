package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rohmanseo/golang-clean-arch/config"
	"github.com/rohmanseo/golang-clean-arch/entity"
	"time"
)

type IJwtToken interface {
	GenerateToken(user entity.User) (entity.Token, error)
	ValidateToken(token entity.Token) bool
	GetValue(token string, key string) (string, error)
}

type JwtToken struct {
	Secret string
}

func NewJwtTokenManager(config config.IConfig) IJwtToken {
	return &JwtToken{Secret: config.Get("JWT_SECRET")}
}
func (tkn *JwtToken) GenerateToken(user entity.User) (entity.Token, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = user.Id
	claim["exp"] = time.Now().Add(time.Hour * 24)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(tkn.Secret))
	tokenType := entity.Token{AccessToken: token}
	if err != nil {
		return tokenType, err
	}
	return tokenType, nil
}
func (tkn *JwtToken) ValidateToken(token entity.Token) bool {
	_, err := jwt.Parse(fmt.Sprintf("%s", token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tkn.Secret), nil
	})
	if err != nil {
		return false
	}
	return true
}

func (tkn *JwtToken) GetValue(token string, key string) (string, error) {
	res, err := jwt.Parse(fmt.Sprintf("%s", token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tkn.Secret), nil
	})

	claims, _ := res.Claims.(jwt.MapClaims)

	if err != nil {
		return "", err
	} else {
		result := fmt.Sprintf("%v", claims[key])
		return result, nil
	}
}
