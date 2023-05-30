package main

import (
	"github.com/thkhxm/tgf/example/common/service"
	"github.com/thkhxm/tgf/example/login"
	"github.com/thkhxm/tgf/util"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/5/29
//***************************************************

func main() {
	util.SetAutoGenerateAPICodePath("../common/api")
	util.GeneratorAPI[service.ILoginService](login.ModuleName, login.Version, "api")
}
