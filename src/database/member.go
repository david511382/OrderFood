package database

import (
	"orderfood/src/database/models"
)

type memberDbSwitch struct {
	redisStatus bool
}

func (d *memberDbSwitch) initRedis() error {
	members, err := redisMemberDb.GetMember(nil)
	if err != nil {
		d.redisStatus = false
		return nil
	}

	if len(members) == 0 {
		members, err = memberDb.GetMember(nil)
		if err != nil {
			d.redisStatus = false
			return err
		}

		for _, member := range members {
			err = redisMemberDb.AddMember(member)
			if err != nil {
				d.redisStatus = false
				return nil
			}
		}
	}

	d.redisStatus = true
	return nil
}

func (d *memberDbSwitch) GetMember(member *models.Member) ([]*models.Member, error) {
	if d.redisStatus {
		result, err := redisMemberDb.GetMember(member)
		if err == nil {
			return result, nil
		}
	}

	d.redisStatus = false

	result, err := memberDb.GetMember(member)
	return result, err
}

func (d *memberDbSwitch) AddMember(member *models.Member) error {
	err := memberDb.AddMember(member)
	if err != nil {
		return err
	}

	if d.redisStatus {
		err = redisMemberDb.AddMember(member)
		if err != nil {
			d.redisStatus = false
		}
	}
	return nil
}
func (d *memberDbSwitch) DeleteMember(member *models.Member) (int64, error) {
	count, err := memberDb.DeleteMember(member)
	if err != nil {
		return count, err
	}

	if d.redisStatus {
		redisCount, err := redisMemberDb.DeleteMember(member)
		if err != nil || count != redisCount {
			d.redisStatus = false
		}
	}
	return count, nil
}
func (d *memberDbSwitch) UpdateMember(member *models.Member) (int64, error) {
	count, err := memberDb.UpdateMember(member)
	if err != nil {
		return count, err
	}

	if d.redisStatus {
		redisCount, err := redisMemberDb.UpdateMember(member)
		if err != nil || count != redisCount {
			d.redisStatus = false
		}
	}
	return count, nil
}
