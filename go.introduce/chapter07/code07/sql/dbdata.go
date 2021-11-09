package sql

//DbData 构建SQL语句的对象
type DbData struct {
	Table     string
	Fields    string
	Where     []WhereItem
	OrderBy   []OrderByItem
	PageIndex int
	PageSize  int
	DbType
}
