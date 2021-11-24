package Models

import (
	"gebi/utils/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Address struct {
	gorm.Model
	ID      int64  `gorm:"column:id" json:"id"`
	Uid     int64  `gorm:"column:uid" json:"uid"`
	Code    int64  `gorm:"column:code" json:"code"`
	Phone   string `gorm:"column:phone" json:"phone"`
	Address string `gorm:"column:address" json:"address"`
}

func (a *Address) Create(data interface{}) *gorm.DB {
	return database.DB.Clauses(clause.OnConflict{ //如果同一个用户的id和地址相同，更新code和电话,需要数据库配置唯一索引结合使用，否则不会触发
		Columns:   []clause.Column{{Name: "uid"}, {Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "phone"}),
	}).Create(data)
}
