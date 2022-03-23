package viperutil

import "github.com/spf13/viper"

func GetString(key ...string) string {
	for _, k := range key {
		if v := viper.GetString(k); v != "" {
			return v
		}
	}
	return ""
}

func GetInt(key ...string) int {
	for _, k := range key {
		if v := viper.GetInt(k); v != 0 {
			return v
		}
	}
	return 0
}
