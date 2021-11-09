package sql

//DbType 数据库类型枚举
type DbType int32

const (
	MYSQL DbType = 1
	MSSQL DbType = 2
)

var DbMapLeft map[DbType]string
var DbMapRight map[DbType]string

func init() {
	DbMapLeft = make(map[DbType]string)
	DbMapLeft[MYSQL] = "`"
	DbMapLeft[MSSQL] = "["
	DbMapRight = make(map[DbType]string)
	DbMapRight[MYSQL] = "`"
	DbMapRight[MSSQL] = "]"
}
