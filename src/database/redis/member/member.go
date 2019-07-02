package member

import (
	"database/sql"
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"strconv"

	linq "github.com/ahmetb/go-linq"
	proto "github.com/golang/protobuf/proto"
)

func (redis *RedisDb) GetMember(member *models.Member) ([]*models.Member, error) {
	r := redis.R
	members := make([]*models.Member, 0)
	id := strconv.Itoa(int(member.GetID()))
	if id != "0" {
		v := r.HGet(common.MemberDt.Name(), id)
		err := v.Err()
		if err != nil {
			return nil, err
		}

		b, err := v.Bytes()
		if err != nil {
			return nil, err
		}

		nm := &models.Member{}
		err = proto.Unmarshal(b, nm)
		if err != nil {
			return nil, err
		}

		members = append(members, nm)
		return members, nil
	}

	v := r.HGetAll(common.MemberDt.Name())
	err := v.Err()
	if err != nil {
		return nil, err
	}

	memberMap, err := v.Result()
	if err != nil {
		return nil, err
	}

	for _, memberStr := range memberMap {
		b := []byte(memberStr)
		nm := &models.Member{}

		err := proto.Unmarshal(b, nm)
		if err != nil {
			return nil, err
		}
		members = append(members, nm)
	}

	linq.From(members).Where(func(m interface{}) bool {
		mem := m.(*models.Member)

		if member.GetName() != "" {
			if mem.GetName() != member.GetName() {
				return false
			}
		}

		if member.GetUsername() != "" {
			if mem.GetUsername() != member.GetUsername() {
				return false
			}
		}

		if member.GetPassword() != "" {
			if mem.GetPassword() != member.GetPassword() {
				return false
			}
		}

		return true
	}).ToSlice(&members)

	return members, nil
}
func (redis *RedisDb) AddMember(member *models.Member, tx *sql.Tx) error {
	data, err := proto.Marshal(member)
	if err != nil {
		return err
	}

	r := redis.R
	id := strconv.Itoa(int(member.GetID()))
	v := r.HSetNX(common.MemberDt.Name(), id, data)
	if !v.Val() {
		return common.InserFailError
	}
	return nil
}
func (redis *RedisDb) UpdateMember(member *models.Member, tx *sql.Tx) (int64, error) {
	r := redis.R
	id := strconv.Itoa(int(member.GetID()))
	v := r.HExists(common.MemberDt.Name(), id)
	err := v.Err()
	if err != nil {
		return 0, err
	}
	if !v.Val() {
		return 0, nil
	}

	data, err := proto.Marshal(member)
	if err != nil {
		return 0, err
	}

	v = r.HSet(common.MemberDt.Name(), id, data)
	err = v.Err()
	if err != nil {
		return 0, err
	}
	if v.Val() {
		return 0, common.UpdateFailError
	}
	return 1, nil
}
func (redis *RedisDb) DeleteMember(member *models.Member, tx *sql.Tx) (int64, error) {
	r := redis.R
	id := strconv.Itoa(int(member.GetID()))
	if id == "0" {
		return 0, common.DbDataError
	}

	v := r.HDel(common.MemberDt.Name(), id)
	err := v.Err()
	if err != nil {
		return 0, err
	}

	count := v.Val()

	return count, nil
}
