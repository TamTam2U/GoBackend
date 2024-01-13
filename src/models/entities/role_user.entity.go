package entities

const TableNameRoleUser = "role_user"

// RoleUser mapped from table <role_user>
type RoleUser struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID   int64  `gorm:"column:userId;not null" json:"userId"`
	RoleID   int64  `gorm:"column:roleId;not null" json:"roleId"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName RoleUser's table name
func (*RoleUser) TableName() string {
	return TableNameRoleUser
}
