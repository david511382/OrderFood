package database

import (
	"database/sql"
	"orderfood/src/database/models"
)

type shopDbSwitch struct {
	redisStatus bool
}

func (d *shopDbSwitch) initRedis() error {
	shops, err := redisMenuDb.GetShop(nil)
	if err != nil {
		d.redisStatus = false
		return nil
	}

	if len(shops) == 0 {
		shops, err = menuDb.GetShop(nil)
		if err != nil {
			d.redisStatus = false
			return err
		}

		for _, shop := range shops {
			err = redisMenuDb.AddShop(shop, nil)
			if err != nil {
				d.redisStatus = false
				return nil
			}
		}
	}

	d.redisStatus = true
	return nil
}

func (d *shopDbSwitch) GetShop(shop *models.Shop) ([]*models.Shop, error) {
	if d.redisStatus {
		result, err := redisMenuDb.GetShop(shop)
		if err == nil {
			return result, nil
		}
	}

	d.redisStatus = false

	result, err := menuDb.GetShop(shop)
	return result, err
}

func (d *shopDbSwitch) AddShop(shop *models.Shop, tx *sql.Tx) error {
	err := menuDb.AddShop(shop, tx)
	if err != nil {
		return err
	}

	if d.redisStatus {
		err = redisMenuDb.AddShop(shop, tx)
		if err != nil {
			d.redisStatus = false
		}
	}
	return nil
}
func (d *shopDbSwitch) DeleteShop(shop *models.Shop, tx *sql.Tx) (int64, error) {
	count, err := menuDb.DeleteShop(shop, tx)
	if err != nil {
		return count, err
	}

	if d.redisStatus {
		redisCount, err := redisMenuDb.DeleteShop(shop, tx)
		if err != nil || count != redisCount {
			d.redisStatus = false
		}
	}
	return count, nil
}
func (d *shopDbSwitch) UpdateShop(shop *models.Shop, tx *sql.Tx) (int64, error) {
	count, err := menuDb.UpdateShop(shop, tx)
	if err != nil {
		return count, err
	}

	if d.redisStatus {
		redisCount, err := redisMenuDb.UpdateShop(shop, tx)
		if err != nil || count != redisCount {
			d.redisStatus = false
		}
	}
	return count, nil
}
