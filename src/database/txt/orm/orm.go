package orm

import (
	"bufio"
	"io"

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

func (dt *DbTable) Select(condiction func(interface{}) bool) ([]interface{},error) {
	f, _, err := Connect(dt.TableName())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	result := make([]interface{}, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}else if err != nil {
			return nil,err
		}

		model, err := readLine(line, dt.Model)
		if err != nil {
			return nil, err
		}

		if condiction==nil||condiction(model){
			result = append(result, model)
		}
	}

	return result, nil
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
	
	err= setID(int32(id),data)
	if err != nil {
		return err
	}
	newData,err:= toPromes(data)
	if err != nil {
		return err
	}
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

		if condiction==nil||condiction(model){
			id:= getID(model)

			err= setID(int32(id),data)
			if err != nil {
				return err
			}
			
			model,err= toPromes(data)
			if err != nil {
				return err
			}
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