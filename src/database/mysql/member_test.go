package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

var (
	memberOutput = models.Member{
		ID:       1,
		Name:     "test name",
		Username: "test username",
		Password: "test password",
	}
)

func TestAddMember(t *testing.T) {
	output1 := memberOutput
	member := models.Member{
		ID:       0,
		Name:     memberOutput.GetName(),
		Username: memberOutput.GetUsername(),
		Password: memberOutput.GetPassword(),
	}

	flagtests := []struct {
		name   string
		input  *models.Member
		err    error
		output *models.Member
	}{
		{
			name:   "add 1",
			input:  &member,
			err:    nil,
			output: &output1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output := &input
			err := memberDb.AddMember(output)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}

func TestGetMember(t *testing.T) {
	flagtests := []struct {
		name   string
		input  *models.Member
		err    error
		output []models.Member
	}{
		{
			name: "get 1 id",
			input: &models.Member{
				ID: memberOutput.GetID(),
			},
			err: nil,
			output: []models.Member{
				memberOutput,
			},
		},
		{
			name: "get 1 name",
			input: &models.Member{
				Name: memberOutput.GetName(),
			},
			err: nil,
			output: []models.Member{
				memberOutput,
			},
		},
		{
			name: "get 1 username",
			input: &models.Member{
				Username: memberOutput.GetUsername(),
			},
			err: nil,
			output: []models.Member{
				memberOutput,
			},
		},
		{
			name: "get 1 password",
			input: &models.Member{
				Password: memberOutput.GetPassword(),
			},
			err: nil,
			output: []models.Member{
				memberOutput,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := memberDb.GetMember(&input)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
