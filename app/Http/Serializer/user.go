package Serializer

import (
	"gebi/app/Models"
)

type Users struct {
	Id        int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Age       int64  `gorm:"column:age" json:"age"`
	Birthday  string `gorm:"column:birthday" json:"birthday"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

func BuildUser(user Models.Users) Users {
	return Users{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Birthday:  user.Birthday.Format("2006-01-02 15:04:05"),
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
