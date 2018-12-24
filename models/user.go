package models

import "github.com/astaxie/beego/orm"


//string user_name = 2;
//string wx_openid = 3;
//string user_icon = 4;
//double user_coin = 5;
//double user_balance = 6;
type User struct {
	Id int64 `orm:"PK";json:"id"`
	UserName string `json:"user_name"`
	UserIcon string `json:"user_icon"`
	WxOpenid string `json:"wx_openid"`
	UserCoin float64 `json:"user_coin"`
	UserBalance float64 `json:"user_balance"`
	IsRobot string `json:"is_robot"`
	IsDel int `json:"is_del"`
}

func init()  {
	orm.RegisterModel(new(User))
}
func (u *User)TableName() string {
	return "users"
}
