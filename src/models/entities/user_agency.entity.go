package entities

const TableNameUserAgency = "user_agency"

// UserAgency mapped from table <user_agency>
type UserAgency struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID   int64  `gorm:"column:userId;not null" json:"userId"`
	AgencyID int64  `gorm:"column:agencyId;not null" json:"agencyId"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName UserAgency's table name
func (*UserAgency) TableName() string {
	return TableNameUserAgency
}
