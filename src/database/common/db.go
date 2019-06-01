package common

type schema string
type table string

type DbTable struct {
	DbName    schema
	TableName table
}

const (
	dbName schema = "orderfood_member"

	memberTableName table = "member"
)

var (
	MemberDt DbTable
)

func init() {
	MemberDt = DbTable{
		DbName:    dbName,
		TableName: memberTableName,
	}
}

func (dt DbTable) Name() string {
	return string(dt.DbName) + "." + string(dt.TableName)
}
