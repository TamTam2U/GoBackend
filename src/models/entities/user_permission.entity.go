package entities

const TableNameUserPermission = "user_permission"

// UserPermission mapped from table <user_permission>
type UserPermission struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       int64  `gorm:"column:userId;not null" json:"userId"`
	PermissionID int64  `gorm:"column:permissionId;not null" json:"permissionId"`
	CreateAt     string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt     string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt     string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName UserPermission's table name
func (*UserPermission) TableName() string {
	return TableNameUserPermission
}
