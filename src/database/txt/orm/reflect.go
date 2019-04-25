package orm

import (
	"orderfood/src/database/models"
	"reflect"
	"strings"

	proto "github.com/golang/protobuf/proto"
)

func getID( m interface{})int{
	t := reflect.TypeOf(m)
	if t == nil || t.Kind() != reflect.Ptr {
		return 0
	}
	t=t.Elem()
	
	valT := reflect.ValueOf(m)
	if valT.Type().Kind() != reflect.Ptr {
		return 0
	}
	valT =reflect.Indirect(valT)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)
		tags := strings.Split(tag, tagSpe)
		tagField := tags[0]

		if tagField == "ID" {
			valField:= valT.Field(i)
			valField =reflect.Indirect(valField)

			if valField.Kind() != reflect.Int32 &&  valField.Kind() != reflect.Int64&&  valField.Kind() != reflect.Int{
				return 0
			}
			return int(valField.Int())
		}
	}

	return 0
}

func setIDTest(id int32, m interface{}) error {
	t := reflect.TypeOf(m)
	if t == nil || t.Kind() != reflect.Ptr {
		return typeError
	}
	t=t.Elem()
	
	valT := reflect.ValueOf(m)
	if valT.Type().Kind() != reflect.Ptr {
		return typeError
	}
	valT =reflect.Indirect(valT)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)
		tags := strings.Split(tag, tagSpe)
		tagField := tags[0]

		if tagField == "ID" {
			valField:= valT.Field(i)
			valField =reflect.Indirect(valField)

			if valField.Kind() != reflect.Int32 &&  valField.Kind() != reflect.Int64&&  valField.Kind() != reflect.Int{
				return typeError
			}
			if !valField.IsValid() || !valField.CanSet() {
				return typeError
			}
			valField.SetInt(int64(id))
			
			break
		}
	}

	return nil
}

func setID(id int32, m interface{}) (proto.Message, error) {
	switch m.(type) {
	case *models.Shop:
		model := &models.Shop{}
		err:=setIDTest(id,model)
		return model,err
	case *models.ShopItem:
		model := &models.ShopItem{}
		err:=setIDTest(id,model)
		return model,err
	case *models.Item:
		model := &models.Item{}
		err:=setIDTest(id,model)
		return model,err
	case *models.ItemSize:
		model := &models.ItemSize{}
		err:=setIDTest(id,model)
		return model,err
	case *models.Size:
		model := &models.Size{}
		err:=setIDTest(id,model)
		return model,err
	case *models.ItemKind:
		model := &models.ItemKind{}
		err:=setIDTest(id,model)
		return model,err
	case *models.Kind:
		model := &models.Kind{}
		err:=setIDTest(id,model)
		return model,err
	case *models.Member:
		model := &models.Member{}
		err:=setIDTest(id,model)
		return model,err
	}

	return nil, undefinedError
}

func readLine(line []byte, m interface{}) (proto.Message, error) {
	switch m.(type) {
	case *models.Shop:
		model := &models.Shop{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.ShopItem:
		model := &models.ShopItem{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.Item:
		model := &models.Item{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.ItemSize:
		model := &models.ItemSize{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.Size:
		model := &models.Size{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.ItemKind:
		model := &models.ItemKind{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.Kind:
		model := &models.Kind{}
		err := proto.Unmarshal(line, model)
		return model, err
	case *models.Member:
		model := &models.Member{}
		err := proto.Unmarshal(line, model)
		return model, err
	}

	return nil, undefinedError
}
