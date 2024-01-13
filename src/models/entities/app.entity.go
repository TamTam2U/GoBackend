package entities

const TableNameApp = "app"

// App mapped from table <app>
type App struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Type     int32  `gorm:"column:type;not null" json:"type"`
	Price    string `gorm:"column:price;not null" json:"price"`
	Image    string `gorm:"column:image;not null" json:"image"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}
