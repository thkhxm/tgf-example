package service

import (
	"context"
	"github.com/thkhxm/tgf/example/common/pb"
	"github.com/thkhxm/tgf/rpc"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/16
//***************************************************

type ILoginService interface {
	Login(ctx context.Context, args *rpc.Args[*pb.LoginReq], reply *rpc.Reply[*pb.LoginRes]) (err error)
}
