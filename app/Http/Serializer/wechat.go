package Serializer

import (
	"gebi/app/Models"
)

type WechatUser struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	City       string `gorm:"column:city" json:"city"`
	Country    string `gorm:"column:country" json:"country"`
	Headimgurl string `gorm:"column:headimgurl" json:"headimgurl"`
	Sex        int32  `gorm:"column:sex" json:"sex"`
}

func BuildWechatUser(user Models.WechatUser) WechatUser {
	return WechatUser{
		Id:         user.ID,
		Name:       user.Nickname,
		City:       user.City,
		Country:    user.Country,
		Headimgurl: user.Headimgurl,
		Sex:        user.Sex,
	}
}
