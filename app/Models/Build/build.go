package Build

import (
	"errors"
	"gebi/app/Http/Middleware"
	"gebi/app/Http/Serializer"
	"gebi/utils/database"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// BuildWhere 构建where条件
//1、and 条件 where := []interface{}{
//    []interface{}{"id", "=", 1},
//    []interface{}{"username", "chen"},
//}
//2、结构体条件  where := user.User{ID: 1, UserName: "chen"}
//3、in,or 条件 where := []interface{}{
//    []interface{}{"id", "in", []int{1, 2}},
//    []interface{}{"username", "=", "chen", "or"},
//}
//4、map条件  where := map[string]interface{}{"id": 1, "username": "chen"}
//5、and or混合条件  where := []interface{}{
//    []interface{}{"id", "in", []int{1, 2}},
//    []interface{}{"username = ? or nickname = ?", "chen", "yond"},
//}
func BuildWhere(db *gorm.DB, where interface{}) *gorm.DB {
	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					Serializer.DBErr("build where step1 error", errors.New("切片长度不能小于2"))
				}
				columnstr := column.(string)
				// 拼接参数形式
				if strings.Index(columnstr, "?") > -1 {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/

					if strings.Index(" in notin ", _opt) > -1 {
						// val 是数组类型
						column = columnstr + " " + opt + " (?)"
					} else if strings.Index(" = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp", _opt) > -1 {
						column = columnstr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			} else if t == reflect.Map /*Map*/ {
				db = db.Where(item)
			} else {
				/*
					// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						db, err = BuildWhere(db, item)
						if err != nil {
							panic(err)
						}
						return db
					})*/

				db = BuildWhere(db, item)
			}
		}
	} else {
		Serializer.DBErr("build where error", errors.New("参数有误"))
	}
	return db
}

//分页list查询
func BuildQueryList(wheres interface{}, columns interface{}, orderBy interface{}) *gorm.DB {
	//var err error
	db := BuildWhere(database.DB, wheres).Select(columns)

	if orderBy != nil && orderBy != "" {
		db = db.Order(orderBy)
	}

	db = db.Limit(Middleware.Limit).Offset((Middleware.Page - 1) * Middleware.Limit)

	return db
}

func BuildUpdates(model interface{}, wheres interface{}, data interface{}) bool {
	if err := BuildWhere(database.DB.Model(model), wheres).Updates(data).Error; err != nil {
		Serializer.DBErr("update address error:", err)
	}

	return true
}
