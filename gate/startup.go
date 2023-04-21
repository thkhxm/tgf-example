package gate

import "github.com/thkhxm/tgf/rpc"

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/20
//***************************************************

func Startup() {
	r := rpc.NewRPCServer().
		WithGateway("8898").
		Run()
	<-r
	//Server Destroy Logic
}
