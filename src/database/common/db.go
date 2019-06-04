package common

import (
	"fmt"
	"strings"
)

type schema string
type table string

type DbTable struct {
	TableName table
	selectSQL string
	insertSQL string
	updateSQL string
	deleteSQL string
}

const (
	memberTableName table = "member"
)

var (
	MemberDt DbTable
)

func init() {
	const (
		selectSQLStr string = `
		SELECT
			%%s
		FROM
			%s
		%%s`

		insertSQLStr string = `
		INSERT INTO
			%s
			(%%s)
		VALUES
			(%%s)`

		updateaSQLStr string = `
		UPDATE 
			%s
		%%s
		%%s`

		deleteSQLStr string = `
		DELETE FROM
			%s
		%%s`
	)

	MemberDt = DbTable{
		TableName: memberTableName,
		selectSQL: fmt.Sprintf(selectSQLStr, memberTableName),
		insertSQL: fmt.Sprintf(insertSQLStr, memberTableName),
		updateSQL: fmt.Sprintf(selectSQLStr, memberTableName),
		deleteSQL: fmt.Sprintf(selectSQLStr, memberTableName),
	}
}

//SelectSQL cols = {id,name}
func (dt DbTable) SelectSQL(cols, conditionCols []string) string {
	colStr := "*"
	if cols != nil && len(cols) > 0 {
		colStr = colSQLStr(cols)
	}

	whereStr := ""
	if conditionCols != nil && len(conditionCols) > 0 {
		whereStr = whereSQLStr(conditionCols)
	}

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, colStr, whereStr)

	return sqlStr
}

//InsertSQL cols = {id,name}
//must have cols
func (dt DbTable) InsertSQL(cols []string) string {
	colStr := colSQLStr(cols)
	valueStr := " :" + strings.Join(cols, " , :")

	sqlStr := dt.insertSQL
	sqlStr = fmt.Sprintf(sqlStr, colStr, valueStr)

	return sqlStr
}

//UpdateSQL cols = {id,name}
//must have condition
func (dt DbTable) UpdateSQL(cols, conditionCols []string) string {
	kv := make([]string, 0)
	for _, col := range cols {
		col = col + "=:" + col
		kv = append(kv, col)
	}
	colStr := "SET " + strings.Join(kv, " , ")

	whereStr := whereSQLStr(conditionCols)

	sqlStr := dt.updateSQL
	sqlStr = fmt.Sprintf(sqlStr, colStr, whereStr)

	return sqlStr
}

//DeleteSQL must have condition
func (dt DbTable) DeleteSQL(conditionCols []string) string {
	whereStr := whereSQLStr(conditionCols)

	sqlStr := dt.selectSQL
	sqlStr = fmt.Sprintf(sqlStr, whereStr)

	return sqlStr
}

func colSQLStr(cols []string) string {
	if cols == nil {
		return ""
	}

	return strings.Join(cols, " , ")
}

func whereSQLStr(conditionCols []string) string {
	kv := make([]string, 0)
	//if conditionCols == nil {
	for _, col := range conditionCols {
		col = col + "=:" + col
		kv = append(kv, col)
	}
	//}

	return "WHERE " + strings.Join(kv, " AND ")
}
