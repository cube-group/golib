# golang基础库

### Golang入门
* [Effective Go](readme/Effective%20Go.pdf)
* [Go Code Review Comments](readme/Go%20Code%20Review%20Comments.pdf)
* [标准Golang项目Demo](demo/project)

### 如何避免循环引用
例如：services包引用了models包文件，则models包内容不能再import services的内容否则会编译错误

### 1.下载方式
要求：go version>=1.13
```shell
export GOPROXY=https://goproxy.cn
#export GOPRIVATE=github.com
git config --global url."git@github.com:".insteadOf "https://github.com/"
go get github.com/cube-group/golib
```

### 2.使用方法
* [helloworld](demo/helloword/main.go)

### 3.http请求
* [req封装类库Curl快速请求](demo/req_curl/main.go)
* [req封装类库Curl支持日志请求链](demo/req_curl_gin/main.go)
* [req请求](demo/req/main.go)
* [更多req帮助](https://github.com/imroc/req)

### 4.配置文件
golib要求标准配置文件必须使用yaml格式，可轻松使用viper配置文件、mysql、redis和elasticearch
* [配置文件demo](demo/conf/main.go)
* [application.yaml单连接](demo/conf/application.yaml)
* [application.yaml多连接](demo/conf/application-many.yaml)

### 5.web服务
* [ginutil server](demo/ginutil_server/main.go)
* [跨域中间件](demo/ginutil_server_cross/main.go)
* [限流器](demo/ginutil_server_limiter/main.go)

### 6.代码分页
* [代码分页demo](demo/pages/main.go)
* [vue专属分页demo](demo/page/main.go)

### 7.标准输出流日志
* [标准输出流日志Demo](demo/log_std/main.go)

### 8.业务类日志
* [分片日志demo](demo/log_file/main.go)
* [分片日志流程链demo](demo/log_flow/main.go)

### 9.分析类日志
注意：LogName需要提前向王新老师申请
* [分析类日志写demo](demo/log_analysis_write/main.go)
* [分析类日志读demo](demo/log_analysis_read/main.go)

### 10.发送钉钉webhook
* [demo](demo/ding/main.go)

### 11.nsq队列使用
* [生产demo](demo/nsq_produce/main.go)
* [消费demo](demo/nsq_consume/main.go)

### 12.快速并发执行
* [并联获取函数结果demo](demo/async/main.go)

### 13.熔断类库使用
* [熔断使用](demo/req_curl_breaker/main.go)

### 14.api网关识别
[如何校验请求是否来自于api网关](demo/gateway/main.go)
* 测试环境https://api.t.xx.com/${service}
* 生产环境https://api.xx.com/${service}