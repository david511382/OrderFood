package redis

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"orderfood/src/database/redis/member"
	"orderfood/src/database/redis/menu"
	"orderfood/tags"
	"strconv"
	"testing"

	proto "github.com/golang/protobuf/proto"

	"github.com/go-redis/redis"
)

const (
	i1 int32 = 1
	i2 int32 = 2
	i3 int32 = 3
	i4 int32 = 4
	i5 int32 = 5

	s1 string = "1"
	s2 string = "2"
	s3 string = "3"
	s4 string = "4"
	s5 string = "5"

	b1 bool = false
	b2 bool = true
)

var (
	memberDb common.IRedisMember
	menuDb   common.IRedisMenu

	memberDbMembers = []models.Member{
		models.Member{
			ID:       i1,
			Name:     s1,
			Username: s1,
			Password: s1,
		},
		models.Member{
			ID:       i2,
			Name:     s2,
			Username: s2,
			Password: s2,
		},
		models.Member{
			ID:       i3,
			Name:     s3,
			Username: s3,
			Password: s3,
		},
		models.Member{
			ID:       i4,
			Name:     s4,
			Username: s4,
			Password: s4,
		},
		models.Member{
			ID:       i5,
			Name:     s5,
			Username: s5,
			Password: s5,
		},
	}

	menuDbShops = []models.Shop{
		models.Shop{
			ID:   i1,
			Name: s1,
		},
		models.Shop{
			ID:   i2,
			Name: s2,
		},
		models.Shop{
			ID:   i3,
			Name: s3,
		},
	}
	// menuDbItems = []models.Item{
	// 	models.Item{
	// 		ID:      i1,
	// 		Name:    s1,
	// 		Shop_ID: i1,
	// 		Price:   i1,
	// 	},
	// 	models.Item{
	// 		ID:      i2,
	// 		Name:    s2,
	// 		Shop_ID: i1,
	// 		Price:   i2,
	// 	},
	// 	models.Item{
	// 		ID:      i3,
	// 		Name:    s3,
	// 		Shop_ID: i2,
	// 		Price:   i3,
	// 	},
	// 	models.Item{
	// 		ID:      i4,
	// 		Name:    s4,
	// 		Shop_ID: i1,
	// 		Price:   i4,
	// 	},
	// 	models.Item{
	// 		ID:      i5,
	// 		Name:    s5,
	// 		Shop_ID: i1,
	// 		Price:   i3,
	// 	},
	// }
	// menuDbItemOption = []models.ItemOption{
	// 	models.ItemOption{
	// 		ID:        i1,
	// 		Item_ID:   i1,
	// 		Option_ID: i1,
	// 	},
	// 	models.ItemOption{
	// 		ID:        i2,
	// 		Item_ID:   i2,
	// 		Option_ID: i2,
	// 	},
	// 	models.ItemOption{
	// 		ID:        i3,
	// 		Item_ID:   i3,
	// 		Option_ID: i3,
	// 	},
	// 	models.ItemOption{
	// 		ID:        i4,
	// 		Item_ID:   i4,
	// 		Option_ID: i1,
	// 	},
	// }
	// menuDbOptions = []models.Option{
	// 	models.Option{
	// 		ID:         i1,
	// 		Select_Num: i1,
	// 	},
	// 	models.Option{
	// 		ID:         i2,
	// 		Select_Num: i2,
	// 	},
	// 	models.Option{
	// 		ID:         i3,
	// 		Select_Num: i3,
	// 	},
	// }
	// menuDbSelections = []models.Selection{
	// 	models.Selection{
	// 		ID:        i1,
	// 		Name:      s1,
	// 		Option_ID: i1,
	// 		Price:     i1,
	// 	},
	// 	models.Selection{
	// 		ID:        i2,
	// 		Name:      s2,
	// 		Option_ID: i2,
	// 		Price:     i2,
	// 	},
	// 	models.Selection{
	// 		ID:        i3,
	// 		Name:      s3,
	// 		Option_ID: i3,
	// 		Price:     i3,
	// 	},
	// 	models.Selection{
	// 		ID:        i4,
	// 		Name:      s4,
	// 		Option_ID: i1,
	// 		Price:     i4,
	// 	},
	// 	models.Selection{
	// 		ID:        i5,
	// 		Name:      s5,
	// 		Option_ID: i2,
	// 		Price:     i5,
	// 	},
	// }
)

func TestMain(m *testing.M) {
	cfg, _ := tags.InitConfig("../../config/test-config.yml")

	rm, err := connect(cfg.RedisMember)
	if err != nil {
		panic(err)
	}
	defer rm.Close()

	rdsM := &member.RedisDb{rm}
	err = initMemberDb(rdsM.R)
	if err != nil {
		panic(err)
	}
	memberDb = rdsM

	rmn, err := connect(cfg.RedisMenu)
	if err != nil {
		panic(err)
	}
	defer rmn.Close()

	rdsMn := &menu.RedisDb{rmn}
	err = initMenuDb(rdsMn.R)
	if err != nil {
		panic(err)
	}
	menuDb = rdsMn

	m.Run()

	rm.FlushAll()
}

func initMemberDb(r *redis.Client) error {
	v := r.FlushDb()
	err := v.Err()
	if err != nil {
		return err
	}

	for _, member := range memberDbMembers {
		data, err := proto.Marshal(&member)
		if err != nil {
			return err
		}

		id := strconv.Itoa(int(member.GetID()))
		v := r.HSetNX(common.MemberDt.Name(), id, data)
		err = v.Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func initMenuDb(r *redis.Client) error {
	v := r.FlushDb()
	err := v.Err()
	if err != nil {
		return err
	}

	for _, shop := range menuDbShops {
		data, err := proto.Marshal(&shop)
		if err != nil {
			return err
		}

		id := strconv.Itoa(int(shop.GetID()))
		v := r.HSetNX(common.ShopDt.Name(), id, data)
		err = v.Err()
		if err != nil {
			return err
		}

	}

	return nil
}
