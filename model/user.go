package model

//User represents users table in database
type User struct {
	UserID     uint64 `gorm:"primary_key:auto_increment" json:"user_id"`
	Role       uint64 `gorm:"type:varchar(255)" json:"role"`
	UserName   string `gorm:"uniqueIndex;type:varchar(255)" json:"user_name"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Token      string `gorm:"-" json:"token,omitempty"`
	Comment    string `gorm:"type:text" json:"comment"`
	IsUsable   string `gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser uint64 `gorm:"type:int" json:"create_user"`
	CreateDate string `gorm:"type:datetime" json:"create_date"`
	UpdateUser uint64 `gorm:"type:int" json:"update_user"`
	UpdateDate string `gorm:"type:datetime" json:"update_date"`
}
