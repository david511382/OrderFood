package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddItemOption(t *testing.T) {
	const (
		i  int32 = 5
		ii int32 = 3
		oi int32 = 2
	)

	flagtests := []struct {
		name   string
		input  *models.ItemOption
		err    error
		output *models.ItemOption
	}{
		{
			name: "add 5",
			input: &models.ItemOption{
				Item_ID:   ii,
				Option_ID: oi,
			},
			err: nil,
			output: &models.ItemOption{
				ID:        i,
				Item_ID:   ii,
				Option_ID: oi,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := menuDb.AddItemOption(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetItemOption(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.ItemOption
		err    error
		output []*models.ItemOption
	}{
		{
			name: "get 1 id",
			input: &models.ItemOption{
				ID: menuDbItemOptions[0].GetID(),
			},
			err: nil,
			output: []*models.ItemOption{
				&(menuDbItemOptions[0]),
			},
		},
		{
			name: "get 2 item_id",
			input: &models.ItemOption{
				Item_ID: menuDbItemOptions[1].GetItem_ID(),
			},
			err: nil,
			output: []*models.ItemOption{
				&(menuDbItemOptions[1]),
			},
		},
		{
			name: "get 3 option_id",
			input: &models.ItemOption{
				Option_ID: menuDbItemOptions[2].GetOption_ID(),
			},
			err: nil,
			output: []*models.ItemOption{
				&(menuDbItemOptions[2]),
			},
		},
		{
			name:  "get 4",
			input: &(menuDbItemOptions[3]),
			err:   nil,
			output: []*models.ItemOption{
				&(menuDbItemOptions[3]),
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetItemOption(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateItemOption(t *testing.T) {
	const (
		newII int32 = 4
		newOI int32 = 4
	)

	flagtests := []struct {
		name   string
		input  models.ItemOption
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.ItemOption{
				ID:        menuDbItemOptions[0].GetID(),
				Item_ID:   newII,
				Option_ID: newOI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 option_id",
			input: models.ItemOption{
				ID:        menuDbItemOptions[1].GetID(),
				Option_ID: newOI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 3 item_id",
			input: models.ItemOption{
				ID:      menuDbItemOptions[2].GetID(),
				Item_ID: newII,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 6",
			input: models.ItemOption{
				ID:      6,
				Item_ID: newII,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateItemOption(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteItemOption(t *testing.T) {
	flagtests := []struct {
		name   string
		input  models.ItemOption
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.ItemOption{
				ID: menuDbItemOptions[0].GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 item_id",
			input: models.ItemOption{
				Item_ID: menuDbItemOptions[1].GetItem_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 3 option_id",
			input: models.ItemOption{
				Option_ID: menuDbItemOptions[2].GetOption_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 4",
			input:  menuDbItemOptions[3],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 6",
			input: models.ItemOption{
				ID: 6,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteItemOption(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
