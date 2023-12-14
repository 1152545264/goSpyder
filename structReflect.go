package main

import (
	"fmt"
	"reflect"
)

func createQuery(q interface{}) string {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		//获取结构体名称
		t := reflect.TypeOf(q).Name()
		//查询语句
		query := fmt.Sprintf("insert into %s value(", t)
		v := reflect.ValueOf(q)

		//遍历结构体字段
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\\%s\\", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \\%s\\", query, v.Field(i).String())
				}
			case reflect.Struct:
				query = fmt.Sprintf("%s %s", query, createQuery(v.Field(i)))
			}
		}
		query = fmt.Sprintf("%s", query)
		fmt.Println(query)
		return query
	}
	return ""
}
