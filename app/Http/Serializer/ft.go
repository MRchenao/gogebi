package Serializer

import "encoding/json"

type ChageAgePic struct {
	Response *struct {

		// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// RspImgType 为 url 时，返回处理后的图片 url 数据。
		ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// BuildAddress 序列化地址
func BuildFt(ft string) ChageAgePic {
	var agePic ChageAgePic

	if err := json.Unmarshal([]byte(ft), &agePic); err != nil {
		Err(400, "json decode error:", err)
	}

	return agePic
}
