package common

import (
	"fmt"
	"strings"
)

type schema string
type table string

type DbTable struct {
	DbName    schema
	TableName table
	selectSQL string
	insertSQL string
	updateSQL string
	deleteSQL string
}

const (
	dbName schema = "orderfood_member"

	memberTableName table = "member"
)

var (
	MemberDt DbTable
)

func init() {
	const (
		selectSQLStr string = `
		SELECT
			*
		FROM
			%s
		%%s`

		insertSQLStr string = `
		INSERT INTO
			%s
			(%%s)
		VALUES
			(%%s)`
	)

	MemberDt = DbTable{
		DbName:    dbName,
		TableName: memberTableName,
	}
	MemberDt.selectSQL = fmt.Sprintf(selectSQLStr, MemberDt.name())
	MemberDt.insertSQL = fmt.Sprintf(insertSQLStr, MemberDt.name())
	MemberDt.updateSQL = fmt.Sprintf(selectSQLStr, MemberDt.name())
	MemberDt.deleteSQL = fmt.Sprintf(selectSQLStr, MemberDt.name())
}

func (dt DbTable) name() string {
	return string(dt.DbName) + "." + string(dt.TableName)
}

func (dt DbTable) SelectSQL(conditions []string) string {
	whereStr := ""
	if len(conditions) > 0 {
		whereStr = "WHERE " + strings.Join(conditions, " AND ")
	}

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, whereStr)

	return sqlStr
}

//InsertSQL cols = {id,name}
func (dt DbTable) InsertSQL(cols []string) string {
	colStr := ""
	valueStr := ""
	if len(cols) > 0 {
		colStr = strings.Join(cols, " , ")
		valueStr = strings.Join(cols, " , ")
	}

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, colStr)

	return sqlStr
}

func (dt DbTable) SelectSQL(condition []string) string {
	whereStr := ""
	if len(condition) > 0 {
		whereStr = "WHERE " + strings.Join(condition, " AND ")
	}

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, whereStr)

	return sqlStr
}

func (dt DbTable) SelectSQL(condition []string) string {
	whereStr := ""
	if len(condition) > 0 {
		whereStr = "WHERE " + strings.Join(condition, " AND ")
	}

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, whereStr)

	return sqlStr
}
