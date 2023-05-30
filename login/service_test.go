package login_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/thkhxm/tgf/example/common/api"
	"github.com/thkhxm/tgf/example/login/pb"
	"github.com/thkhxm/tgf/robot"
	"github.com/thkhxm/tgf/util"
	"testing"
)

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2023/4/17
// ***************************************************

func TestConvertToPB(t *testing.T) {
	req := &pb.LoginReq{
		Account:  "tim",
		Password: "1234",
	}
	data, _ := proto.Marshal(req)
	c := util.ConvertToPB[*pb.LoginReq](data)
	t.Log(c.Account, ":", c.Password)
}

func Test_service_Login(t *testing.T) {
	robot := robot.NewRobotWs("/tgf")
	robot.Connect("127.0.0.1:8443")
	robot.SendMessage(api.Login.ModuleName, api.Login.Name, &pb.LoginReq{})
}
