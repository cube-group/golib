## golang标准开发项目

### 项目目录介绍
* app: 程序启动目录
* app/main.go 程序启动入口
* app/cmd: 程序选择目录
* app/controllers: 路由目录
* app/library: 公共类库
* app/models: 模型目录
* app/services: 数据处理目录
* views: html模板目录
* app/public: 静态文件目录

### 相关类库
* github.com/cube-group/golib golang公共库（强烈推荐）
* github.com/getsentry/sentry-go 全局捕捉异常
* github.com/gin-gonic/gin web框架
* github.com/go-redis/redis
* github.com/go-sql-driver/mysql
* github.com/jinzhu/gorm
* github.com/spf13/viper

### Corecd构建和部署
构建脚本
```shell
cd app && CGO_ENABLED=0 GOOS=linux go build -o ../bin/app
cd ..
```
Dockerfile如下
```dockerfile
FROM alpine
ENV APP_PATH /go
#COPY application.yaml $APP_PATH/application.yaml
COPY bin/app $APP_PATH/app
COPY app/views $APP_PATH/views
COPY app/public $APP_PATH/public
WORKDIR $APP_PATH
ENTRYPOINT ["./app"]
CMD ["--cmd","web"]

```