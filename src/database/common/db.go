package common

import (
	"fmt"
	"orderfood/src/database/models"
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
	itemTableName   table = "item"
)

var (
	MemberDt DbTable
	ItemDt   DbTable
)

func init() {
	MemberDt = newDt(memberTableName, &models.Member{})
	ItemDt = newDt(itemTableName, &models.Item{})
}

func newDt(tableName table, m interface{}) DbTable {
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

	return DbTable{
		TableName: tableName,
		selectSQL: fmt.Sprintf(selectSQLStr, tableName),
		insertSQL: fmt.Sprintf(insertSQLStr, tableName),
		updateSQL: fmt.Sprintf(updateaSQLStr, tableName),
		deleteSQL: fmt.Sprintf(deleteSQLStr, tableName),
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
func (dt DbTable) UpdateSQL(cols []string) string {
	kv := make([]string, 0)
	for _, col := range cols {
		col = col + "=:" + col
		kv = append(kv, col)
	}
	colStr := "SET " + strings.Join(kv, " , ")

	whereStr := whereSQLStr([]string{"id"})

	sqlStr := dt.updateSQL
	sqlStr = fmt.Sprintf(sqlStr, colStr, whereStr)

	return sqlStr
}

//DeleteSQL must have condition
func (dt DbTable) DeleteSQL(conditionCols []string) string {
	whereStr := whereSQLStr(conditionCols)

	sqlStr := dt.deleteSQL
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
	for _, col := range conditionCols {
		col = col + "=:" + col
		kv = append(kv, col)
	}

	return "WHERE " + strings.Join(kv, " AND ")
}
