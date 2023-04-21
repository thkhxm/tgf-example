package login

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/thkhxm/tgf/example/login/pb"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/util"
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
	version     = "1.0"
	serviceName = "login"
)

// service
// @Description: 登录相关服务
type service struct {
	rpc.Module
	m *manager
}

func (this *service) Login(ctx context.Context, args *[]byte, reply *[]byte) (err error) {
	var (
		//将字节数组转换为pb数据
		req = util.ConvertToPB[*pb.LoginReq](*args)
		res = &pb.LoginRes{}
	)
	defer func() {
		*reply, _ = proto.Marshal(res)
	}()
	res.Error, res.UserId = this.m.doLogin(req.Account, req.Password)
	if res.Error == 0 {

		r, e := rpc.SendRPCMessage(ctx, rpc.Login.New(&rpc.LoginReq{
			UserId:         "res.UserId",
			TemplateUserId: rpc.GetTemplateUserId(ctx),
		}, &rpc.LoginRes{}))
		if e != nil {
			log.DebugTag("login", "user login fail account=%v password=%v code=%v ", req.Account, req.Password, r.ErrorCode)
		}
	}
	log.InfoTag("login", "user login account=%v password=%v userId=%v", req.Account, req.Password, res.UserId)
	return

}

func (this *service) Init() {
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
	s.Name = serviceName
	s.Version = version
	return s
}
