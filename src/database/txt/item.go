package txt

import (
	"orderfood/src/database/models"

	proto "github.com/golang/protobuf/proto"
)

func (d *txtDb) AddItem(item *models.Item) (*models.Item, error) {
	f, err := d.Connect(itemDT.TableName())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out, err := proto.Marshal(item)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(out)
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString("\n")
	return item, nil
}

func (d *txtDb) GetItems() ([]*models.Item, error) {
	iitems, err := itemDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	result := make([]*models.Item, 0)
	for _, v := range iitems {
		result = append(result, v.(*models.Item))
	}

	return result, nil
}
