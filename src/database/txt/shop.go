package txt

import (
	"orderfood/src/database/models"

	proto "github.com/golang/protobuf/proto"
)

func (d *txtDb) GetMenus(shopName string) ([]models.MenuItem, error) {
	// get shop id
	ishops, err := shopDT.Select().Where(
		func(model interface{}) bool {
			shop := model.(*models.Shop)
			if shop.Name == shopName {
				return true
			}

			return false
		}).Exec()
	if err != nil {
		return nil, err
	}

	shops := make([]*models.Shop, 0)
	for _, v := range ishops {
		shop := v.(*models.Shop)
		shops = append(shops, shop)
	}

	// get item ids

	// make every item
	// get item size and price
	// get item kind and price
	return nil, nil
}

func (d *txtDb) AddShop(shop *models.Shop) (*models.Shop, error) {
	f, err := d.Connect(shopDT.TableName())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out, err := proto.Marshal(shop)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(out)
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString("\n")
	return shop, nil
}
