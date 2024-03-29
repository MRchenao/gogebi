package Api

import (
	"encoding/json"
	"gebi/app/Http/Serializer"
	"gebi/app/Models"
	"gebi/app/Repositories"
	"github.com/gin-gonic/gin"
)

var userRepo = Repositories.UserRepository{}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) Models.Users {
	uid, _ := c.Get("user_id")
	if uid == nil {
		Serializer.Err(200, "用户不存在", nil)
	}

	return userRepo.GetById(int(uid.(float64)))
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) Serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Serializer.ParamErr("JSON类型不匹配", err)
	}

	return Serializer.ParamErr("参数错误", err)
}

func SuccessResponse(data interface{}) Serializer.Response {
	return Serializer.Response{
		Data: data,
	}
}
