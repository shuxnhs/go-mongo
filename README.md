# go安装与环境变量配置
```shell script
brew install go
# 导入环境变量设置GOPROXY
export GOPATH=/usr/local/go/package
export GOPROXY=https://goproxy.cn
export PATH=$PATH:/usr/local/go/package/bin
```

# 下载包
```shell script
# 下载beego源码和bee工具
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee

# 下载mongodb驱动
go get -u go.mongodb.org/mongo-driver

# mysql
go get -u github.com/go-sql-driver/mysql
```

# 初始化
```shell script
go mod init
go mod download
```

# 启动
```shell script
# 普通启动方式
bee run

# 启动并配置自动生成接口文档
bee run -gendoc=true -downdoc=true

# 生成接口文档
bee generate docs
  
# 访问接口文档
127.0.0.1:8080/swagger

# 打包
bee pack -be GOOS=linux
```


# 项目配置
```shell script
appname = go-mongodb      #  项目名
httpport = 8080           #  端口
runmode = dev             #  运行模式(prod/test/dev)
autorender = false        #  不用模版渲染
copyrequestbody = true    #  是否返回原始请求体数据
EnableDocs = true         #  开启文档内置功能
EnableAdmin = true        #  开启进程内监控
```


# RSA公钥私钥
```
../common           
├── generateRSA.go      生成公钥私钥方法(默认1024位)
├── private.pem         私钥
├── public.pem          公钥
└── rsa                 通用方法:加解密，生成公私钥
    └── rsa.go
```

参考文档：https://blog.csdn.net/qq_36431213/article/details/82982181


# 查询shell命令
```
# 杀掉全部进程：
kill -9 `ps -ef | grep go-mongo | awk '{print $2}'`

```

