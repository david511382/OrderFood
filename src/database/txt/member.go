package txt

import (
	"bufio"
	"io"
	"io/ioutil"
	"orderfood/src/database/models"

	proto "github.com/golang/protobuf/proto"
)

func (db *txtDb) GetMembers() ([]models.Member, error) {
	imembers, err := memberDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	members := make([]models.Member, 0)
	for _, v := range imembers {
		member, ok := v.(*models.Member)
		if !ok {
			return nil, undefinedError
		}

		members = append(members, *member)
	}

	return members, nil
}

func (db *txtDb) AddMembers(member models.Member) error {
	f, err := db.Connect("order_member.user_info.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := proto.Marshal(&member)
	if err != nil {
		return err
	}

	_, err = f.Write(out)
	if err != nil {
		return err
	}

	_, err = f.WriteString("\n")
	return err

	// f.Sync()

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered" + "\n")
	// fmt.Printf("wrote %d bytes\n", n4)

	// w.Flush()
}

func (db *txtDb) UpdateMembers(member models.Member) error {
	f, err := db.Connect("order_member.user_info.txt")
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

	err = ioutil.WriteFile(db.Filepath, allData, 0644)

	return err
}
