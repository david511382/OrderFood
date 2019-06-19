package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"strconv"

	linq "github.com/ahmetb/go-linq"
	proto "github.com/golang/protobuf/proto"
)

func (redis *RedisDb) GetShop(shop *models.Shop) ([]*models.Shop, error) {
	r := redis.R
	shops := make([]*models.Shop, 0)
	id := strconv.Itoa(int(shop.GetID()))
	if id != "0" {
		v := r.HGet(common.ShopDt.Name(), id)
		err := v.Err()
		if err != nil {
			return nil, err
		}

		b, err := v.Bytes()
		if err != nil {
			return nil, err
		}

		nm := &models.Shop{}
		err = proto.Unmarshal(b, nm)
		if err != nil {
			return nil, err
		}

		shops = append(shops, nm)
		return shops, nil
	}

	v := r.HGetAll(common.ShopDt.Name())
	err := v.Err()
	if err != nil {
		return nil, err
	}

	memberMap, err := v.Result()
	if err != nil {
		return nil, err
	}

	for _, memberStr := range memberMap {
		b := []byte(memberStr)
		nm := &models.Shop{}

		err := proto.Unmarshal(b, nm)
		if err != nil {
			return nil, err
		}
		shops = append(shops, nm)
	}

	linq.From(shops).Where(func(m interface{}) bool {
		mem := m.(*models.Shop)

		if shop.GetName() != "" {
			if mem.GetName() != shop.GetName() {
				return false
			}
		}

		return true
	}).ToSlice(&shops)

	return shops, nil
}
func (redis *RedisDb) AddShop(shop *models.Shop) error {
	data, err := proto.Marshal(shop)
	if err != nil {
		return err
	}

	r := redis.R
	id := strconv.Itoa(int(shop.GetID()))
	v := r.HSetNX(common.ShopDt.Name(), id, data)
	if !v.Val() {
		return common.InserFailError
	}
	return nil
}
func (redis *RedisDb) UpdateShop(shop *models.Shop) (int64, error) {
	r := redis.R
	id := strconv.Itoa(int(shop.GetID()))
	v := r.HExists(common.ShopDt.Name(), id)
	err := v.Err()
	if err != nil {
		return 0, err
	}
	if !v.Val() {
		return 0, nil
	}

	data, err := proto.Marshal(shop)
	if err != nil {
		return 0, err
	}

	v = r.HSet(common.ShopDt.Name(), id, data)
	err = v.Err()
	if err != nil {
		return 0, err
	}
	if v.Val() {
		return 0, common.UpdateFailError
	}
	return 1, nil
}
func (redis *RedisDb) DeleteShop(shop *models.Shop) (int64, error) {
	r := redis.R
	id := strconv.Itoa(int(shop.GetID()))
	if id == "0" {
		return 0, common.DbDataError
	}

	v := r.HDel(common.ShopDt.Name(), id)
	err := v.Err()
	if err != nil {
		return 0, err
	}

	count := v.Val()

	return count, nil
}
