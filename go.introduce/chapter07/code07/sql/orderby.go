package sql

import (
	"strings"
)

type Order string

const (
	ASC  Order = " ASC "
	DESC Order = " DESC "
)

type OrderByItem struct {
	Field string
	Order
}

func (da *DbData) GenOrderBy() string {
	s := make([]string, 0)
	for _, v := range da.OrderBy {
		s = append(s, DbMapLeft[da.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[da.DbType])
		s = append(s, string(v.Order)) //强制转换为string
		s = append(s, ",")

	}
	s = s[0 : len(s)-1]
	return strings.Join(s, "")
}
