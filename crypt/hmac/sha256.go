package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

//hmac sha256
func Sha256(key string, content string) string {
	keyBytes := []byte(key)
	h := hmac.New(sha256.New, keyBytes)
	h.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
