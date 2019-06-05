package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

var (
	dbItem = models.Item{
		ID:      1,
		Name:    "test name",
		Price:   100,
		Shop_ID: 123,
	}
	itemInput = models.Item{
		Name:    dbItem.GetName(),
		Price:   dbItem.GetPrice(),
		Shop_ID: dbItem.GetShop_ID(),
	}
)

func TestAddItem(t *testing.T) {
	output1 := dbItem
	item := itemInput

	flagtests := []struct {
		name   string
		input  *models.Item
		err    error
		output *models.Item
	}{
		{
			name:   "add 1",
			input:  &item,
			err:    nil,
			output: &output1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := menuDb.AddItem(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetItem(t *testing.T) {
	item := itemInput
	err := menuDb.AddItem(&item)
	assert.Nil(t, err)

	flagtests := []struct {
		name   string
		input  *models.Item
		err    error
		output []*models.Item
	}{
		{
			name: "get 1 id",
			input: &models.Item{
				ID: dbItem.GetID(),
			},
			err: nil,
			output: []*models.Item{
				&dbItem,
			},
		},
		{
			name: "get 1 name",
			input: &models.Item{
				Name: dbItem.GetName(),
			},
			err: nil,
			output: []*models.Item{
				&dbItem,
			},
		},
		{
			name: "get 1 price",
			input: &models.Item{
				Price: dbItem.GetPrice(),
			},
			err: nil,
			output: []*models.Item{
				&dbItem,
			},
		},
		{
			name: "get 1 shop_id",
			input: &models.Item{
				Shop_ID: dbItem.GetShop_ID(),
			},
			err: nil,
			output: []*models.Item{
				&dbItem,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetItem(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateItem(t *testing.T) {
	item := itemInput
	err := menuDb.AddItem(&item)
	assert.Nil(t, err)

	newItem := &models.Item{
		ID:      dbItem.GetID(),
		Name:    "new name",
		Price:   1531,
		Shop_ID: 15234,
	}

	flagtests := []struct {
		name   string
		input  models.Item
		err    error
		output int64
	}{
		{
			name: "update 1 name",
			input: models.Item{
				ID:   dbItem.GetID(),
				Name: newItem.GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 1 price",
			input: models.Item{
				ID:    dbItem.GetID(),
				Price: newItem.GetPrice(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 1 shop_id",
			input: models.Item{
				ID:      dbItem.GetID(),
				Shop_ID: newItem.GetShop_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "update 1",
			input:  item,
			err:    nil,
			output: 1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateItem(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteItem(t *testing.T) {
	item := itemInput

	flagtests := []struct {
		name   string
		input  models.Item
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Item{
				ID: dbItem.GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 name",
			input: models.Item{
				Name: dbItem.GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 price",
			input: models.Item{
				Price: dbItem.GetPrice(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 shop_id",
			input: models.Item{
				Shop_ID: dbItem.GetShop_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 1",
			input:  dbItem,
			err:    nil,
			output: 1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			err := menuDb.AddItem(&item)
			assert.Nil(t, err)

			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteItem(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
