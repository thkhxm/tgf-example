package login_test

import (
	"github.com/golang/protobuf/proto"
	login_pb "github.com/thkhxm/tgf/example/login/pb"
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
	req := &login_pb.LoginReq{
		Account:  "tim",
		Password: "1234",
	}
	data, _ := proto.Marshal(req)
	c := util.ConvertToPB[*login_pb.LoginReq](data)
	t.Log(c.Account, ":", c.Password)
}
