package jsonutil

import "encoding/json"

func ToString(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ToBytes(i interface{}) []byte {
	bytes, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return bytes
}

func ToJson(i string) interface{} {
	if i == "" {
		return map[string]interface{}{}
	}
	var res interface{}
	err := json.Unmarshal([]byte(i), &res)
	if err != nil {
		return map[string]interface{}{}
	}
	return res
}
