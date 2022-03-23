package nsq

import (
	"encoding/base64"
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type NsqProducer struct {
	conns []*nsq.Producer
}

func NewNsqProducer(addresses []string) (*NsqProducer, error) {
	i := new(NsqProducer)
	err := i.init(addresses)
	return i, err
}

//close
func (this *NsqProducer) Close() {
	if this.conns != nil {
		for _, c := range this.conns {
			c.Stop()
		}
	}
}

//nsq client
func (this *NsqProducer) init(addresses []string) error {
	conns := make([]*nsq.Producer, 0)
	for _, address := range addresses {
		i, err := nsq.NewProducer(address, nsq.NewConfig())
		i.SetLogger(nil, 0) //屏蔽系统日志
		if err != nil {
			continue
		}
		conns = append(conns, i)
	}
	if len(conns) == 0 {
		return errors.New("nsq producer is nil.")
	}
	this.conns = conns
	return nil
}

//nsq produce
func (this *NsqProducer) produce(topic string, bytes []byte, delay time.Duration) error {
	if bytes == nil {
		return errors.New("nsq produce value is nil.")
	}
	if conn := this.conns[rand.Intn(len(this.conns))]; conn != nil {
		if delay == 0 {
			return conn.Publish(topic, bytes)
		} else {
			return conn.DeferredPublish(topic, delay, bytes)
		}
	}
	return errors.New("nsq producer is nil.")
}

//nsq produce multiple
func (this *NsqProducer) produces(topic string, bytes [][]byte) error {
	if bytes == nil {
		return errors.New("nsq produce value is nil.")
	}
	if conn := this.conns[rand.Intn(len(this.conns))]; conn != nil {
		return conn.MultiPublish(topic, bytes)
	}
	return errors.New("nsq producer is nil.")
}

//nsq produce one
func (t *NsqProducer) One(topic string, i interface{}, delay time.Duration, ascii ...bool) error {
	isAscii := len(ascii) > 0 && ascii[0] == true
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	if isAscii {
		bytes = t.quoteToAscii(bytes)
	}
	return t.produce(topic, bytes, delay)
}

//nsq produce one
func (t *NsqProducer) OneBase64(topic string, i interface{}, delay time.Duration) error {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return t.produce(topic, t.quoteToBase64(bytes), delay)
}

//nsq produce multiple
func (t *NsqProducer) Multiple(topic string, values []interface{}, ascii ...bool) error {
	isAscii := len(ascii) > 0 && ascii[0] == true
	list := make([][]byte, 0)
	for _, i := range values {
		if bytes, err := json.Marshal(i); err == nil {
			if isAscii {
				bytes = t.quoteToAscii(bytes)
			}
			list = append(list, bytes)
		}
	}

	return t.produces(topic, list)
}

//nsq produce multiple
func (t *NsqProducer) MultipleBase64(topic string, values []interface{}) error {
	list := make([][]byte, 0)
	for _, i := range values {
		if bytes, err := json.Marshal(i); err == nil {
			list = append(list, t.quoteToBase64(bytes))
		}
	}

	return t.produces(topic, list)
}

func (t *NsqProducer) quoteToAscii(bytes []byte) []byte {
	sText := string(bytes)
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]

	//\"替换为"
	return []byte(strings.ReplaceAll(textUnquoted, `\"`, `"`))
}

func (t *NsqProducer) quoteToBase64(bytes []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(bytes))
}
