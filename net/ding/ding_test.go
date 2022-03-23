package ding

import "testing"

func TestDing(t *testing.T) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=10e5d13776c67e25752191506a61a71850e183b3a49f1b2c03b3cd7e1b388f59"
	if err := Ding(url, "golib net/ding/ding.go demo"); err != nil {
		t.Fatal(err)
	}
}
