package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestGetOptionSelectionView(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.OptionSelectionView
		err    error
		output []*models.OptionSelectionView
	}{
		{
			name: "get 1 option_id",
			input: &models.OptionSelectionView{
				Option_ID:  menuDbOptions[2].GetID(),
				Select_Num: -1,
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[2].GetID(),
					Select_Num:   menuDbOptions[2].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[2].GetID()),
					Name:         stringP(menuDbSelections[2].GetName()),
					Price:        intP(menuDbSelections[2].GetPrice()),
				},
			},
		},
		{
			name: "get 2 select_num",
			input: &models.OptionSelectionView{
				Select_Num: menuDbOptions[2].GetSelect_Num(),
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[2].GetID(),
					Select_Num:   menuDbOptions[2].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[2].GetID()),
					Name:         stringP(menuDbSelections[2].GetName()),
					Price:        intP(menuDbSelections[2].GetPrice()),
				},
			},
		},
		{
			name: "get 3 selection_id",
			input: &models.OptionSelectionView{
				Selection_ID: intP(menuDbSelections[0].GetID()),
				Select_Num:   -1,
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[0].GetID(),
					Select_Num:   menuDbOptions[0].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[0].GetID()),
					Name:         stringP(menuDbSelections[0].GetName()),
					Price:        intP(menuDbSelections[0].GetPrice()),
				},
			},
		},
		{
			name: "get 4 name",
			input: &models.OptionSelectionView{
				Name:       stringP(menuDbSelections[1].GetName()),
				Select_Num: -1,
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[1].GetID(),
					Select_Num:   menuDbOptions[1].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[1].GetID()),
					Name:         stringP(menuDbSelections[1].GetName()),
					Price:        intP(menuDbSelections[1].GetPrice()),
				},
			},
		},
		{
			name: "get 5 price",
			input: &models.OptionSelectionView{
				Price:      intP(menuDbSelections[3].GetPrice()),
				Select_Num: -1,
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[0].GetID(),
					Select_Num:   menuDbOptions[0].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[3].GetID()),
					Name:         stringP(menuDbSelections[3].GetName()),
					Price:        intP(menuDbSelections[3].GetPrice()),
				},
			},
		},
		{
			name: "get 6",
			input: &models.OptionSelectionView{
				Option_ID:    menuDbOptions[1].GetID(),
				Select_Num:   menuDbOptions[1].GetSelect_Num(),
				Selection_ID: intP(menuDbSelections[4].GetID()),
				Name:         stringP(menuDbSelections[4].GetName()),
				Price:        intP(menuDbSelections[4].GetPrice()),
			},
			err: nil,
			output: []*models.OptionSelectionView{
				&models.OptionSelectionView{
					Option_ID:    menuDbOptions[1].GetID(),
					Select_Num:   menuDbOptions[1].GetSelect_Num(),
					Selection_ID: intP(menuDbSelections[4].GetID()),
					Name:         stringP(menuDbSelections[4].GetName()),
					Price:        intP(menuDbSelections[4].GetPrice()),
				},
			},
		},
		{
			name: "get 7",
			input: &models.OptionSelectionView{
				Option_ID:    3,
				Select_Num:   0,
				Selection_ID: intP(menuDbSelections[2].GetID()),
				Name:         stringP(menuDbSelections[2].GetName()),
				Price:        intP(menuDbSelections[2].GetPrice()),
			},
			err:    nil,
			output: []*models.OptionSelectionView{},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetOptionSelectionView(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
