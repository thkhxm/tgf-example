package login

import (
	"context"
	"github.com/thkhxm/tgf/example/common/pb"
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

var (
	Version    = "1.0"
	ModuleName = "login"
)

// service
// @Description: 登录相关服务
type service struct {
	rpc.Module
	m *manager
}

func (s *service) Login(ctx context.Context, args *rpc.Args[*pb.LoginReq], reply *rpc.Reply[*pb.LoginRes]) (err error) {
	var (
		//将字节数组转换为pb数据
		req = args.GetData()
		res = &pb.LoginRes{}
	)
	defer func() {
		reply.SetData(res)
	}()
	//执行登录逻辑，如果登录成功返回用户userId
	res.Error, res.UserId = s.m.doLogin(req.Account, req.Password)
	//登录成功，调用网关同步用户userId
	if res.Error == 0 {
		r, e := rpc.SendRPCMessage(ctx, rpc.Login.New(&rpc.LoginReq{
			UserId:         res.UserId,
			TemplateUserId: rpc.GetTemplateUserId(ctx),
		}, &rpc.LoginRes{}))
		if e != nil {
			log.DebugTag("login", "user login fail account=%v password=%v code=%v ", req.Account, req.Password, r.ErrorCode)
			res.Error = uint32(r.ErrorCode)
			return
		}
	}
	//rpc.SendToGate(ctx, api.Login.ModuleName)
	log.InfoTag("login", "user login account=%v password=%v userId=%v", req.Account, req.Password, res.UserId)
	return
}

func (s *service) Init() {
	var ()
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
	log.InfoTag("login", "login service destroy")
}

func newService() rpc.IService {
	s := &service{}
	s.Name = ModuleName
	s.Version = Version
	return s
}
