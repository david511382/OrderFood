package redis

import (
	"orderfood/src/database/models"
)

func (redis *redisDb) GetShop(*models.Shop) ([]*models.Shop, error) {
	//r := redis.r
	return nil, nil
}
func (redis *redisDb) AddShop(*models.Shop) error {
	return nil
}
func (redis *redisDb) UpdateShop(*models.Shop) (int64, error) {
	return 0, nil
}
func (redis *redisDb) DeleteShop(*models.Shop) (int64, error) {
	return 0, nil
}
