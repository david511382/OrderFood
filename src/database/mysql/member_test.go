package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

var (
	dbMember = models.Member{
		ID:       1,
		Name:     "test name",
		Username: "test username",
		Password: "test password",
	}
	memberInput = models.Member{
		Name:     dbMember.GetName(),
		Username: dbMember.GetUsername(),
		Password: dbMember.GetPassword(),
	}
)

func TestAddMember(t *testing.T) {
	output1 := dbMember
	member := memberInput

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
	member := memberInput
	err := memberDb.AddMember(&member)
	assert.Nil(t, err)

	flagtests := []struct {
		name   string
		input  *models.Member
		err    error
		output []*models.Member
	}{
		{
			name: "get 1 id",
			input: &models.Member{
				ID: dbMember.GetID(),
			},
			err: nil,
			output: []*models.Member{
				&dbMember,
			},
		},
		{
			name: "get 1 name",
			input: &models.Member{
				Name: dbMember.GetName(),
			},
			err: nil,
			output: []*models.Member{
				&dbMember,
			},
		},
		{
			name: "get 1 username",
			input: &models.Member{
				Username: dbMember.GetUsername(),
			},
			err: nil,
			output: []*models.Member{
				&dbMember,
			},
		},
		{
			name: "get 1 password",
			input: &models.Member{
				Password: dbMember.GetPassword(),
			},
			err: nil,
			output: []*models.Member{
				&dbMember,
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
func TestUpdateMember(t *testing.T) {
	member := memberInput
	err := memberDb.AddMember(&member)
	assert.Nil(t, err)

	newMember := &models.Member{
		ID:       dbMember.GetID(),
		Name:     "new name",
		Username: "new username",
		Password: "new password",
	}

	flagtests := []struct {
		name   string
		input  models.Member
		err    error
		output int64
	}{
		{
			name: "update 1 name",
			input: models.Member{
				ID:   dbMember.GetID(),
				Name: newMember.GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 1 username",
			input: models.Member{
				ID:       dbMember.GetID(),
				Username: newMember.GetUsername(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 1 password",
			input: models.Member{
				ID:       dbMember.GetID(),
				Password: newMember.GetPassword(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "update 1",
			input:  member,
			err:    nil,
			output: 1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := flag.input
			inputp := &input
			output, err := memberDb.UpdateMember(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
func TestDeleteMember(t *testing.T) {
	member := memberInput

	flagtests := []struct {
		name   string
		input  models.Member
		err    error
		output int64
	}{
		{
			name: "delete 1 id",
			input: models.Member{
				ID:       dbMember.GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 name",
			input: models.Member{
				Name: dbMember.GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 username",
			input: models.Member{
				Username: dbMember.GetUsername(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 1 password",
			input: models.Member{
				Password: dbMember.GetPassword(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 1",
			input:  dbMember,
			err:    nil,
			output: 1,
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			err := memberDb.AddMember(&member)
			assert.Nil(t, err)

			input := flag.input
			inputp := &input
			output, err := memberDb.DeleteMember(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
