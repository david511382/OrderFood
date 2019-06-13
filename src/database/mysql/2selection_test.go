package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddSelection(t *testing.T) {
	const (
		n  string = "test"
		oi int  = 1
		p  int  = 154531
	)

	flagtests := []struct {
		name   string
		input  *models.Selection
		err    error
		output *models.Selection
	}{
		{
			name: "add 6",
			input: &models.Selection{
				Name:      n,
				Option_ID: oi,
				Price:     p,
			},
			err: nil,
			output: &models.Selection{
				ID:        6,
				Name:      n,
				Option_ID: oi,
				Price:     p,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := menuDb.AddSelection(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetSelection(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.Selection
		err    error
		output []*models.Selection
	}{
		{
			name: "get 1 id",
			input: &models.Selection{
				ID:    menuDbSelections[0].GetID(),
				Price: -1,
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[0]),
			},
		},
		{
			name: "get 2 name",
			input: &models.Selection{
				Name:  menuDbSelections[1].GetName(),
				Price: -1,
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[1]),
			},
		},
		{
			name: "get 3 option_id",
			input: &models.Selection{
				Option_ID: menuDbSelections[2].GetOption_ID(),
				Price:     -1,
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[2]),
			},
		},
		{
			name: "get 4 Price",
			input: &models.Selection{
				Price: menuDbSelections[3].GetPrice(),
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[3]),
			},
		},
		{
			name:  "get 5",
			input: &(menuDbSelections[4]),
			err:   nil,
			output: []*models.Selection{
				&(menuDbSelections[4]),
			},
		},
		{
			name: "get 7",
			input: &models.Selection{
				ID:    7,
				Price: -1,
			},
			err:    nil,
			output: []*models.Selection{},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetSelection(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateSelection(t *testing.T) {
	const (
		new       = "new"
		oi  int = 1
		p   int = 87641
	)

	flagtests := []struct {
		name   string
		input  models.Selection
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.Selection{
				ID:    menuDbSelections[0].GetID(),
				Name:  new,
				Option_ID: oi,
				Price: -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 option_id",
			input: models.Selection{
				ID:        menuDbSelections[1].GetID(),
				Option_ID: oi,
				Price:     -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 3 price",
			input: models.Selection{
				ID:    menuDbSelections[2].GetID(),
				Price: p,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 4 name",
			input: models.Selection{
				ID:    menuDbSelections[3].GetID(),
				Name:  new,
				Price: -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 7",
			input: models.Selection{
				ID:    7,
				Price: p,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateSelection(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteSelection(t *testing.T) {
	flagtests := []struct {
		name   string
		input  models.Selection
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Selection{
				ID:    menuDbSelections[0].GetID(),
				Price: -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 name",
			input: models.Selection{
				Name:  menuDbSelections[1].GetName(),
				Price: -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 3 option_id",
			input: models.Selection{
				Option_ID: menuDbSelections[2].GetOption_ID(),
				Price:     -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 4 price",
			input: models.Selection{
				Price: menuDbSelections[3].GetPrice(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 7",
			input: models.Selection{
				ID:    7,
				Price: -1,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteSelection(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
