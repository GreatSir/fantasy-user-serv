package handler

import (
	"golang.org/x/net/context"
	"fantasy-user-service/proto"
	"fantasy-user-service/tool"
	"fantasy-user-service/models"
)

type UserServiceHandler struct {
}

func NewHandler()  *UserServiceHandler{
	return &UserServiceHandler{}
}
func (u *UserServiceHandler)WxAuth(ctx context.Context,rq *user.WxAuthRq,rp *user.WxAuthRp) error {
	code := rq.Code
	openid,accessToken, err := tool.GetOpenId(code)
	if err != nil {
		return err
	}
	id , err := tool.GetUserIdByOpenid(openid,accessToken)
	//
	if err != nil {
		return err
	}
	//生成token
	token , expire , err := tool.EncodeToken(id)
	if err != nil {
		return err
	}
	rp.Token = token
	rp.Expire = expire
	return nil
}
func (u *UserServiceHandler)EditInfo(ctx context.Context,rq *user.EditInfoRq,rp *user.EditInfoRp) error {
	user := models.User{Id:rq.Id,UserName:rq.UserName,UserIcon:rq.UserIcon,UserCoin:rq.UserCoin,UserBalance:rq.UserBalance}
	err := tool.EditUserInfo(&user)
	if err != nil {
		return nil
	}
	rp.Rp = rq
	return nil
}
func (u *UserServiceHandler)GetInfo(ctx context.Context,rq *user.GetInfoRq,rp *user.GetInfoRp) error {
	user , err := tool.GetUserInfo(rq.Id)
	if err != nil {
		return err
	}
	rp.Info.Id = user.Id
	rp.Info.UserBalance = user.UserBalance
	rp.Info.UserCoin = user.UserCoin
	rp.Info.UserName = user.UserName
	rp.Info.UserIcon = user.UserIcon
	return nil
}
