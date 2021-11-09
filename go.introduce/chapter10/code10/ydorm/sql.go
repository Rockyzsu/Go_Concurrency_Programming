package ydorm

import (
	"strings"
)

//genInsertSQL 根据实体生成插入语句
func genInsertSQL(entity interface{}) (string, []interface{}) {
	tbInfo := parseEntity(entity)
	//拼Sql语句，以及参数值
	strSQL := "insert into " + tbInfo.Name
	strFileds := ""
	strValues := ""
	var params []interface{}
	for _, v := range tbInfo.Fields {
		strFileds += v.Name + ","
		strValues += "?,"
		//非导出字段不能Interface()
		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
		}
	}
	strFileds = strings.TrimRight(strFileds, ",")
	strValues = strings.TrimRight(strValues, ",")
	strSQL += " (" + strFileds + ") values(" + strValues + ")"
	return strSQL, params
}

//genDeleteSQL 自动生成删除的sql语句，以主键为删除条件
func genDeleteSQL(entity interface{}) (string, []interface{}) {
	//获取表信息
	tbInfo := parseEntity(entity)
	//拼Sql语句，以及参数值
	strSQL := "delete from " + tbInfo.Name + " where "
	var idVal interface{}
	for _, v := range tbInfo.Fields {
		if v.IsPrimaryKey {
			strSQL += v.Name + "=?"
			if v.refValue.CanInterface() {
				idVal = v.refValue.Interface()
			}
		}
	}
	params := []interface{}{idVal}
	return strSQL, params
}

//genUpdateSQL 自动生成更新的sql语句
func genUpdateSQL(entity interface{}) (string, []interface{}) {
	//获取表信息
	tbInfo := parseEntity(entity)
	//拼Sql语句，以及参数值
	strSQL := "update " + tbInfo.Name + " set "
	strFileds := ""
	strWhere := ""
	var p interface{}
	var params []interface{}
	for _, v := range tbInfo.Fields {
		if v.IsPrimaryKey {
			strWhere += v.Name + "=?"
			if v.refValue.CanInterface() {
				p = v.refValue.Interface()
			}
			continue
		}
		strFileds += v.Name + "=?,"
		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
		}
	}
	params = append(params, p)
	strFileds = strings.TrimRight(strFileds, ",")
	strSQL += strFileds + " where " + strWhere
	return strSQL, params
}
