package model

//Ticket struct represents tickets table in database
type Content struct {
	ContentID       uint64 `gorm:"primary_key:auto_increment" json:"content_id"`
	ContentTitle    string `gorm:"type:text" json:"content_title"`
	ContentDetailID uint64 `gorm:"type:int" json:"content_detail_id"`
	Comment         string `gorm:"type:text" json:"comment"`
	IsUsable        string `gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser      uint64 `gorm:"type:int" json:"create_user"`
	CreateDate      string `gorm:"type:datetime" json:"create_date"`
	UpdateUser      uint64 `gorm:"type:int" json:"update_user"`
	UpdateDate      string `gorm:"type:datetime" json:"update_date"`
	// ContentDetails  ContentDetail `gorm:"foreignKey:content_detail_id;references:content_detail_id"`
	// ContentDetails   ContentDetail `gorm:"foreignKey:content_id; ASSOCIATION_FOREIGNKEY:content_detail_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"content_detail_id"`
}
