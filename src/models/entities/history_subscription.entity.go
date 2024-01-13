package entities

import (
	"time"
)

const TableNameHistorySubscription = "history_subscription"

// HistorySubscription mapped from table <history_subscription>
type HistorySubscription struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SubscriptionID int64     `gorm:"column:subscriptionId;not null" json:"subscriptionId"`
	AgencyID       int64     `gorm:"column:agencyId;not null" json:"agencyId"`
	DetailAgency   string    `gorm:"column:detailAgency;not null" json:"detailAgency"`
	AppID          int64     `gorm:"column:appId;not null" json:"appId"`
	DetailApp      string    `gorm:"column:detailApp;not null" json:"detailApp"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Logo           string    `gorm:"column:logo;not null" json:"logo"`
	Price          string    `gorm:"column:price;not null" json:"price"`
	Setting        string    `gorm:"column:setting;not null" json:"setting"`
	StartDate      time.Time `gorm:"column:startDate" json:"startDate"`
	EndDate        time.Time `gorm:"column:endDate" json:"endDate"`
	CreateAt       string    `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt       string    `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt       string    `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName HistorySubscription's table name
func (*HistorySubscription) TableName() string {
	return TableNameHistorySubscription
}
