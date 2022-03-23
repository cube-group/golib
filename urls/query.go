package urls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/convert"
	"net/url"
)

//追加新url参数
func UrlAddQuery(uri string, params gin.H) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}
	newUri := fmt.Sprintf(
		"%s://%s%s",
		u.Scheme,
		u.Host,
		u.Path,
	)
	newQuery := u.Query()
	for k, v := range params {
		newQuery.Add(k, convert.MustString(v))
	}
	if encodeQuery := newQuery.Encode(); encodeQuery != "" {
		newUri += "?" + encodeQuery
	}
	if u.Fragment != "" {
		newUri += "#" + u.Fragment
	}
	return newUri
}

//从url地址里获取query参数值
func UrlWithoutQuery(uri, key string) string {
	newUri, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	newQuery := newUri.Query()
	newQuery.Del(key)

	newUrl := fmt.Sprintf(
		"%s://%s%s",
		newUri.Scheme,
		newUri.Host,
		newUri.Path,
	)
	if encodeQuery := newQuery.Encode(); encodeQuery != "" {
		newUrl += "?" + encodeQuery
	}
	if fragment := newUri.Fragment; fragment != "" {
		newUrl += "#" + fragment
	}
	return newUrl
}

//从url地址里获取query参数值
func UrlQuery(uri string, key string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Query().Get(key)
}
