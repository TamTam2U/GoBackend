package entities

const TableNameSetting = "setting"

// Setting mapped from table <setting>
type Setting struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID    int64  `gorm:"column:appId;not null" json:"appId"`
	Key      string `gorm:"column:key" json:"key"`
	Value    string `gorm:"column:value" json:"value"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}
