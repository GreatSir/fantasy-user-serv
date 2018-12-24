package tool

import (
	"github.com/astaxie/beego/orm"
	"fantasy-user-service/models"
)

func GetUserIdByOpenid(openid , accessToken string) (int64, error) {
	o := orm.NewOrm()
	user := models.User{WxOpenid:openid}
	err := o.Read(&user,"WxOpenid")
	if err != nil {
		info ,err := getUserInfo(openid,accessToken)
		if err != nil {
			return 0,err
		}
		user.UserName = info.NickName
		user.UserIcon = info.Headimgurl
		id, err := o.Insert(&user)
		if err != nil {
			return 0 , err
		}
		return id , nil
	}
	return user.Id ,nil
}
func EditUserInfo(u *models.User) error {
	o := orm.NewOrm()
	_,err := o.Update(&u)
	if err != nil {
		return err
	}
	return nil
}
func GetUserInfo(id int64) (models.User,error) {
	user := models.User{Id:id}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		return user,err
	}
	return user ,nil
}
