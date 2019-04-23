package txt

import (
	"bufio"
	"io"
	"orderfood/src/database/models"
	"reflect"
	"strings"

	proto "github.com/golang/protobuf/proto"
)

type table string

type dbTable struct {
	name  table
	model interface{}
}

const (
	tagName = "json"
	tagSpe  = ","
)

func (dt *dbTable) TableName() string {
	return string(dt.name)
}

func (dt *dbTable) Select(cols ...string) *query {
	q := &query{
		table:        dt,
		cols:         make([]string, 0),
		selectedCols: make([]*string, 0),
	}

	t := reflect.TypeOf(dt.model)
	if t == nil || t.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)
		tags := strings.Split(tag, tagSpe)
		tagField := tags[0]

		if tagField != "-" {
			q.cols = append(q.cols, tagField)
		}
	}

	if cols == nil || len(cols) == 0 {
		// select all
		for i := 0; i < len(q.cols); i++ {
			q.selectedCols = append(q.selectedCols, &q.cols[i])
		}

		return q
	}

	for i := 0; i < len(q.cols); i++ {
		for colsI, v := range cols {
			if v == q.cols[i] {
				// Add col
				q.selectedCols = append(q.selectedCols, &q.cols[i])

				// Remove col
				colsLen := len(cols)
				if colsI == colsLen-1 {
					cols = cols[:colsI]
				} else if colsI == 0 {
					cols = cols[1:colsLen]
				} else {
					cols = append(cols[:colsI], cols[colsI:colsLen]...)
				}

				break
			}
		}

		if len(cols) == 0 {
			return q
		}
	}

	if len(cols) != 0 {
		return nil
	}

	return q
}

type query struct {
	table          *dbTable
	cols           []string
	selectedCols   []*string
	checkCondition func(model interface{}) bool
}

type col struct {
	table table
	name  string
}

type rowData struct {
	Data []interface{}
}

type dataTable struct {
	cols []*col
	rows []*rowData
}

func (q *query) Where(condition func(model interface{}) bool) *query {
	q.checkCondition = condition
	return q
}

func (q *query) Exec() ([]interface{}, error) {
	db := txtDb{}
	f, err := db.Connect(q.table.TableName())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	result := make([]interface{}, 0)
	for id := 1; ; id++ {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		model, err := readLine(line, q.table.model, id)
		if err != nil {
			return nil, err
		}

		if q.checkCondition != nil {
			if !q.checkCondition(model) {
				continue
			}
		}

		result = append(result, model)
	}

	return result, nil
}

func readLine(line []byte, m interface{}, id int) (interface{}, error) {
	switch m.(type) {
	case models.Shop:
		model := &models.Shop{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		model.ID = int32(id)
		return model, nil
	case models.ShopItem:
		model := &models.ShopItem{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		//model.ID = int32(id)
		return model, nil
	case models.Item:
		model := &models.Item{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		model.ID = int32(id)
		return model, nil
	case models.ItemSize:
		model := &models.ItemSize{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		//model.ID = int32(id)
		return model, nil
	case models.Size:
		model := &models.Size{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		model.ID = int32(id)
		return model, nil
	case models.ItemKind:
		model := &models.ItemKind{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		//model.ID = int32(id)
		return model, nil
	case models.Kind:
		model := &models.Kind{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		model.ID = int32(id)
		return model, nil
	case models.Member:
		model := &models.Member{}
		err := proto.Unmarshal(line, model)
		if err != nil {
			return nil, err
		}
		model.ID = int32(id)
		return model, nil
	}

	return nil, undefinedError
}