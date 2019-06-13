package models

// Shop 。
type Shop struct {
	ID   int  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *Shop) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Shop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Item 。
type Item struct {
	ID      int  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Shop_ID int  `protobuf:"varint,3,opt,name=Shop_ID,json=ShopID,proto3" json:"Shop_ID,omitempty"`
	Price   int  `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *Item) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetShop_ID() int {
	if m != nil {
		return m.Shop_ID
	}
	return 0
}

func (m *Item) GetPrice() int {
	if m != nil {
		return m.Price
	}
	return 0
}

// ItemOption 。
type ItemOption struct {
	ID        int `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Item_ID   int `protobuf:"varint,2,opt,name=Item_ID,json=ItemID,proto3" json:"Item_ID,omitempty"`
	Option_ID int `protobuf:"varint,3,opt,name=Option_ID,json=OptionID,proto3" json:"Option_ID,omitempty"`
}

func (m *ItemOption) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ItemOption) GetItem_ID() int {
	if m != nil {
		return m.Item_ID
	}
	return 0
}

func (m *ItemOption) GetOption_ID() int {
	if m != nil {
		return m.Option_ID
	}
	return 0
}

// Option 。
type Option struct {
	ID         int `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Select_Num int `protobuf:"varint,2,opt,name=Select_Num,json=SelectNum,proto3" json:"Select_Num,omitempty"`
}

func (m *Option) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Option) GetSelect_Num() int {
	if m != nil {
		return m.Select_Num
	}
	return 0
}

// selection
type Selection struct {
	ID        int  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Option_ID int  `protobuf:"varint,3,opt,name=Option_ID,json=OptionID,proto3" json:"Option_ID,omitempty"`
	Price     int  `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *Selection) GetID() int {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Selection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Selection) GetOption_ID() int {
	if m != nil {
		return m.Option_ID
	}
	return 0
}

func (m *Selection) GetPrice() int {
	if m != nil {
		return m.Price
	}
	return 0
}

// ItemOptionView
type ItemOptionView struct {
	Shop_ID   *int    `protobuf:"bytes,1,opt,name=Shop_ID,json=ShopID,proto3" json:"Shop_ID,omitempty"`
	Item_ID   *int    `protobuf:"bytes,2,opt,name=Item_ID,json=ItemID,proto3" json:"Item_ID,omitempty"`
	Option_ID *int    `protobuf:"bytes,3,opt,name=Option_ID,json=OptionID,proto3" json:"Option_ID,omitempty"`
	Name      *string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Price     *int    `protobuf:"bytes,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *ItemOptionView) GetShop_ID() *int {
	if m != nil {
		return m.Shop_ID
	}
	return nil
}

func (m *ItemOptionView) GetItem_ID() *int {
	if m != nil {
		return m.Item_ID
	}
	return nil
}

func (m *ItemOptionView) GetOption_ID() *int {
	if m != nil {
		return m.Option_ID
	}
	return nil
}

func (m *ItemOptionView) GetName() *string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ItemOptionView) GetPrice() *int {
	if m != nil {
		return m.Price
	}
	return nil
}
