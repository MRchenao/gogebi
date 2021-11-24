package Repositories

import (
	"encoding/json"
	"gebi/app/Http/Serializer"
	"gebi/app/Models"
	"gebi/app/Models/Build"
	"gebi/utils/database"
	"gebi/utils/redis_factory"
	"strconv"
)

type UserRepository struct {
}

func (r UserRepository) Create(user Models.Users) int {
	if err := database.DB.Create(&user).Error; err != nil {
		Serializer.DBErr("插入数据失败", err)
	}

	return user.Id
}

func (r UserRepository) Get(wheres interface{}) Models.Users {
	var user Models.Users
	db := Build.BuildWhere(database.DB, wheres)
	if err := db.First(&user).Error; err != nil {
		Serializer.DBErr("用户数据查询错误", err)
	}

	return user
}

func (r UserRepository) GetById(id int) Models.Users {
	var user Models.Users
	cacheKey := redis_factory.RedisKey(strconv.Itoa(id))
	userByte := redis_factory.GetBytes(cacheKey)

	if userByte == nil {
		wheres := Models.Users{Id: id}
		user = r.Get(wheres)
		redis_factory.SetByJson(cacheKey, user)
		return user
	}

	if err := json.Unmarshal(userByte, &user); err != nil {
		Serializer.Err(30001, "用户序列化失败", err)
	}

	return user
}

func (r UserRepository) Count(wheres interface{}) int64 {
	count := int64(0)
	db := Build.BuildWhere(database.DB.Model(Models.Users{}), wheres)
	if err := db.Count(&count).Error; err != nil {
		Serializer.DBErr("用户统计查询错误", err)
	}

	return count
}
