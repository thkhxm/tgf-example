//Auto generated by tgf util
//created at 2023-06-10 16:58:49.7457218 +0800 CST m=+0.029026701

package api

import (
	"github.com/thkhxm/tgf/example/common/pb"

	"github.com/thkhxm/tgf/rpc"
)

var loginService = &rpc.Module{Name: "login", Version: "1.0"}

var (
	Login = rpc.ServiceAPI[*pb.LoginReq, *pb.LoginRes]{
		ModuleName:  loginService.Name,
		Name:        "Login",
		MessageType: loginService.Name + "." + "Login",
	}
)
