package ydorm

import (
	"reflect"
	"strings"
)

//parseEntity 解析实体对象
func parseEntity(entity interface{}) (tInfo *TableInfo) {
	tInfo = &TableInfo{}
	rt := reflect.TypeOf(entity)
	rv := reflect.ValueOf(entity)
	//如果是指针，则需要用Elem()
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	//字段解析
	for i, j := 0, rt.NumField(); i < j; i++ {
		rtf := rt.Field(i) //StructField
		rvf := rv.Field(i) //Value
		var f FieldInfo
		//没有Tag,结构体字段和表字段名一致
		if rtf.Tag == "" {
			f = FieldInfo{Name: rtf.Name, IsPrimaryKey: false, refValue: rvf}
		} else {
			strTag := string(rtf.Tag)
			if strings.Index(strTag, ":") == -1 {
				//Tag中没有":"时，代表表名字段
				tInfo.Name = strTag
				continue
			} else {
				//解析Tag中的field键值为表字段名
				field := rtf.Tag.Get("field")
				//解析Tag中的PK
				isKey := false
				strIsKey := rtf.Tag.Get("iskey")
				if strIsKey == "1" {
					isKey = true
				}
				f = FieldInfo{Name: field, IsPrimaryKey: isKey, refValue: rvf}
			}
		}
		tInfo.Fields = append(tInfo.Fields, f)
	}
	return
}
