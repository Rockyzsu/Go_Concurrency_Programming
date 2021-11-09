package sql

type SqlBuilder struct {
}
type _sqlBuilder struct {
	IBuilder
}

func (*SqlBuilder) GenSelectSQL(da *DbData) string {
	if da == nil {
		return ""
	}
	switch da.DbType {
	case MYSQL:
		{
			mysql := MySQL{}
			sqlBuilder := _sqlBuilder{&mysql}
			sql := sqlBuilder.GenSelect(da)
			return sql
		}
	case MSSQL:
		{
			mssql := MsSQL{}
			sqlBuilder := _sqlBuilder{&mssql}
			sql := sqlBuilder.GenSelect(da)
			return sql
		}
	default:
		{
			return ""
		}
	}
}

func (*SqlBuilder) GenPageSQL(da *DbData) string {
	if da == nil {
		return ""
	}
	switch da.DbType {
	case MYSQL:
		{
			mysql := MySQL{}
			sqlBuilder := _sqlBuilder{&mysql}
			sql := sqlBuilder.GenPage(da)
			return sql
		}
	case MSSQL:
		{
			mssql := MsSQL{}
			sqlBuilder := _sqlBuilder{&mssql}
			sql := sqlBuilder.GenPage(da)
			return sql
		}
	default:
		{
			return ""
		}
	}
}
