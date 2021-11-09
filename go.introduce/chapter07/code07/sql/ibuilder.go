package sql

//IBuilder 接口
type IBuilder interface {
	GenSelect(*DbData) string
	GenPage(*DbData) string
}
