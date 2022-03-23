package user

import (
	"app/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/types/convert"
	"time"
)

func Get(c *gin.Context) (interface{}, error) {
	fmt.Println(ContextUsername(c))
	id := convert.MustUint(c.Param("id"))
	if id == 0 {
		return nil, errors.New("id param is nil.")
	}

	var user models.User
	cacheKey := fmt.Sprintf("user:%d", id)
	if b, err := conf.Redis().Get(cacheKey).Bytes(); err == nil {
		if err := json.Unmarshal(b, &user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	if err := conf.DB().First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}

	//todo save cache
	go func() {
		conf.Redis().Set(cacheKey, user.ToBytes(), time.Minute)
	}()

	//http request
	resp, err := req.Get(viper.GetString("api.service.url"))
	if err != nil {
		fmt.Println(err)
	}
	var res map[string]string
	fmt.Println(resp.ToJSON(&resp), res)

	return &user, nil
}
