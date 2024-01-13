package entities

const TableNameSchemaSubscription = "schema_subscription"

// SchemaSubscription mapped from table <schema_subscription>
type SchemaSubscription struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SchemaID       int64  `gorm:"column:schemaId;not null" json:"schemaId"`
	SubscriptionID int64  `gorm:"column:subscriptionId;not null" json:"subscriptionId"`
	Hostname       string `gorm:"column:hostname;not null" json:"hostname"`
	Username       string `gorm:"column:username;not null" json:"username"`
	Password       string `gorm:"column:password;not null" json:"password"`
	Database       string `gorm:"column:database;not null" json:"database"`
	Port           string `gorm:"column:port;not null" json:"port"`
	Engine         string `gorm:"column:engine;not null" json:"engine"`
	CreateAt       string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt       string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt       string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName SchemaSubscription's table name
func (*SchemaSubscription) TableName() string {
	return TableNameSchemaSubscription
}
