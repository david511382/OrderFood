package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddOption(t *testing.T) {
	const (
		li int32 = 3
	)

	flagtests := []struct {
		name   string
		input  *models.Option
		err    error
		output *models.Option
	}{
		{
			name: "add 4",
			input: &models.Option{
				Least_Select_Num: li,
			},
			err: nil,
			output: &models.Option{
				ID:               4,
				Least_Select_Num: li,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := menuDb.AddOption(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestGetOption(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.Option
		err    error
		output []*models.Option
	}{
		{
			name: "get 1 id",
			input: &models.Option{
				ID:               menuDbOptions[0].GetID(),
				Least_Select_Num: -1,
			},
			err: nil,
			output: []*models.Option{
				&(menuDbOptions[0]),
			},
		},
		{
			name: "get 2 least_select_num",
			input: &models.Option{
				Least_Select_Num: menuDbOptions[1].GetLeast_Select_Num(),
			},
			err: nil,
			output: []*models.Option{
				&(menuDbOptions[1]),
			},
		},
		{
			name:  "get 3",
			input: &(menuDbOptions[2]),
			err:   nil,
			output: []*models.Option{
				&(menuDbOptions[2]),
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := menuDb.GetOption(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestUpdateOption(t *testing.T) {
	const (
		newLI int32 = 4
	)

	flagtests := []struct {
		name   string
		input  models.Option
		err    error
		output int64
	}{
		{
			name: "update 1",
			input: models.Option{
				ID:               menuDbOptions[0].GetID(),
				Least_Select_Num: newLI,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 5",
			input: models.Option{
				ID:               5,
				Least_Select_Num: newLI,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.UpdateOption(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteOption(t *testing.T) {
	flagtests := []struct {
		name   string
		input  models.Option
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Option{
				ID:               menuDbOptions[0].GetID(),
				Least_Select_Num: -1,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 least_select_num",
			input: models.Option{
				Least_Select_Num: menuDbOptions[1].GetLeast_Select_Num(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 3",
			input:  menuDbOptions[2],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 5",
			input: models.Option{
				ID:               5,
				Least_Select_Num: -1,
			},
			err:    nil,
			output: 0,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := menuDb.DeleteOption(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
