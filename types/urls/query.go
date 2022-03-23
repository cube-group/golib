package urls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/convert"
	"net/url"
)

//给url添加query字符串
//注意：requestUrl中key如在query字段中则会被覆盖
func AddQuery(requestUrl, query string) string {
	uri, err := url.Parse(requestUrl)
	if err != nil {
		return requestUrl
	}
	values, err := url.ParseQuery(query)
	if err != nil {
		return requestUrl
	}
	newQuery := uri.Query()
	for k, _ := range values {
		newQuery.Set(k, values.Get(k))
	}

	newUrl := fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, uri.Path)
	if newQueryStr := newQuery.Encode(); newQueryStr != "" {
		newUrl += "?" + newQueryStr
	}
	if uri.Fragment != "" {
		newUrl += "#" + uri.Fragment
	}
	return newUrl
}

//给url添加query maps
//注意：requestUrl中key如在query字段中则会被覆盖
func AddQueryMap(requestUrl string, queryMap gin.H) string {
	uri, err := url.Parse(requestUrl)
	if err != nil {
		return requestUrl
	}
	newQuery := uri.Query()
	for k, v := range queryMap {
		newQuery.Set(k, convert.MustString(v))
	}

	newUrl := fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, uri.Path)
	if newQueryStr := newQuery.Encode(); newQueryStr != "" {
		newUrl += "?" + newQueryStr
	}
	if uri.Fragment != "" {
		newUrl += "#" + uri.Fragment
	}
	return newUrl
}
