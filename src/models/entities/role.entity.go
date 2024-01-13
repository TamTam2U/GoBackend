package entities

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AgencyID int64  `gorm:"column:agencyId;not null" json:"agencyId"`
	Role     string `gorm:"column:role;not null" json:"role"`
	Status   string `gorm:"column:status;not null" json:"status"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
