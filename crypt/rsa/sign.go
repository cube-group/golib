package rsa

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/cube-group/golib/types/convert"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"sort"
)

// 签名
func RsaSignWithMd5(origData string, rsaPrivateKey string) (sign string, err error) {
	//加密
	hashMd5 := md5.Sum([]byte(origData))
	hashed := hashMd5[:]
	privateKeyString, err := base64.StdEncoding.DecodeString(rsaPrivateKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyString)
	if err != nil {
		return
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, hashed)
	return base64.StdEncoding.EncodeToString(signature), err
}

/**RSA验签
 * $data待签名数据(需要先排序，然后拼接)
 * $sign需要验签的签名,需要base64_decode解码
 * 验签用支付公钥
 * return 验签是否通过 bool值
 */
func RsaSignWithMd5Verify(originalData, signData, rsaPublicKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	publicKey, _ := base64.StdEncoding.DecodeString(rsaPublicKey) //RsaLlkPublickey  RsaPublickey
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	hash := md5.New()
	hash.Write([]byte(originalData))
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.MD5, hash.Sum(nil), sign)
	if err != nil {
		return err
	}
	return nil
}

func GetMapSortString(m gin.H) (res string) {
	values := url.Values{}
	for k, v := range m {
		values.Set(k, convert.MustString(v))
	}
	res, _ = url.QueryUnescape(values.Encode())
	return
}

func GetRsaSignWithMd5ReverseValueString(input string) (res, sign string, err error) {
	var inputBytes = []byte(input)
	var any = jsoniter.Get(inputBytes)
	var keys = any.Keys()
	//sort
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	//根据key从m中拿元素，就是按顺序拿了
	for _, k := range keys {
		if k == "sign" {
			sign = any.Get(k).ToString()
		} else {
			res += any.Get(k).ToString()
		}
	}
	return
}