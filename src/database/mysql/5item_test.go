package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddItem(t *testing.T) {
	const (
		n  string = "fjdsakl;tg"
		si int  = 1
		p  int  = 15315
	)

	flagtests := []struct {
		name   string
		input  *models.Item
		err    error
		output *models.Item
	}{
		{
			name: "add 6",
			input: &models.Item{
				Name:    n,
				Shop_ID: si,
				Price:   p,
			},
			err: nil,
			output: &models.Item{
				ID:      6,
				Name:    n,
				Shop_ID: si,
				Price:   p,
			},
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
	flagtests := []struct {
		name   string
		input  *models.Item
		err    error
		output []*models.Item
	}{
		{
			name: "get 1 id",
			input: &models.Item{
				ID: menuDbItems[0].GetID(),
				Price:   -1,
			},
			err: nil,
			output: []*models.Item{
				&(menuDbItems[0]),
			},
		},
		{
			name: "get 2 name",
			input: &models.Item{
				Name: menuDbItems[1].GetName(),
				Price:   -1,
			},
			err: nil,
			output: []*models.Item{
				&(menuDbItems[1]),
			},
		},
		{
			name: "get 3 shop_id",
			input: &models.Item{
				Shop_ID: menuDbItems[2].GetShop_ID(),
				Price:   -1,
			},
			err: nil,
			output: []*models.Item{
				&(menuDbItems[2]),
			},
		},
		{
			name: "get 4 price",
			input: &models.Item{
				Price: menuDbItems[3].GetPrice(),
			},
			err: nil,
			output: []*models.Item{
				&(menuDbItems[3]),
			},
		},
		{
			name:  "get 5",
			input: &(menuDbItems[4]),
			err:   nil,
			output: []*models.Item{
				&(menuDbItems[4]),
			},
		},
		{
			name: "get 7",
			input: &models.Item{
				ID: 7,
				Price:   -1,
			},
			err:    nil,
			output: []*models.Item{},
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
	const (
		new1  string = "new"
		new2  string = "new2"
		newP  int  = 5465347
		newSI int  = 3
	)

	flagtests := []struct {
		name   string
		input  models.Item
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.Item{
				ID:      menuDbItems[0].GetID(),
				Name:    new1,
				Shop_ID: newSI,
				Price:   newP,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 shop_id",
			input: models.Item{
				ID:      menuDbItems[1].GetID(),
				Shop_ID: newSI,
				Price:   -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 3 price",
			input: models.Item{
				ID:    menuDbItems[2].GetID(),
				Price: newP,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 4 name",
			input: models.Item{
				ID:   menuDbItems[3].GetID(),
				Name: new2,
				Price:   -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 7",
			input: models.Item{
				ID:   7,
				Name: new2,
				Price:   -1,
			},
			err:    nil,
			output: 0,
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
	flagtests := []struct {
		name   string
		input  models.Item
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Item{
				ID: menuDbItems[0].GetID(),
				Price:   -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 name",
			input: models.Item{
				Name: menuDbItems[1].GetName(),
				Price:   -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 3 shop_id",
			input: models.Item{
				Shop_ID: menuDbItems[2].GetShop_ID(),
				Price:   -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 4 price",
			input: models.Item{
				Price: menuDbItems[3].GetPrice(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 5",
			input:  menuDbItems[4],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 7",
			input: models.Item{
				ID: 7,
				Price:   -1,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteItem(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
