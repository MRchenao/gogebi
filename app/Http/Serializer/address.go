package Serializer

import (
	"gebi/app/Models"
)

type Address struct {
	ID        int64  `json:"id"`
	Uid       int64  `json:"uid"`
	Code      int64  `json:"code"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAddress 序列化地址
func BuildAddress(address []Models.Address) []Address {
	formatAddr := make([]Address, len(address))
	for index, addr := range address {
		formatAddr[index] = Address{
			ID:        addr.ID,
			Uid:       addr.Uid,
			Code:      addr.Code,
			Phone:     addr.Phone,
			Address:   addr.Address,
			CreatedAt: addr.CreatedAt.Unix(),
		}
	}

	return formatAddr
}
