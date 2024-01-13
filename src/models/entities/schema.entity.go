package entities

const TableNameSchema = "schema"

// Schema mapped from table <schema>
type Schema struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID     int64  `gorm:"column:appId;not null" json:"appId"`
	ServiceID int64  `gorm:"column:serviceId;not null" json:"serviceId"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Path      string `gorm:"column:path;not null" json:"path"`
	CreateAt  string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt  string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt  string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Schema's table name
func (*Schema) TableName() string {
	return TableNameSchema
}
