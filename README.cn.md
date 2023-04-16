# tgf-example
入门案例，从空项目开始的一系列流程和讲解

## 初始化项目

项目根目录下初始化work 和 mod，最终目录接口参考项目案例

```go
//初始化work空间
go work init
//创建一个common项目，用于存放通用模块
mkdir common
//切换到common目录
cd common
//初始化common的mod, 后面使用自定义的命名即可
go mod init github.com/thkhxm/tgf/example/common
//导入tgf框架
go get -u -v github.com/thkhxm/tgf
//切换到根目录
cd ../
//将common加入到work空间中
go work use common
//重复common的创建方式,分别创建其余的业务项目
```

编辑common目录下的go.mod文件,替换cors版本,在rpcx1.8.0中,使用的是1.8.2版本的cors,所以需要对cors进行替换

```go
replace github.com/rs/cors v1.8.3 => github.com/rs/cors v1.8.2
```



## 业务项目初始化

**以下行为并不是约束，但是建议遵循这种规范**

通过初始化项目的方式，我们创建login项目，并将它与common项目一起加入到了我们的work空间中。

### service

创建login的业务service

service.go

```go
package login

import (
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/15
//***************************************************

var version = "1.0"
var serviceName = "login"

// service
// @Description: 登录相关服务
type service struct {
	rpc.Module
}

func (s *service) GetName() string {
	return s.Name
}

func (s *service) GetVersion() string {
	return s.Version
}

func (s *service) Startup() (bool, error) {
	log.InfoTag("login", "login service startup")
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {
	log.Info("login", "login service destroy")
}

func newService() rpc.IService {
	s := &service{}
	s.Name = serviceName
	s.Version = version
	return s
}
```

- Startup()		该函数会在启动服务的时候,自动执行,我们会在这里进行该业务相关的一些初始化
- Destroy()       该函数会在服务关闭的时候,自动执行.
- GetName()    该函数提供服务的名称,用于服务发现中对该组服务的命名
- GetVersion()  该函数提供服务的版本,用于在进行灰度更新或者停服的时候,对服务的识别
- rpc.Modul      服务通用模块,基础通用的函数和字段



### startup

创建服务启动函数

startup.go

```go
package login

import "github.com/thkhxm/tgf/rpc"

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/15
//***************************************************

func Startup() {
	r := rpc.NewRPCServer().
		WithRandomServicePort(8000, 8010).
		WithService(newService()).
		Run()
	<-r
	//Server Destroy Logic
}

```

- WithRandomServicePort   在一定范围内随机监听端口, 例子中为8000-8010之间,随机一个端口
- WithService    注入单个服务, 可以使用链式调用,注入多个服务在单个进程中. 例如在一个启动项目中,启动login和hall两个服务.只需要在后面加WithService(newHallService())即可
- Run    启动rpc服务





