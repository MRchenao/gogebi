package Services

import (
	"encoding/json"
	"gebi/app/Http/Serializer"
	"gebi/app/Models"
	"gebi/app/Repositories"
	"gebi/utils/wechat"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
)

var wechatOps = wechat.Wc{}
var wechatRepo = Repositories.WechatRepository{}

type WechatOauthService struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type WechatMeService struct {
}

func (receiver WechatOauthService) WechatUser() string {
	auth := getOauth()
	resToken, err := auth.GetUserAccessToken(receiver.Code)
	if err != nil {
		Serializer.Err(500, "微信授权错误！", err)
	}

	openid := resToken.OpenID
	user := wechatRepo.GetByOpenId(openid)
	if user.ID > 0 {
		return GetToken(user.ID)
	}

	//getUserInfo
	userInfo, err := auth.GetUserInfo(resToken.AccessToken, openid)
	if err != nil {
		Serializer.Err(500, "微信授权获取用户信息错误！", err)
	}

	privilege, _ := json.Marshal(userInfo.Privilege)

	wechatUser := Models.WechatUser{
		Headimgurl: userInfo.HeadImgURL,
		Country:    userInfo.Country,
		City:       userInfo.City,
		Openid:     userInfo.OpenID,
		Nickname:   userInfo.Nickname,
		Sex:        userInfo.Sex,
		Province:   userInfo.Province,
		Privilege:  string(privilege),
		Unionid:    userInfo.Unionid,
	}

	return GetToken(wechatRepo.Create(wechatUser))
}

func getOauth() *oauth.Oauth {
	OfficialAccount := wechatOps.OfficialAccount()
	return OfficialAccount.GetOauth()
}
