package member

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"
)

type MemberDb struct {
}

func (db *MemberDb) GetMembers() ([]models.Member, error) {
	imembers, err := orm.MemberDT.Select(nil)
	if err != nil {
		return nil, err
	}

	members := make([]models.Member, 0)
	for _, v := range imembers {
		member, ok := v.(*models.Member)
		if !ok {
			return nil, common.UndefinedError
		}

		members = append(members, *member)
	}

	return members, nil
}

func (db *MemberDb) AddMembers(member *models.Member) error {
	err := orm.MemberDT.Insert(member)
	return err

	// f.Sync()

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered" + "\n")
	// fmt.Printf("wrote %d bytes\n", n4)

	// w.Flush()
}

func (db *MemberDb) UpdateMembers(member models.Member) error {
	err:=orm.MemberDT.Update(&member ,func(c interface{})bool{
		m:= c.(*models.Member)
		if m.GetUsername() == member.GetUsername(){
			return true
		}
		return false
	})

	return err
}
