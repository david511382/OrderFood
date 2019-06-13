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
				Shop_ID: intP(2),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(2),
					Item_ID:   intP(menuDbItems[2].GetID()),
					Option_ID: intP(3),
					Name:      stringP(menuDbItems[2].GetName()),
					Price:     intP(menuDbItems[2].GetPrice()),
				},
			},
		},
		{
			name: "get 2 item_id",
			input: &models.ItemOptionView{
				Item_ID: intP(menuDbItems[1].GetID()),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(1),
					Item_ID:   intP(menuDbItems[1].GetID()),
					Option_ID: intP(2),
					Name:      stringP(menuDbItems[1].GetName()),
					Price:     intP(menuDbItems[1].GetPrice()),
				},
			},
		},
		{
			name: "get 3 option_id",
			input: &models.ItemOptionView{
				Option_ID: intP(menuDbOptions[2].GetID()),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(2),
					Item_ID:   intP(menuDbItems[2].GetID()),
					Option_ID: intP(3),
					Name:      stringP(menuDbItems[2].GetName()),
					Price:     intP(menuDbItems[2].GetPrice()),
				},
			},
		},
		{
			name: "get 4 name",
			input: &models.ItemOptionView{
				Name: stringP(menuDbItems[4].GetName()),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(1),
					Item_ID:   intP(menuDbItems[4].GetID()),
					Option_ID: nil,
					Name:      stringP(menuDbItems[4].GetName()),
					Price:     intP(menuDbItems[4].GetPrice()),
				},
			},
		},
		{
			name: "get 5 price",
			input: &models.ItemOptionView{
				Price: intP(menuDbItems[3].GetPrice()),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(1),
					Item_ID:   intP(menuDbItems[3].GetID()),
					Option_ID: intP(1),
					Name:      stringP(menuDbItems[3].GetName()),
					Price:     intP(menuDbItems[3].GetPrice()),
				},
			},
		},
		{
			name: "get 6",
			input: &models.ItemOptionView{
				Shop_ID:   intP(1),
				Item_ID:   intP(menuDbItems[0].GetID()),
				Option_ID: intP(1),
				Name:      stringP(menuDbItems[0].GetName()),
				Price:     intP(menuDbItems[0].GetPrice()),
			},
			err: nil,
			output: []*models.ItemOptionView{
				&models.ItemOptionView{
					Shop_ID:   intP(1),
					Item_ID:   intP(menuDbItems[0].GetID()),
					Option_ID: intP(1),
					Name:      stringP(menuDbItems[0].GetName()),
					Price:     intP(menuDbItems[0].GetPrice()),
				},
			},
		},
		{
			name: "get 7",
			input: &models.ItemOptionView{
				Shop_ID:   intP(2),
				Item_ID:   intP(menuDbItems[1].GetID()),
				Option_ID: intP(1),
				Name:      stringP(menuDbItems[1].GetName()),
				Price:     intP(menuDbItems[1].GetPrice()),
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
