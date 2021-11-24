package Services

import (
	"errors"
	"gebi/app/Http/Serializer"
	"gebi/app/Models"
	"gebi/app/Repositories"
	"gebi/utils/sign"
	"time"
)

var userRepo = Repositories.UserRepository{}
var jwtService = sign.JwtService{}

type UsersRegisterService struct {
	Name            string    `json:"name" form:"name" binding:"required"`
	Password        string    `json:"password" form:"password" binding:"required"`
	ConfirmPassword string    `json:"confirmpassword" form:"confirmpassword" binding:"required,eqfield=Password"`
	Age             int64     `json:"age" form:"age" binding:"numeric"`
	Birthday        time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02"`
}

type UsersLoginService struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (receiver UsersRegisterService) Register() string {
	if count := userRepo.Count(Models.Users{Name: receiver.Name}); count > 0 {
		Serializer.Err(0, "用户已注册", errors.New(receiver.Name+"已经被注册了"))
	}

	user := Models.Users{
		Name:     receiver.Name,
		Age:      receiver.Age,
		Birthday: receiver.Birthday,
		Password: receiver.Password,
	}
	id := userRepo.Create(user)

	return GetToken(id)
}

func (receiver UsersLoginService) Login() string {
	wheres := Models.Users{
		Name:     receiver.Name,
		Password: receiver.Password,
	}
	user := userRepo.Get(wheres)

	return GetToken(user.Id)
}
