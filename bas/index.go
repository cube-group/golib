//Author: linyang
//Date: 2020-07
//ms-bas基础微服务认证相关方法
package bas

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/bas/consts"
	"github.com/cube-group/golib/bas/token/protos"
	"github.com/cube-group/golib/crypt/base64"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/e"
	"github.com/cube-group/golib/types/convert"
	"github.com/cube-group/golib/types/jsonutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
	"time"
)

const (
	BAS_TOKEN_HTTP_ADDRESS     = "http://ms-bas-token.inner"
	BAS_TOKEN_GRPC_TLS_ADDRESS = "ms-bas-token.inner:11003"
	BAS_TOKEN_GRPC_H2C_ADDRESS = "ms-bas-token.inner:11004"
	BAS_TOKEN_TLS_CERT         = `-----BEGIN CERTIFICATE-----
MIIDOzCCAiOgAwIBAgIQCkUA8jJArsqTRIf6H66kbDANBgkqhkiG9w0BAQsFADBF
MRIwEAYDVQQKEwllb2ZmY25jb20xEjAQBgNVBAsTCWVvZmZjbmNvbTEbMBkGA1UE
AxMSbXMtYmFzLXRva2VuLmlubmVyMB4XDTIwMDcyNzAxNTYxOFoXDTQ3MTIxMzAx
NTYxOFowRTESMBAGA1UEChMJZW9mZmNuY29tMRIwEAYDVQQLEwllb2ZmY25jb20x
GzAZBgNVBAMTEm1zLWJhcy10b2tlbi5pbm5lcjCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAMNNBL6zFX24EwxT8WXBGHI5MxrhxRcGqRWy79SSvYbIX6Ru
+hHffVHVG329DZ1A8oOC4nQd5jFCYwvB+2cSdDBbb6UWlwDQhrbYwdcxLL7Ceq7X
jwj1ph4Qb5uMKcQIkW3nhtNHfcU86MDorZaoopXvcT3Od1WMy82vQLBwcNoRyc95
YW+UX8EgLm8re3UOE+NQyMPL8inEowe77aubOSBYRrtt9lJrlyc/Sfm1SpXIQfTZ
59GpzrCksZNKxq68Dv4HFt+x/4GRVK4v+oR9qlQhlFJ7JcLrL+I9B3Nmq+mtEouS
G8mCzwzxY5d5uhsBcOoyOzZLUZCdhubVMEpeItUCAwEAAaMnMCUwDgYDVR0PAQH/
BAQDAgUgMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA0GCSqGSIb3DQEBCwUAA4IBAQCG
M7umVfstXikAowcnCsV6f5PSnMu52lRoo8X67R7PHibAKcx2kxkJsbwSEiaMTHf/
09m7F+DV4tbtTbxFKGdkw7BSJpca9iPTkSNQejqeX64UDOw/ixnpHw2k6PTQAGS/
d4ACbfcIBdZ4Hed2Mj62s44U/3oztINWAoOhf26PB5QgGe0jHGCiZVfFlad0Vg3n
qt5s0h2XPwqo9z3O25rEPkQPTj8cRIoi2xnTKVL/6g4l/b9YLydqDRlND/FsiByL
GtzXPUdYjsi6Mh42MfaBTFMsc75lHnsQZSVeUWK0627yxIpiQSwPriKgqvtY74C+
xnkMdy9A6LBp8/bbLaNX
-----END CERTIFICATE-----
`
)

type Bas struct {
}

func NewBasService() *Bas {
	i := new(Bas)
	return i
}

//获取ms-bas临时token授权
//GRPC模式
func (t *Bas) GetTokenFromGrpc(gid, ak string, expire int64) (string, int64, error) {
	// Create CertPool
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM([]byte(BAS_TOKEN_TLS_CERT)) {
		return "", 0, errors.New("BAS_TOKEN_TLS_CERT Append Certs Error.")
	}
	creds := credentials.NewClientTLSFromCert(roots, "ms-bas-token.inner")
	conn, err := grpc.Dial(BAS_TOKEN_GRPC_TLS_ADDRESS, grpc.WithTransportCredentials(creds))
	if err != nil {
		return "", 0, err
	}
	defer conn.Close()

	client := protos.NewTokenServiceClient(conn)
	timestamp := convert.MustString(time.Now().Unix())
	md5Str := md5.MD5(fmt.Sprintf("gid=%s&t=%s", gid, timestamp))
	startTime := time.Now().UnixNano()
	resp, err := client.GetToken(context.Background(), &protos.TokenRequest{
		XForwardedApiGid:  gid,
		XForwardedApiAk:   ak,
		XForwardedApiTime: timestamp,
		XForwardedApiMd5:  md5Str,
		Dt:                expire,
	})
	if err != nil {
		if resp != nil {
			return "", 0, errors.New(resp.Error)
		}
		return "", 0, err
	}
	return resp.Value, time.Now().UnixNano() - startTime, nil
}

