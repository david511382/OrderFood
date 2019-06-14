package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestGetItemOptionView(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.ItemOptionView
		err    error
		output []*models.ItemOptionView
	}{
		{
			name: "get 1 shop_id",
			input: &models.ItemOptionView{
				Shop_ID: 2,
				Price: -1,
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   2,
					Item_ID:   menuDbItems[2].GetID(),
					Option_ID: intP(3),
					Name:      menuDbItems[2].GetName(),
					Price:     menuDbItems[2].GetPrice(),
				},
			},
		},
		{
			name: "get 2 item_id",
			input: &models.ItemOptionView{
				Item_ID: menuDbItems[1].GetID(),
				Price: -1,
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   1,
					Item_ID:   menuDbItems[1].GetID(),
					Option_ID: intP(2),
					Name:      menuDbItems[1].GetName(),
					Price:     menuDbItems[1].GetPrice(),
				},
			},
		},
		{
			name: "get 3 option_id",
			input: &models.ItemOptionView{
				Option_ID: intP(menuDbOptions[2].GetID()),
				Price: -1,
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   2,
					Item_ID:   menuDbItems[2].GetID(),
					Option_ID: intP(3),
					Name:      menuDbItems[2].GetName(),
					Price:     menuDbItems[2].GetPrice(),
				},
			},
		},
		{
			name: "get 4 name",
			input: &models.ItemOptionView{
				Name: menuDbItems[4].GetName(),
				Price: -1,
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   1,
					Item_ID:   menuDbItems[4].GetID(),
					Option_ID: nil,
					Name:      menuDbItems[4].GetName(),
					Price:     menuDbItems[4].GetPrice(),
				},
			},
		},
		{
			name: "get 5 price",
			input: &models.ItemOptionView{
				Price: menuDbItems[3].GetPrice(),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   1,
					Item_ID:   menuDbItems[3].GetID(),
					Option_ID: intP(1),
					Name:      menuDbItems[3].GetName(),
					Price:     menuDbItems[3].GetPrice(),
				},
			},
		},
		{
			name: "get 6",
			input: &models.ItemOptionView{
				Shop_ID:   1,
				Item_ID:   menuDbItems[0].GetID(),
				Option_ID: intP(1),
				Name:      menuDbItems[0].GetName(),
				Price:     menuDbItems[0].GetPrice(),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   1,
					Item_ID:   menuDbItems[0].GetID(),
					Option_ID: intP(1),
					Name:      menuDbItems[0].GetName(),
					Price:     menuDbItems[0].GetPrice(),
				},
			},
		},
		{
			name: "get 7",
			input: &models.ItemOptionView{
				Shop_ID:   2,
				Item_ID:   menuDbItems[1].GetID(),
				Option_ID: intP(1),
				Name:      menuDbItems[1].GetName(),
				Price:     menuDbItems[1].GetPrice(),
			},
			err:    nil,
			output: []*models.ItemOptionView{},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetItemOptionView(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}

func intP(i int) *int {
	return &i
}

func stringP(s string) *string {
	return &s
}
