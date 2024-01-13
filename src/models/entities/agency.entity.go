package entities

const TableNameAgency = "agency"

// Agency mapped from table <agency>
type Agency struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;not null" json:"email"`
	NoHp     string `gorm:"column:noHp;not null" json:"noHp"`
	Address  string `gorm:"column:address;not null" json:"address"`
	Logo     string `gorm:"column:logo" json:"logo"`
	Domain   string `gorm:"column:domain;not null" json:"domain"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Agency's table name
func (*Agency) TableName() string {
	return TableNameAgency
}
