package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/common"
	"orderfood/src/database/models"
)

func TestAddShop(t *testing.T) {
	const (
		i int32  = 6
		n string = "fjdsakl;tg"
	)

	flagtests := []struct {
		name   string
		input  *models.Shop
		err    error
		output []interface{}
	}{
		{
			name: "add 6",
			input: &models.Shop{
				ID:   i,
				Name: n,
			},
			err: nil,
			output: []interface{}{
				i,
				n,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			data := &input
			err := menuDb.AddShop(data,nil)
			assert.Equal(t, flag.err, err)

			output := []interface{}{data.GetID(), data.GetName()}
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetShop(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.Shop
		err    error
		output []*models.Shop
	}{
		{
			name: "get 1 id",
			input: &models.Shop{
				ID: menuDbShops[0].GetID(),
			},
			err: nil,
			output: []*models.Shop{
				&(menuDbShops[0]),
			},
		},
		{
			name: "get 2 name",
			input: &models.Shop{
				Name: menuDbShops[1].GetName(),
			},
			err: nil,
			output: []*models.Shop{
				&(menuDbShops[1]),
			},
		},
		{
			name:  "get 3",
			input: &(menuDbShops[2]),
			err:   nil,
			output: []*models.Shop{
				&(menuDbShops[2]),
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetShop(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateShop(t *testing.T) {
	const new = "new"

	flagtests := []struct {
		name   string
		input  models.Shop
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.Shop{
				ID:   menuDbShops[0].GetID(),
				Name: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 name",
			input: models.Shop{
				ID:   menuDbShops[1].GetID(),
				Name: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 5",
			input: models.Shop{
				ID:   5,
				Name: new,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateShop(inputp,nil)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteShop(t *testing.T) {
	flagtests := []struct {
		name   string
		input  models.Shop
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Shop{
				ID: menuDbShops[0].GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 name",
			input: models.Shop{
				Name: menuDbShops[1].GetName(),
			},
			err:    common.DbDataError,
			output: 0,
		},
		{
			name:   "delete 3",
			input:  menuDbShops[2],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 5",
			input: models.Shop{
				ID: 5,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteShop(inputp,nil)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
