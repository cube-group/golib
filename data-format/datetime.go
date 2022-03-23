package dataFormat

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"time"
)

func DateTimeFormat(datetime string) (string, error) {
	validate := validator.New()
	if validate.Var(datetime, "number") == nil {
		if len(datetime) < 10 {
			return "", errors.New(fmt.Sprintf("%s 时间解析失败", datetime))
		}
		datetime = datetime[:10]
		at, err := strconv.ParseInt(datetime, 10, 64)
		if err != nil {
			return "", err
		}
		if at > 1000000000 {
			datetime = time.Unix(at, 0).Format("2006-01-02 15:04:05")
		} else {
			return "", errors.New(fmt.Sprintf("%s 时间解析失败", datetime))
		}
	} else if validate.Var(datetime, "datetime=2006-1-2 15:4:5") == nil {
		t, err := time.ParseInLocation("2006-1-2 15:4:5", datetime, time.Local)
		if err != nil {
			return "", err
		}
		datetime = t.Format("2006-01-02 15:04:05")
	} else if err := validate.Var(datetime, "datetime=2006-01-02 15:04:05"); err != nil {
		return "", err
	}
	return datetime, nil
}
