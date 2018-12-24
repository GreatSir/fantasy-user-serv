package tool

import (
	"os"
	"net/http"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/micro/go-log"
)

type WechatError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type OpenidResponse struct {
	WechatError
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
}
type UserInfoResponse struct {
	WechatError
	Openid string `json:"openid"`
	NickName string `json:"nickname"`
	Sex string `json:"sex"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`

}
func getUserInfo(openid, accessToken string) (*UserInfoResponse,error) {
	res := new(UserInfoResponse)
	url := "https://api.weixin.qq.com/sns/userinfo?access_token="+accessToken+"&openid="+openid+"&lang=zh_CN"
	response,err:=http.Get(url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Fatalf(err.Error())
		return res ,err
	}
	if res.ErrCode >0 {
		e := errors.New(res.ErrMsg)
		log.Fatalf(res.ErrMsg)
		return res , e
	}
	return res ,nil
}
func GetOpenId(code string) (string , string , error) {
	res := new(OpenidResponse)
	appid := os.Getenv("WECHAT_APPID")
	secret := os.Getenv("WECHAT_SECRET")
	url:="https://api.weixin.qq.com/sns/oauth2/access_token?appid="+appid+"&secret="+secret+"&code="+code+"&grant_type=authorization_code"
	response,err:=http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return "","",err
	}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return "","",err
	}
	if res.ErrCode > 0 {
		e := errors.New(res.ErrMsg)
		return "","",e
	}
	return res.Openid, res.AccessToken,nil
}
