package model

type ContentMedia struct {
	Media_Id     	uint64  `gorm:"primary_key:auto_increment" json:"media_id"`
	Media_Name  	string  `gorm:"type:varchar(255)" json:"media_name"`
	Media_File  	string  `gorm:"type:varchar(255)" json:"media_file"`
	Media_Size  	uint64  `gorm:"type:int" json:"media_size"`
	Comment    		string 	`gorm:"type:text" json:"comment"`
	IsUsable   		string 	`gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser 		uint64 	`gorm:"type:int" json:"create_user"`
	CreateDate 		string 	`gorm:"type:datetime" json:"create_date"`
	UpdateUser 		uint64 	`gorm:"type:int" json:"update_user"`
	UpdateDate 		string 	`gorm:"type:datetime" json:"update_date"`
	ContentDetailID uint64  `gorm:"type:int" json:"contentdetail_id"`
	// Content     Content `gorm:"foreignKey:ContentID; ASSOCIATION_FOREIGNKEY:ContentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ticket"`
}
