package orm

import (
	"bufio"
	"io"
	"reflect"
	"strings"

	"io/ioutil"
	proto "github.com/golang/protobuf/proto"
)

type Table string

type DbTable struct {
	Name  Table
	Model proto.Message
}

const (
	tagName = "json"
	tagSpe  = ","
)

func (dt *DbTable) TableName() string {
	return string(dt.Name)
}

func (dt *DbTable) Select(cols ...string) *query {
	q := &query{
		table:        dt,
		cols:         make([]string, 0),
		selectedCols: make([]*string, 0),
	}

	t := reflect.TypeOf(dt.Model)
	if t == nil || t.Kind() != reflect.Ptr {
		return nil
	}
	t=t.Elem()

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

func (dt *DbTable) Insert(data interface{}) error {
	f, _, err := Connect(dt.TableName())
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var last []byte
	id:=0
	for {
		line, _, err :=  reader.ReadLine()
		if err == io.EOF {
			break
		}else if err != nil {
			return err
		}
		id=1
		last = line	
	}	
	preData, err := readLine(last, dt.Model)
	if err != nil {
		return err
	}
	id += getID(preData)
	
	newData, err:= setID(int32(id),data)
	out, err := proto.Marshal(newData)
	if err != nil {
		return err
	}

	_, err = f.Write(out)
	if err != nil {
		return err
	}

	_, err = f.WriteString("\n")
	return err
}

func (dt *DbTable) Update(data proto.Message,condiction func(interface{}) bool) error {
	f, filepath, err := Connect(dt.TableName())
	if err != nil {
		return err
	}
	defer f.Close()

	const n = 10 // \n
	allData := make([]byte, 0)	
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}else if err != nil {
			return err
		}

		model, err := readLine(line, dt.Model)
		if err != nil {
			return err
		}

		if condiction(model){
			model=data
		}

		line,err=proto.Marshal(model)
		if err != nil {
			return err
		}

		allData = append(allData, line...)
		allData = append(allData, n)
	}

	err = ioutil.WriteFile(filepath, allData, 0644)
	return err
}

type query struct {
	table          *DbTable
	cols           []string
	selectedCols   []*string
	checkCondition func(model interface{}) bool
}

type col struct {
	table Table
	name  string
}

type rowData struct {
	Data []interface{}
}

type dataTable struct {
	cols []*col
	rows []*rowData
}

func (q *query) Where(condition func(interface{}) bool) *query {
	q.checkCondition = condition
	return q
}

func (q *query) Exec() ([]interface{}, error) {
	f, _, err := Connect(q.table.TableName())
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

		model, err := readLine(line, q.table.Model)
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