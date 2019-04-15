package txt

import (
	"bufio"
	"io"
	"orderfood/src/database/models"

	proto "github.com/golang/protobuf/proto"
)

type table string

type dbTable struct {
	Name  table
	Model interface{}
}

func (dt *dbTable) TableName() string {
	return string(dt.Name)
}

func (dt *dbTable) Select() *query {
	q := &query{
		table: dt,
	}
	return q
}

type query struct {
	table       *dbTable
	selectedCol []string
	condition   []string
}

func (q *query) Where() *query {
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

		model, err := readLine(line, q.table.Model, id)
		if err != nil {
			return nil, err
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
