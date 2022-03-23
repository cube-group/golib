//gin redis session middle
//author: linyang
//mail: lin2798003@sina.com
//date: 2019-08
package session

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	SESSION_NAME             = "golib-session-id"
	SESSION_CONTEXT_INSTANCE = "golib-session-context-instance"
)

type Options struct {
	SessionName string //options

	MaxAge int    //默认为86400s
	Secret string //默认为default
	Domain string
	//Path   string
	//Secure   bool
	//HttpOnly bool

	RedisAddress  string
	RedisPassword string
	RedisDB       int
	RedisPoolSize int
}

var sessionOptions *Options
var redisConn *redis.Client

//初始化redis连接
func initRedis() {
	if sessionOptions.SessionName == "" {
		sessionOptions.SessionName = SESSION_NAME
	}
	if sessionOptions.Secret == "" {
		sessionOptions.Secret = "default"
	}
	if sessionOptions.MaxAge == 0 {
		sessionOptions.MaxAge = 86400
	}

	var initOnce sync.Once
	initOnce.Do(func() {
		//log.StdOut("MiddleWare", "Redis init")
		redisConn = redis.NewClient(&redis.Options{
			Addr:     sessionOptions.RedisAddress,
			Password: sessionOptions.RedisPassword,
			DB:       sessionOptions.RedisDB,
			PoolSize: sessionOptions.RedisPoolSize,
		})
		if err := redisConn.Ping().Err(); err != nil {
			log.StdFatal("RedisSession", "Redis Connect Err", err.Error())
		}
	})
}

//生成写入cookie的域
func hostname(c *gin.Context) string {
	if sessionOptions.Domain != "" {
		return sessionOptions.Domain
	}
	return strings.Split(c.Request.Host, ":")[0]
}

func createSid() (string) {
	sid := md5.MD5(fmt.Sprintf(
		"%s-%s-%d-%d",
		sessionOptions.Secret,
		sessionOptions.SessionName,
		time.Now().Nanosecond(),
		rand.Intn(100000),
	))
	return sid
}

type IRedisSession interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) (error)
	Clear() (error)
	GetSid() string
	GetSKey() string
	SetSid(sid string)
}

type RedisSession struct {
	Sid      string //session sid
	SKey     string //session cache key
	Hostname string

	context *gin.Context

	IRedisSession
}

func NewRedisSession(sid string, c *gin.Context) *RedisSession {
	s := &RedisSession{Sid: sid, context: c, Hostname: hostname(c)}
	s.SKey = fmt.Sprintf("%s:%s", sessionOptions.SessionName, sid)
	return s
}

//设置key
func (t *RedisSession) Set(key, value string) error {
	if err := redisConn.HSet(t.SKey, key, value).Err(); err != nil {
		return err
	}
	if err := redisConn.Expire(t.SKey, time.Duration(sessionOptions.MaxAge)*time.Second).Err(); err != nil {
		return err
	}
	t.context.SetCookie(
		sessionOptions.SessionName,
		t.Sid,
		sessionOptions.MaxAge,
		"/",
		t.Hostname,
		false,
		true,
	)
	return nil
}

//获取key
func (t *RedisSession) Get(key string) (string, error) {
	cmd := redisConn.HGet(t.SKey, key)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

//删除key
func (t *RedisSession) Delete(key string) (error) {
	if err := redisConn.HDel(t.SKey, key).Err(); err != nil {
		return err
	}
	return nil
}

//彻底清除session
func (t *RedisSession) Clear() (error) {
	if err := redisConn.Del(t.SKey).Err(); err != nil {
		return err
	}
	t.context.SetCookie(
		sessionOptions.SessionName,
		"",
		0,
		"/",
		t.Hostname,
		false,
		true,
	)
	return nil
}

func (t *RedisSession) GetSid() string {
	return t.Sid
}

func (t *RedisSession) SetSid(sid string) {
	t.Sid = sid
	t.SKey = fmt.Sprintf("%s:%s", sessionOptions.SessionName, sid)
}

func (t *RedisSession) GetSKey() string {
	return t.SKey
}

type RedisSessionNil struct {
	IRedisSession
}

//设置key
func (t *RedisSessionNil) Set(key, value string) error {
	return errors.New("没有初始化Session.GinHandler")
}

//获取key
func (t *RedisSessionNil) Get(key string) (string, error) {
	return "", errors.New("没有初始化Session.GinHandler")
}

//删除key
func (t *RedisSessionNil) Delete(key string) (error) {
	return errors.New("没有初始化Session.GinHandler")
}

//彻底清除session
func (t *RedisSessionNil) Clear() (error) {
	return errors.New("没有初始化Session.GinHandler")
}

func (t *RedisSessionNil) GetSid() string {
	return ""
}

func (t *RedisSessionNil) SetSid(sid string) {
}

func (t *RedisSessionNil) GetSKey() string {
	return ""
}

//通过context获取session实例
func Session(c *gin.Context) (IRedisSession) {
	i, exist := c.Get(SESSION_CONTEXT_INSTANCE)
	if !exist || i == nil {
		return new(RedisSessionNil)
	}
	return i.(*RedisSession)
}

//gin redis session中间件
func GinHandler(options *Options) gin.HandlerFunc {
	sessionOptions = options
	initRedis()

	return func(c *gin.Context) {
		sid, err := c.Cookie(sessionOptions.SessionName)
		if err != nil || sid == "" {
			sid = createSid()
		}
		c.Set(sessionOptions.SessionName, sid)
		c.Set(SESSION_CONTEXT_INSTANCE, NewRedisSession(sid, c))
		c.Next()
	}
}
