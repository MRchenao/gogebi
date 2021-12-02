package Repositories

import (
	"gebi/app/Http/Serializer"
	"gebi/app/Models"
	"gebi/app/Models/Build"
	"gebi/utils/database"
)

type AddressRepository struct {
}

var addressModel = Models.Address{}

func (r AddressRepository) GetList(wheres interface{}) []Models.Address {
	var list []Models.Address

	if err := Build.BuildQueryList(wheres, Build.SetOrderBy("id desc")).Find(&list).Error; err != nil {
		Serializer.DBErr("address list error:", err)
	}

	return list
}

func (r AddressRepository) DelById(id int64) bool {
	address := Models.Address{
		ID: id,
	}

	if err := database.DB.Delete(&address).Error; err != nil {
		Serializer.DBErr("del address error:", err)
	}

	return true
}

func (r AddressRepository) Create(data Models.Address) bool {
	if err := addressModel.Create(&data).Error; err != nil {
		Serializer.DBErr("add address error:", err)
	}

	return true
}

func (r AddressRepository) Updates(wheres interface{}, data interface{}) bool {
	return Build.BuildUpdates(addressModel, wheres, data)
}
