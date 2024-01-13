package entities

import (
	"time"
)

const TableNameSubscription = "subscription"

// Subscription mapped from table <subscription>
type Subscription struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AgencyID  int64     `gorm:"column:agencyId;not null" json:"agencyId"`
	AppID     int64     `gorm:"column:appId;not null" json:"appId"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Logo      string    `gorm:"column:logo;not null" json:"logo"`
	Price     string    `gorm:"column:price;not null" json:"price"`
	Setting   string    `gorm:"column:setting;not null" json:"setting"`
	StartDate time.Time `gorm:"column:startDate" json:"startDate"`
	EndDate   time.Time `gorm:"column:endDate" json:"endDate"`
	Status    int32     `gorm:"column:status;not null" json:"status"`
	CreateAt  string    `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt  string    `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt  string    `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Subscription's table name
func (*Subscription) TableName() string {
	return TableNameSubscription
}
