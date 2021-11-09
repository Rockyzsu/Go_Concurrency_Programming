package ydorm

import "reflect"

//FieldInfo 表字段信息
type FieldInfo struct {
	Name         string
	IsPrimaryKey bool
	refValue     reflect.Value
}
