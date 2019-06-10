package orm

import (
	// "orderfood/src/database/models"
	// "reflect"
	// "strings"

	// proto "github.com/golang/protobuf/proto"
)

// func getID( m interface{})int{
// 	t := reflect.TypeOf(m)
// 	if t == nil || t.Kind() != reflect.Ptr {
// 		return 0
// 	}
// 	t=t.Elem()
	
// 	valT := reflect.ValueOf(m)
// 	if valT.Type().Kind() != reflect.Ptr {
// 		return 0
// 	}
// 	valT =reflect.Indirect(valT)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)

// 		// Get the field tag value
// 		tag := field.Tag.Get(tagName)
// 		tags := strings.Split(tag, tagSpe)
// 		tagField := tags[0]

// 		if tagField == "ID" {
// 			valField:= valT.Field(i)
// 			valField =reflect.Indirect(valField)

// 			if valField.Kind() != reflect.Int32 &&  valField.Kind() != reflect.Int64&&  valField.Kind() != reflect.Int{
// 				return 0
// 			}
// 			return int(valField.Int())
// 		}
// 	}

// 	return 0
// }

// func toPromes(m interface{})(pm proto.Message, err error) {
// 	switch m.(type) {
// 	case *models.Shop:
// 		pm = m.(*models.Shop)
// 	case *models.Item:
// 		pm = m.(*models.Item)
// 	case *models.ItemOption:
// 		pm = m.(*models.ItemOption)
// 	case *models.Option:
// 		pm = m.(*models.Option)
// 	case *models.OptionSelection:
// 		pm = m.(*models.OptionSelection)
// 	case *models.Selection:
// 		pm = m.(*models.Selection)

// 	case *models.Member:
// 		pm = m.(*models.Member)
// 	default:
// 		return nil, undefinedError	
// 	}

// 	return pm, nil
// }

// func setID(id int32, m interface{}) error {
// 	t := reflect.TypeOf(m)
// 	if t == nil || t.Kind() != reflect.Ptr {
// 		return typeError
// 	}
// 	t=t.Elem()
	
// 	valT := reflect.ValueOf(m)
// 	if valT.Type().Kind() != reflect.Ptr {
// 		return typeError
// 	}
// 	valT =reflect.Indirect(valT)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)

// 		// Get the field tag value
// 		tag := field.Tag.Get(tagName)
// 		tags := strings.Split(tag, tagSpe)
// 		tagField := tags[0]

// 		if tagField == "ID" {
// 			valField:= valT.Field(i)
// 			valField =reflect.Indirect(valField)

// 			if valField.Kind() != reflect.Int32 &&  valField.Kind() != reflect.Int64&&  valField.Kind() != reflect.Int{
// 				return typeError
// 			}
// 			if !valField.IsValid() || !valField.CanSet() {
// 				return typeError
// 			}
// 			valField.SetInt(int64(id))
			
// 			break
// 		}
// 	}

// 	return nil
// }

// func readLine(line []byte, m interface{}) (proto.Message, error) {
// 	switch m.(type) {
// 	case *models.Shop:
// 		model := &models.Shop{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.Item:
// 		model := &models.Item{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.ItemOption:
// 		model := &models.ItemOption{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.Option:
// 		model := &models.Option{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.OptionSelection:
// 		model := &models.OptionSelection{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.Selection:
// 		model := &models.Selection{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	case *models.Member:
// 		model := &models.Member{}
// 		err := proto.Unmarshal(line, model)
// 		return model, err
// 	}

// 	return nil, undefinedError
// }
