package sql

import (
	"strings"
)

//SqlOp Where字段比较类型枚举
type SqlOp string

const (
	GE SqlOp = ">="
	GR SqlOp = ">"
	EQ SqlOp = "="
	LQ SqlOp = "<="
	LR SqlOp = "<"
)

type OrAnd string

const (
	OR   OrAnd = " OR "
	AND  OrAnd = " AND "
	NONE OrAnd = ""
)

type WhereItem struct {
	Field string
	Value string
	SqlOp
	OrAnd
}

func (da *DbData) GenWhere() string {
	s := make([]string, 0)
	for _, v := range da.Where {
		s = append(s, DbMapLeft[da.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[da.DbType])
		s = append(s, string(v.SqlOp)) //强制转换为string
		s = append(s, v.Value)
		//s = append(s, " ")
		s = append(s, string(v.OrAnd)) //强制转换为string
	}
	return strings.Join(s, "")
}
