package txt

import (
	"bufio"
	"io"
	"orderfood/src/database/models"

	proto "github.com/golang/protobuf/proto"
)

func (db *txtDb) GetMembers() ([]models.Member, error) {
	f, err := db.Connect("order_member.user_info.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	members := make([]models.Member, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		member := &models.Member{}
		err = proto.Unmarshal(line, member)
		if err != nil {
			return nil, err
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

	data := member.String()
	_, err = f.WriteString(data + "\n")
	return err

	// f.Sync()

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered" + "\n")
	// fmt.Printf("wrote %d bytes\n", n4)

	// w.Flush()
}
