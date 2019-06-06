package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/models"
)

func TestAddMember(t *testing.T) {
	const (
		i  int32  = 6
		n  string = "fjdsakl;tg"
		un string = "fjdsakl;ffjslfjkla"
		p  string = "fdsagewgege"
	)

	flagtests := []struct {
		name   string
		input  *models.Member
		err    error
		output *models.Member
	}{
		{
			name: "add 6",
			input: &models.Member{
				Name:     n,
				Username: un,
				Password: p,
			},
			err: nil,
			output: &models.Member{
				ID:       i,
				Name:     n,
				Username: un,
				Password: p,
			},
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
		output []*models.Member
	}{
		{
			name: "get 1 id",
			input: &models.Member{
				ID: memberDbMembers[0].GetID(),
			},
			err: nil,
			output: []*models.Member{
				&(memberDbMembers[0]),
			},
		},
		{
			name: "get 2 name",
			input: &models.Member{
				Name: memberDbMembers[1].GetName(),
			},
			err: nil,
			output: []*models.Member{
				&(memberDbMembers[1]),
			},
		},
		{
			name: "get 3 username",
			input: &models.Member{
				Username: memberDbMembers[2].GetUsername(),
			},
			err: nil,
			output: []*models.Member{
				&(memberDbMembers[2]),
			},
		},
		{
			name: "get 4 password",
			input: &models.Member{
				Password: memberDbMembers[3].GetPassword(),
			},
			err: nil,
			output: []*models.Member{
				&(memberDbMembers[3]),
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
	const new = "new"
	flagtests := []struct {
		name   string
		input  models.Member
		err    error
		output int64
	}{
		{
			name: "update 1 name",
			input: models.Member{
				ID:   memberDbMembers[0].GetID(),
				Name: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 2 username",
			input: models.Member{
				ID:       memberDbMembers[1].GetID(),
				Username: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 3 password",
			input: models.Member{
				ID:       memberDbMembers[2].GetID(),
				Password: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 4",
			input: models.Member{
				ID:       memberDbMembers[3].GetID(),
				Name:     new,
				Username: new,
				Password: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "update 5",
			input:  memberDbMembers[4],
			err:    nil,
			output: 0,
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
	flagtests := []struct {
		name   string
		input  models.Member
		err    error
		output int64
	}{
		{
			name: "delete 1 username",
			input: models.Member{
				Username: memberDbMembers[0].GetUsername(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 password",
			input: models.Member{
				Password: memberDbMembers[1].GetPassword(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 3 name",
			input: models.Member{
				Name: memberDbMembers[2].GetName(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 4 id",
			input: models.Member{
				ID: memberDbMembers[3].GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name:   "delete 5",
			input:  memberDbMembers[4],
			err:    nil,
			output: 1,
		},
		{
			name: "delete 7",
			input: models.Member{
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
			output, err := memberDb.DeleteMember(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
