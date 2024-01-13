package entities

const TableNameSession = "session"

// Session mapped from table <session>
type Session struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID   int64  `gorm:"column:userId" json:"userId"`
	IP       string `gorm:"column:ip;not null" json:"ip"`
	Agent    string `gorm:"column:agent;not null" json:"agent"`
	Try      int32  `gorm:"column:try;not null" json:"try"`
	LastTry  string `gorm:"column:lastTry" json:"lastTry"`
	CreateAt string `gorm:"column:createAt;not null" json:"createAt"`
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
	DeleteAt string `gorm:"column:deleteAt" json:"deleteAt"`
}

// TableName Session's table name
func (*Session) TableName() string {
	return TableNameSession
}
