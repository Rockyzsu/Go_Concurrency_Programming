package sql

import (
	"strconv"
	"strings"
)

//MsSQL 对象
type MsSQL struct {
}

func (*MsSQL) GenSelect(da *DbData) string {
	if da == nil {
		return ""
	}
	s := make([]string, 0)
	s = append(s, "SELECT ")
	fields := strings.Split(da.Fields, ",")
	for _, f := range fields {
		s = append(s, DbMapLeft[da.DbType])
		s = append(s, f)
		s = append(s, DbMapRight[da.DbType])
		s = append(s, ",")
	}
	s = s[0 : len(s)-1]
	s = append(s, " FROM ")
	s = append(s, DbMapLeft[da.DbType])
	s = append(s, da.Table)
	s = append(s, DbMapRight[da.DbType])
	s = append(s, " WHERE ")
	if da.Where == nil {
		s = append(s, "1=1")
	} else {
		s = append(s, da.GenWhere())
	}
	if da.OrderBy != nil {
		s = append(s, " ORDER BY ")
		s = append(s, da.GenOrderBy())
	}

	return strings.Join(s, "")
}

//2012版本+
func (this *MsSQL) GenPage(da *DbData) string {
	if da == nil {
		return ""
	}
	s := make([]string, 0)
	str := this.GenSelect(da)
	//limit (curPage-1)*pageSize,pageSize
	s = append(s, str)
	s = append(s, " offset ")
	i := (da.PageIndex - 1) * da.PageSize
	s = append(s, strconv.Itoa(i))
	s = append(s, " rows fetch next ")
	s = append(s, strconv.Itoa(da.PageSize))
	s = append(s, " rows only ")
	return strings.Join(s, "")
}
