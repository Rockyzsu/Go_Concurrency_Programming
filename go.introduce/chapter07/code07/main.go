package main

import (
	"fmt"

	"go.introduce/chapter07/code07/sql"
)

func main() {

	db := sql.DbData{}
	db.Table = "Student"
	db.Fields = "ID,NAME,AGE"
	where := make([]sql.WhereItem, 0)
	where = append(where, sql.WhereItem{
		Field: "XH",
		Value: "'0906'",
		SqlOp: sql.EQ,
		OrAnd: sql.NONE,
	})

	db.Where = where
	orderby := make([]sql.OrderByItem, 0)
	orderby = append(orderby, sql.OrderByItem{
		Field: "XH",
		Order: sql.ASC,
	})
	db.Where = where
	db.OrderBy = orderby
	db.DbType = sql.MYSQL
	sqlBuilder := sql.SqlBuilder{}
	sql0 := sqlBuilder.GenSelectSQL(&db)
	fmt.Println(sql0)
	db.PageIndex = 1
	db.PageSize = 20
	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)
	db.DbType = sql.MSSQL
	sql0 = sqlBuilder.GenSelectSQL(&db)
	fmt.Println(sql0)
	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)

}
