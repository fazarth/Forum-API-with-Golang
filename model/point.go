package model

//User struct represents user table in database
type Point struct {
	PointID    uint64 `gorm:"primary_key:auto_increment" json:"point_id"`
	UserID     uint64 `gorm:"type:int" json:"user_id"`
	Comment    string `gorm:"type:text" json:"comment"`
	IsUsable   string `gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser uint64 `gorm:"type:int" json:"create_user"`
	CreateDate string `gorm:"type:datetime" json:"create_date"`
	UpdateUser uint64 `gorm:"type:int" json:"update_user"`
	UpdateDate string `gorm:"type:datetime" json:"update_date"`
	// Users      User   `gorm:"foreignKey:UserID; ASSOCIATION_FOREIGNKEY:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
