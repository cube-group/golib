## conf.Init
conf.Init会自动解析application.yaml中的redis和mysql配置
```go
//直接读取远程yaml配置文件
conf.Init(conf.Template{
	AppYamlPath: "https://xx.com/application.yaml",
})

//读取本地yaml配置文件
conf.Init(conf.Template{
	AppYamlPath: "/conf/application.yaml",
})
```
### redis配置
yaml单链接配置
```yaml
redis:
  address: 127.0.0.1:6379
  password: 
  db: 0
  poolSize: 2
```
单连接实例：
```go
conf.Redis().Set("a","1",time.Minute)
```

yaml多链接配置
```yaml
redis:
  default:
    address: 127.0.0.1:6379
    password: 
    db: 0
    poolSize: 2
  other:
    address: 127.0.0.1:6378
    password: 
    db: 0
    poolSize: 2
```
多连接实例：
```go
conf.Redis("default").Set("a","1",time.Minute)
//与下方链接一致
//conf.Redis().Set("a","1",time.Minute)

conf.Redis("other").Set("a","2",time.Minute)
```

### mysql配置
yaml单链接配置
```yaml
mysql:
  address: root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local
  maxIdle: 5
  maxOpen: 100
  logMode: true
```
单连接实例：
```go
conf.DB().Exec("show databases").Error
```

yaml多链接配置
```yaml
mysql:
  default:
    address: root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local
    maxIdle: 5
    maxOpen: 100
    logMode: true
  other:
    address: root:root@tcp(127.0.0.1:3307)/demo?charset=utf8&parseTime=True&loc=Local
    maxIdle: 5
    maxOpen: 100
    logMode: true
```
多连接实例：
```go
conf.DB().Exec("show databases").Error
//与下方链接一致
//conf.DB("default").Exec("show databases").Error

conf.DB("other").Exec("show databases").Error
```