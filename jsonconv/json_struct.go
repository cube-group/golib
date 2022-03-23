package jsonconv

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"
)

//支持gorm的json类型
//文档：https://github.com/jinzhu/gorm/issues/1935
type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

type Datetime time.Time

func (dt Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(dt).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (dt Datetime) UnmarshalJSON(text []byte) error {
	stamp, err := time.ParseInLocation("2006-01-02 15:04:05", string(text[1:20]), time.Local)
	if err != nil {
		return err
	}
	reflect.ValueOf(dt).Elem().Set(reflect.ValueOf(Datetime(stamp)))
	return err
}

func (dt Datetime) String() string {
	return time.Time(dt).Format("2006-01-02 15:04:05")
}
