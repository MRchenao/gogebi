package Services

import (
	"gebi/app/Models"
	"gebi/app/Repositories"
)

type AddressListService struct {
	ID      int64  `json:"id" form:"id"`
	Code    int64  `json:"code" form:"code"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

type AddressUpdateService struct {
	ID      int64  `json:"id" form:"id" binding:"required"`
	Code    int64  `json:"code" form:"code"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

type AddressAddService struct {
	Uid     int64  `json:"uid" form:"uid" binding:"required"`
	Code    int64  `json:"code" form:"code"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

type AddressDelService struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

var addrRepo = Repositories.AddressRepository{}

func (receiver AddressListService) List() []Models.Address {
	return addrRepo.GetList(receiver)
}

func (receiver AddressUpdateService) Update() bool {
	where := Models.Address{ID: receiver.ID}
	return addrRepo.Updates(where, receiver)
}

func (receiver AddressAddService) Add() bool {
	address := Models.Address{
		Uid:     receiver.Uid,
		Code:    receiver.Code,
		Phone:   receiver.Phone,
		Address: receiver.Address,
	}

	return addrRepo.Create(address)
}

func (receiver AddressDelService) Del() bool {
	return addrRepo.DelById(receiver.ID)
}