//获取ms-bas临时token授权
//GRPC NO TLS模式
func (t *Bas) GetTokenFromGrpcWithOutTLS(gid, ak string, expire int64) (string, int64, error) {
	conn, err := grpc.Dial(BAS_TOKEN_GRPC_H2C_ADDRESS, grpc.WithInsecure())
	if err != nil {
		return "", 0, err
	}
	defer conn.Close()

	client := protos.NewTokenServiceClient(conn)
	timestamp := convert.MustString(time.Now().Unix())
	md5Str := md5.MD5(fmt.Sprintf("gid=%s&t=%s", gid, timestamp))
	startTime := time.Now().UnixNano()
	resp, err := client.GetToken(context.Background(), &protos.TokenRequest{
		XForwardedApiGid:  gid,
		XForwardedApiAk:   ak,
		XForwardedApiTime: timestamp,
		XForwardedApiMd5:  md5Str,
		Dt:                expire,
	})
	if err != nil {
		if resp != nil {
			return "", 0, errors.New(resp.Error)
		}
		return "", 0, err
	}
	return resp.Value, time.Now().UnixNano() - startTime, nil
}

//获取ms-bas临时token授权
//HTTP模式
func (t *Bas) GetTokenFromHttp(gid, ak, securityMsName string, expire int64) (string, int64, error) {
	startTime := time.Now().UnixNano()
	requestUrl := fmt.Sprintf("%s?expire=%d&security=%s", BAS_TOKEN_HTTP_ADDRESS, expire, securityMsName)
	resp, err := req.Get(requestUrl, t.CreateXForwardedHeaders(gid, ak))
	if err != nil {
		return "", 0, err
	}
	if resp.Response().StatusCode != http.StatusOK {
		return "", 0, errors.New(resp.String())
	}
	var token string
	if err := e.TryCatch(func() {
		var res map[string]interface{}
		resp.ToJSON(&res)
		if convert.MustInt(res["code"]) == 0 {
			token = res["data"].(string)
		} else {
			panic(res["msg"])
		}
	}); err != nil {
		return "", 0, err
	}
	return token, time.Now().UnixNano() - startTime, nil
}

//获取ms-bas-auth sign签名模式
//ak：签名sign对应的开发者秘钥明文
//sk：签名sign对应的开发者秘钥密文
//ext：签名sign对应的辅助信息{"expire":"高级过期配置","path":"永久有效根据url.Path来签名","security":"安全服务名称"}
//expire：签名sign管理员功能可延长签名时效
func (t *Bas) GetSecuritySignQuery(ak, sk string, ext gin.H) (string, string, error) {
	var sign string
	var query string

	timestamp := time.Now().Unix()
	if len(ext) > 0 {
		extBase64Str := base64.Base64Encode(jsonutil.ToString(ext)) //自动排序升序
		sign = md5.MD5(fmt.Sprintf("ak=%s&t=%d&ext=%s&sk=%s", ak, timestamp, extBase64Str, sk))
		query = fmt.Sprintf("ak=%s&t=%d&ext=%s&sign=%s", ak, timestamp, extBase64Str, sign)
	} else {
		sign = md5.MD5(fmt.Sprintf("ak=%s&t=%d&sk=%s", ak, timestamp, sk))
		query = fmt.Sprintf("ak=%s&t=%d&sign=%s", ak, timestamp, sign)
	}
	return sign, query, nil
}

//获取bas header请求标准
func (t *Bas) CreateXForwardedHeaders(gid, ak string) req.Header {
	timestamp := time.Now().Unix()
	md5Str := md5.MD5(fmt.Sprintf("gid=%s&t=%d", gid, timestamp))
	headers := req.Header{
		consts.AUTH_FORWARDED_GID:  gid,
		consts.AUTH_FORWARDED_TIME: convert.MustString(timestamp),
		consts.AUTH_FORWARDED_AK:   ak,
		consts.AUTH_FORWARDED_MD5:  md5Str,
	}
	return headers
}
