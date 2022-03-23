package server

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil/middlewares"
	"github.com/cube-group/golib/log"
	"html/template"
	"net/http"
	"os"
)

type FuncController func(engine *gin.Engine)

type Config struct {
	AppName            string            //app name
	SentryDsn          string            //sentry地址
	Address            string            //服务地址，如：0.0.0.0:8080
	FuncController     FuncController    //初始化路由函数
	FuncMap            template.FuncMap  //自定义模板全局函数
	StaticRelativePath string            //静态文件请求相对路径
	StaticRoot         string            //静态文件基础目录
	HtmlPattern        string            //html模板路径解析规则
	BeforeMiddleware   []gin.HandlerFunc //中间件之前运行的所有中间件

	UseRedisSession bool //是否启用redis session（会自动查找yaml配置文件中的session.*)
	UseRateLimiter  int  //是否启用限流器
}

//创建配置
func Create(cfg Config) {
	engine := gin.New()
	engine.Use(cfg.BeforeMiddleware...)
	isLocal := os.Getenv("APP_MODE") == ""
	isGinDebug := os.Getenv("GIN_DEBUG") == "1"
	if isLocal || isGinDebug {
		engine.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if !isLocal {
		engine.Use(middlewares.Recovery(cfg.AppName, cfg.SentryDsn))
	}
	//启用限流器
	if cfg.UseRateLimiter > 0 {
		engine.Use(middlewares.RateLimiter(cfg.UseRateLimiter))
	}
	//自定义模板方法函数
	if cfg.FuncMap != nil {
		engine.SetFuncMap(cfg.FuncMap)
	}
	//设置静态文件
	if cfg.StaticRelativePath != "" && cfg.StaticRoot != "" {
		engine.Static(cfg.StaticRelativePath, cfg.StaticRoot)
	}
	//加载视图模板
	if cfg.HtmlPattern != "" {
		engine.LoadHTMLGlob(cfg.HtmlPattern)
	}
	//redis session
	if cfg.UseRedisSession {
		if err := conf.InitRedisSession(engine); err != nil {
			log.StdFatal("Server", "Session", "Fatal:", err)
		}
	}
	//执行路由函数
	if cfg.FuncController != nil {
		cfg.FuncController(engine)
	}

	if cfg.Address == "" {
		cfg.Address = "0.0.0.0:8080"
	}
	log.StdOut("Gin", "Address", cfg.Address)
	s := &http.Server{Addr: cfg.Address, Handler: engine}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.StdFatal("Server", "Fatal:", err)
	}
}
