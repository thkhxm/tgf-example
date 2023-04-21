package login

import (
	"github.com/thkhxm/tgf/db"
	"github.com/thkhxm/tgf/example/common/define"
	"github.com/thkhxm/tgf/example/login/entity"
	"github.com/thkhxm/tgf/log"
	"time"
)

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2023/4/20
// ***************************************************
const (
	account_mem_time_out              = 60 * 60 * 12
	account_cache_time_out            = time.Hour * 24 * 3
	account_longevity_update_interval = time.Second * 5
)

type manager struct {
	accountDataManager db.IAutoCacheService[string, *entity.AccountModel]
}

func (this *manager) doLogin(account, password string) (errorCode uint32, userId string) {
	accountModel, _ := this.accountDataManager.Get(account)
	if accountModel == nil || accountModel.Password != password {
		errorCode = define.ErrorCodeAccountNotFound
		log.DebugTag("login", "user login failed , account not found")
		return
	}
	userId = accountModel.UserId
	return
}

func (this *manager) initStruct() {
	//实例化一个完整的
	accountDataBuilder := db.NewAutoCacheBuilder[string, *entity.AccountModel]()
	this.accountDataManager = accountDataBuilder.
		WithMemCache(account_mem_time_out).
		WithAutoCache(RedisKeyAccount, account_cache_time_out).
		WithLongevityCache(account_longevity_update_interval).
		New()
}

func newLoginManager() *manager {
	m := new(manager)
	m.initStruct()
	return m
}
