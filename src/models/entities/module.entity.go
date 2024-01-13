package entities

const TableNameModule = "module"

// Module mapped from table <module>
type Module struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID    int64  `gorm:"column:appId" json:"appId"`
	Name     string `gorm:"column:name;not null" json:"name"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Module's table name
func (*Module) TableName() string {
	return TableNameModule
}
