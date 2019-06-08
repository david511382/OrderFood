package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddSelection(t *testing.T) {
	const (
		s = "3"
	)

	flagtests := []struct {
		name   string
		input  *models.Selection
		err    error
		output *models.Selection
	}{
		{
			name: "add 4",
			input: &models.Selection{
				Name: s,
			},
			err: nil,
			output: &models.Selection{
				ID:               4,
				Name: s,
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
				ID:               menuDbSelections[0].GetID(),
				
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[0]),
			},
		},
		{
			name: "get 2 name",
			input: &models.Selection{
				Name: menuDbSelections[1].GetName(),
			},
			err: nil,
			output: []*models.Selection{
				&(menuDbSelections[1]),
			},
		},
		{
			name:  "get 3",
			input: &(menuDbSelections[2]),
			err:   nil,
			output: []*models.Selection{
				&(menuDbSelections[2]),
			},
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
		new  = "new"
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
				ID:               menuDbSelections[0].GetID(),
				Name: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 5",
			input: models.Selection{
				ID:               5,
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
				ID:               menuDbSelections[0].GetID(),
				
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 name",
			input: models.Selection{
				Name: menuDbSelections[1].GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 3",
			input:  menuDbSelections[2],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 5",
			input: models.Selection{
				ID:               5,
				
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
