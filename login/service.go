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
