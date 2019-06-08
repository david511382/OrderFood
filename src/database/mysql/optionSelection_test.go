package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddOptionSelection(t *testing.T) {
	const (
		oi int32 = 3
		si int32 = 2
		p  int32 = 1531
	)

	flagtests := []struct {
		name   string
		input  *models.OptionSelection
		err    error
		output *models.OptionSelection
	}{
		{
			name: "add 6",
			input: &models.OptionSelection{
				Option_ID:    oi,
				Price:        p,
				Selection_ID: si,
			},
			err: nil,
			output: &models.OptionSelection{
				ID:           6,
				Option_ID:    oi,
				Price:        p,
				Selection_ID: si,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := menuDb.AddOptionSelection(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetOptionSelection(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.OptionSelection
		err    error
		output []*models.OptionSelection
	}{
		{
			name: "get 1 id",
			input: &models.OptionSelection{
				ID: menuDbOptionSelections[0].GetID(),
			},
			err: nil,
			output: []*models.OptionSelection{
				&(menuDbOptionSelections[0]),
			},
		},
		{
			name: "get 2 option_id",
			input: &models.OptionSelection{
				Option_ID: menuDbOptionSelections[1].GetOption_ID(),
			},
			err: nil,
			output: []*models.OptionSelection{
				&(menuDbOptionSelections[1]),
			},
		},
		{
			name: "get 3 price",
			input: &models.OptionSelection{
				Price: menuDbOptionSelections[2].GetPrice(),
			},
			err: nil,
			output: []*models.OptionSelection{
				&(menuDbOptionSelections[2]),
			},
		},
		{
			name: "get 4 selection_id",
			input: &models.OptionSelection{
				Selection_ID: menuDbOptionSelections[3].GetSelection_ID(),
			},
			err: nil,
			output: []*models.OptionSelection{
				&(menuDbOptionSelections[3]),
			},
		},
		{
			name:  "get 5",
			input: &(menuDbOptionSelections[4]),
			err:   nil,
			output: []*models.OptionSelection{
				&(menuDbOptionSelections[4]),
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetOptionSelection(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateOptionSelection(t *testing.T) {
	const (
		newSI int32 = 2
		newOI int32 = 3
		newP  int32 = 15313
	)

	flagtests := []struct {
		name   string
		input  models.OptionSelection
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.OptionSelection{
				ID:           menuDbOptionSelections[0].GetID(),
				Price:        newP,
				Selection_ID: newSI,
				Option_ID:    newOI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 price",
			input: models.OptionSelection{
				ID:    menuDbOptionSelections[1].GetID(),
				Price: newP,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 3 selection_id",
			input: models.OptionSelection{
				ID:           menuDbOptionSelections[2].GetID(),
				Selection_ID: newSI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 4 option_id",
			input: models.OptionSelection{
				ID:        menuDbOptionSelections[3].GetID(),
				Option_ID: newOI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 7",
			input: models.OptionSelection{
				ID:        7,
				Option_ID: newOI,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateOptionSelection(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteOptionSelection(t *testing.T) {
	flagtests := []struct {
		name   string
		input  models.OptionSelection
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.OptionSelection{
				ID: menuDbOptionSelections[0].GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 option_id",
			input: models.OptionSelection{
				Option_ID: menuDbOptionSelections[1].GetOption_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 3 price",
			input: models.OptionSelection{
				Price: menuDbOptionSelections[2].GetPrice(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 4 selection_id",
			input: models.OptionSelection{
				Selection_ID: menuDbOptionSelections[3].GetSelection_ID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 5",
			input:  menuDbOptionSelections[4],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 7",
			input: models.OptionSelection{
				ID: 7,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteOptionSelection(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
