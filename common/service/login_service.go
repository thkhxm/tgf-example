package service

import "context"

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/16
//***************************************************

type ILoginService interface {
	Login(ctx context.Context, args *[]byte, reply *[]byte) (err error)
}
