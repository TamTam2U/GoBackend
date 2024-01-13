package entities

const TableNamePermission = "permission"

// Permission mapped from table <permission>
type Permission struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	ModuleID    int64  `gorm:"column:moduleId;not null" json:"moduleId"`
	CreateAt    string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt    string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt    string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Permission's table name
func (*Permission) TableName() string {
	return TableNamePermission
}
