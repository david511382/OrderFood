package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"orderfood/src/database/common"
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
		output []interface{}
	}{
		{
			name: "add 6",
			input: &models.Member{
				ID:       i,
				Name:     n,
				Username: un,
				Password: p,
			},
			err: nil,
			output: []interface{}{
				i,
				n,
				un,
				p,
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			data := &input
			err := IRedis.AddMember(data)
			assert.Equal(t, flag.err, err)

			output := []interface{}{data.GetID(), data.GetName(), data.GetUsername(), data.GetPassword()}
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
		{
			name:  "get 5",
			input: &(memberDbMembers[4]),
			err:   nil,
			output: []*models.Member{
				&(memberDbMembers[4]),
			},
		},
	}

	for _, flag := range flagtests {
		t.Run(flag.name, func(t *testing.T) {
			input := *flag.input
			output, err := IRedis.GetMember(&input)
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
			name: "update 1",
			input: models.Member{
				ID:       memberDbMembers[0].GetID(),
				Name:     new,
				Username: new,
				Password: new,
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
			name: "update 4 name",
			input: models.Member{
				ID:   memberDbMembers[3].GetID(),
				Name: new,
			},
			err:    nil,
			output: 1,
		},
		{
			name: "update 7",
			input: models.Member{
				ID:   7,
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
			output, err := IRedis.UpdateMember(inputp)
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
			name: "delete 1 id",
			input: models.Member{
				ID: memberDbMembers[0].GetID(),
			},
			err:    nil,
			output: 1,
		},
		{
			name: "delete 2 name",
			input: models.Member{
				Name: memberDbMembers[1].GetName(),
			},
			err:    common.DbDataError,
			output: 0,
		},
		{
			name: "delete 3 username",
			input: models.Member{
				Username: memberDbMembers[2].GetUsername(),
			},
			err:    common.DbDataError,
			output: 0,
		},
		{
			name: "delete 4 password",
			input: models.Member{
				Password: memberDbMembers[3].GetPassword(),
			},
			err:    common.DbDataError,
			output: 0,
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
			output, err := IRedis.DeleteMember(inputp)
			assert.Equal(t, flag.err, err)
			assert.Equal(t, flag.output, output)
		})
	}
}
