package dataFormat

import jsoniter "github.com/json-iterator/go"

func StringToMap(s string) map[string]interface{} {
	mm := make(map[string]interface{})
	if s == "" {
		return mm
	} else if s[0] == '{' {
		_ = jsoniter.Unmarshal([]byte(s), &mm)
		return mm
	} else if s[0] == '"' {
		_ = jsoniter.Unmarshal([]byte(s), &s)
		return StringToMap(s)
	} else {
		return mm
	}
}

type MapDataFormat struct {
	any  jsoniter.Any
	data map[string]interface{}
	err  error
}

func NewMapData(data map[string]interface{}) *MapDataFormat {
	return &MapDataFormat{
		any:  jsoniter.Wrap(data),
		data: data,
		err:  nil,
	}
}

func NewMapDataFormatFromJson(data []byte) *MapDataFormat {
	m := make(map[string]interface{})
	err := jsoniter.Unmarshal(data, &m)
	any := jsoniter.Wrap(m)
	return &MapDataFormat{
		any:  any,
		data: m,
		err:  err,
	}
}

func NewMapDataFormatFromJsonString(data string) *MapDataFormat {
	return NewMapDataFormatFromJson([]byte(data))
}

func (t *MapDataFormat) Data() (map[string]interface{}, error) {
	return t.data, t.err
}

func (t *MapDataFormat) ToJson() ([]byte, error) {
	if t.err != nil {
		return nil, t.err
	}
	return jsoniter.Marshal(t.data)
}

func (t *MapDataFormat) ToJsonString() (string, error) {
	b, err := t.ToJson()
	return string(b), err
}

func (t *MapDataFormat) FormatInt(default_ int, key ...string) {
	for _, k := range key {
		t.data[k] = FirstNotZeroInt(t.any.Get(k).ToInt(), default_)
	}
}

func (t *MapDataFormat) FormatInt64(default_ int64, key ...string) {
	for _, k := range key {
		t.data[k] = FirstNotZeroInt64(t.any.Get(k).ToInt64(), default_)
	}
}

func (t *MapDataFormat) FormatFloat64(default_ float64, key ...string) {
	for _, k := range key {
		t.data[k] = FirstNotZeroFloat64(t.any.Get(k).ToFloat64(), default_)
	}
}

func (t *MapDataFormat) FormatString(default_ string, key ...string) {
	for _, k := range key {
		t.data[k] = FirstNotNullString(t.any.Get(k).ToString(), default_)
	}
}

func (t *MapDataFormat) FormatMap(key ...string) {
	for _, k := range key {
		any := t.any.Get(k)
		if any.ValueType() != jsoniter.ObjectValue {
			t.data[k] = StringToMap(any.ToString())
		}
	}
}

func (t *MapDataFormat) UniteMap(key string, other ...string) {
	t.FormatMap(key)
	aims := t.data[key].(map[string]interface{})
	for _, k := range other {
		if k == key {
			continue
		}
		any := t.any.Get(k)
		var m map[string]interface{}
		if any.ValueType() != jsoniter.ObjectValue {
			m = StringToMap(any.ToString())
		} else {
			m = t.data[k].(map[string]interface{})
		}
		for kk, vv := range m {
			aims[kk] = vv
		}
	}
}
