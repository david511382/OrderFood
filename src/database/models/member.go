package models

// Member 。
type Member struct {
	ID                   int    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (m *Member) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Member) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}