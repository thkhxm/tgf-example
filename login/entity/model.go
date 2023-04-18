package entity

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/17
//***************************************************

// AccountModel
// @Description: 用户表
type AccountModel struct {
	Account  string `orm:"pk"`
	Password string
	UserId   string
}

func (a *AccountModel) GetTableName() string {
	return "game_account"
}
