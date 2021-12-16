package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rohmanseo/golang-clean-arch/entity"
	"testing"
	"time"
)

const secret = "super_secret_tkn"

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAipOiIyMDIxLTEyLTE3VDA4OjQ1OjIzLjUyMDU5MDkrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.OK8kI_SJPZ_35NPOcO-tt93miEmlxY8c0VMcdL1xafo
func TestJwtToken_GenerateTokenGenerateToken(t *testing.T) {
	user := entity.User{
		Id:       1,
		Name:     "random",
		Email:    "email@example.com",
		Password: "password",
	}
	claim := jwt.MapClaims{}
	claim["user_id"] = user.Id
	claim["exp"] = time.Now().Add(time.Hour * 24)
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	fmt.Println(token)
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTE3VDA4OjQ4OjU4LjIxMzE3MzMrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.wcq6qLpDM7buiUywns1Wd0x5vknDP0IoxBhYarhMBeI"
	res, err := jwt.Parse(fmt.Sprintf("%s", token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	fmt.Println(res)
	if err != nil {
		fmt.Println("Token invalid")
	} else {
		fmt.Println("Token valid")
	}
}

type MyClaimsType struct {
	Exp    string `json:"exp"`
	UserId int64  `json:"user_id"`
	jwt.StandardClaims
}

func TestJwtCustomClaim(t *testing.T) {

	claim := MyClaimsType{
		Exp:    time.Now().Add(time.Hour * 24).String(),
		UserId: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	fmt.Println(token)
}

func TestDecodeToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTE3IDA5OjI5OjI1Ljk0NjQ5MjUgKzA3MDAgKzA3IG09Kzg2NDAwLjAwMjk5ODQwMSIsInVzZXJfaWQiOjEsImlzcyI6InRlc3QifQ.EUu17uckfRs_PTJc56h6LtuCBOPh12Zkpw5rIJDY7n8"
	res, err := jwt.Parse(fmt.Sprintf("%s", token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	claims, ok := res.Claims.(jwt.MapClaims)
	s := fmt.Sprintf("%v", claims["user_id"])
	fmt.Println(s + " a")

	if !ok {
		// handle type assertion failure
	}
	// do something with "claims"
	if err != nil {
		fmt.Println("Token invalid")
	} else {
		fmt.Println("Token valid")
	}
}
