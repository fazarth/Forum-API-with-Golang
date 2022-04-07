package model

//Ticket struct represents tickets table in database
type Role struct {
	RoleID     uint64 `gorm:"primary_key:auto_increment" json:"role_id"`
	RoleName   string `gorm:"type:varchar(255)" json:"role_title"`
	Comment    string `gorm:"type:text" json:"comment"`
	IsUsable   string `gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser uint64 `gorm:"type:int" json:"create_user"`
	CreateDate string `gorm:"type:datetime" json:"create_date"`
	UpdateUser uint64 `gorm:"type:int" json:"update_user"`
	UpdateDate string `gorm:"type:datetime" json:"update_date"`
}
