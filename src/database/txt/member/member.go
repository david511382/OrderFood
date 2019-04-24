package member

import (
	"bufio"
	"io"
	"io/ioutil"
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"

	proto "github.com/golang/protobuf/proto"
)

type MemberDb struct {
}

func (db *MemberDb) GetMembers() ([]models.Member, error) {
	imembers, err := orm.MemberDT.Select().Exec()
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

func (db *MemberDb) AddMembers(member models.Member) error {
	err := orm.MemberDT.Insert(&member)
	return err

	// f.Sync()

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered" + "\n")
	// fmt.Printf("wrote %d bytes\n", n4)

	// w.Flush()
}

func (db *MemberDb) UpdateMembers(member models.Member) error {
	f, filepath, err := orm.Connect("order_member.user_info.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	id := int(member.GetID())

	allData := make([]byte, 0)

	const n = 10
	reader := bufio.NewReader(f)
	for i := 1; i != id; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}

		allData = append(allData, line...)
		allData = append(allData, n)
	}

	_, _, err = reader.ReadLine()
	if err != nil {
		return err
	}

	newData, err := proto.Marshal(&member)
	if err != nil {
		return err
	}

	allData = append(allData, []byte(newData)...)
	allData = append(allData, n)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		allData = append(allData, line...)
		allData = append(allData, n)
	}

	err = ioutil.WriteFile(filepath, allData, 0644)

	return err
}
