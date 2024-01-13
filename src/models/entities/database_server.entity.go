package entities

const TableNameDatabaseServer = "database_server"

// DatabaseServer mapped from table <database_server>
type DatabaseServer struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ServiceID int64  `gorm:"column:serviceId;not null" json:"serviceId"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Hostname  string `gorm:"column:hostname;not null" json:"hostname"`
	Username  string `gorm:"column:username;not null" json:"username"`
	Password  string `gorm:"column:password;not null" json:"password"`
	Port      string `gorm:"column:port;not null" json:"port"`
	Engine    string `gorm:"column:engine;not null" json:"engine"`
	Count     int32  `gorm:"column:count;not null" json:"count"`
	CreateAt  string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt  string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt  string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName DatabaseServer's table name
func (*DatabaseServer) TableName() string {
	return TableNameDatabaseServer
}
