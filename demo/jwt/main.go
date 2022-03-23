package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/cube-group/golib/types/convert"
	"time"
)

const (
	USERNAME = "12312312313123123aaa"
	SECRET   = "123456"
)

func main() {
	//create jwt
	token, err := createJwtToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("token:", token)

	//parse jwt
	username, err := getUsername(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("username:", username)
}

func createJwtToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(86400 * time.Second).Unix(),
		Issuer:    "corecd",
		Audience:  USERNAME,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //生成token，此cookie可直接用于cookie
	accessToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func getUsername(token string) (string, error) {
	// to the callback, providing flexibility.
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRET), nil
	})
	if err != nil {
		return "", err
	}
	if !res.Valid {
		return "", errors.New("jwt token invalid")
	}
	i, ok := res.Claims.(jwt.MapClaims)["aud"]
	if !ok {
		return "", errors.New("jwt token not found userInfo")
	}
	return convert.MustString(i), nil
}
