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
