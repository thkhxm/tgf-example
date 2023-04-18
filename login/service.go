package login

import (
	"context"
	"github.com/thkhxm/tgf/db"
	"github.com/thkhxm/tgf/example/login/entity"
	login_pb "github.com/thkhxm/tgf/example/login/pb"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/util"
	"time"
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

const (
	account_mem_time_out              = 60 * 60 * 12
	account_cache_time_out            = time.Hour * 24 * 3
	account_longevity_update_interval = time.Second * 5
)

// service
// @Description: 登录相关服务
type service struct {
	rpc.Module
	accountDataManager db.IAutoCacheService[string, *entity.AccountModel]
}

func (this *service) Login(ctx context.Context, args *[]byte, reply *[]byte) error {
	var (
		//将字节数组转换为pb数据
		req = util.ConvertToPB[*login_pb.LoginReq](*args)
		//res = &login_pb.LoginRes{}
	)
	accountModel, _ := this.accountDataManager.Get(req.Account)
	if accountModel == nil {

	}
	return nil
}

func (this *service) Init() {
	var ()
	//实例化一个完整的
	accountDataBuilder := db.NewAutoCacheBuilder[string, *entity.AccountModel]()
	this.accountDataManager = accountDataBuilder.WithMemCache(account_mem_time_out).
		WithAutoCache(RedisKeyAccount, account_cache_time_out).
		WithLongevityCache(account_longevity_update_interval).New()
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
