package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/cube-group/golib/log"
)

func main() {
	key, err := totp.Generate(totp.GenerateOpts{Issuer: "err", AccountName: "demo", Digits: 6, Period: 30})
	if err != nil {
		log.StdFatal("err", "Generate", err)
	}

	secret := key.Secret()
	fmt.Println(secret)
	fmt.Println(key.URL())
	for {
		var str string
		fmt.Scanln(&str)
		fmt.Println("validate:", totp.Validate(str, secret))
	}
}
