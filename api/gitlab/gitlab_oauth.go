package gitlab

import (
	"context"
	"errors"
	"fmt"
	"github.com/imroc/req"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

// gitlab operate struct
type Gitlab struct {
	accessToken string //gitlab系统token

	//最近一次请求的http状态码
	lastResponseCode int
	Option           Option
}

type Option struct {
	GitlabAddress        string
	GitlabAccessToken    string
	GitlabAppId          string
	GitlabAppSecret      string
	GitlabAppCallbackUrl string
}

type State struct {
	Value string `json:"v"`
	Ref   string `json:"ref"`
}

// gitlab callback 携带参数结构体
type GitlabCallbackVal struct {
	Code  string `form:"code" binding:"required"`
	State string `form:"state" binding:"required"`
}

func NewReq(option Option) *Gitlab {
	return &Gitlab{Option: option}
}

// 取gitlab授权跳转地址
func (t *Gitlab) GetAuthURL(state string) string {
	redirectUrl := fmt.Sprintf(
		"%s/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&state=%s",
		t.Option.GitlabAddress,
		t.Option.GitlabAppId,
		url.QueryEscape(t.Option.GitlabAppCallbackUrl),
		state,
	)
	return redirectUrl
}

// 获取用户信息
// param code: gitlab redirect code
func (t *Gitlab) GetUserInfo(code string) (string, error) {
	c := &oauth2.Config{
		ClientID:     t.Option.GitlabAppId,
		ClientSecret: t.Option.GitlabAppSecret,
		RedirectURL:  t.Option.GitlabAppCallbackUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/oauth/authorize", t.Option.GitlabAddress),
			TokenURL: fmt.Sprintf("%s/oauth/token", t.Option.GitlabAddress),
		},
	}
	token, err := c.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	resp, err := req.Get(fmt.Sprintf("%s/api/v4/user?access_token=%s", t.Option.GitlabAddress, token.AccessToken))
	if err != nil {
		return "", err
	}
	if resp.Response().StatusCode != http.StatusOK {
		return "", errors.New(resp.String())
	}
	return resp.String(), nil
}
