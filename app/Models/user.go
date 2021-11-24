package Models

import (
	"database/sql"
	"time"
)

type Users struct {
	Id        int          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string       `gorm:"column:name" json:"name"`
	Password  string       `gorm:"column:password" json:"password"` // 密码
	Age       int64          `gorm:"column:age" json:"age"`
	Birthday  time.Time    `gorm:"column:birthday" json:"birthday"`
	CreatedAt sql.NullTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at" json:"deleted_at"`
}

