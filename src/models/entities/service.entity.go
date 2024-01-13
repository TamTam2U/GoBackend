package entities

const TableNameService = "service"

// Service mapped from table <service>
type Service struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Service's table name
func (*Service) TableName() string {
	return TableNameService
}
