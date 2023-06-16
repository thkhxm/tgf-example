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
