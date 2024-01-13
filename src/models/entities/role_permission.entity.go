package entities

const TableNameRolePermission = "role_permission"

// RolePermission mapped from table <role_permission>
type RolePermission struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PermissionID int64  `gorm:"column:permissionId;not null" json:"permissionId"`
	RoleID       int64  `gorm:"column:roleId;not null" json:"roleId"`
	CreateAt     string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt     string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt     string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName RolePermission's table name
func (*RolePermission) TableName() string {
	return TableNameRolePermission
}
