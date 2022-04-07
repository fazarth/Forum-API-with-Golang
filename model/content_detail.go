package model

//Common struct represents commons table in database
type ContentDetail struct {
	ContentDetailID uint64  `gorm:"primary_key:auto_increment" json:"content_detail_id"`
	Content         string  `gorm:"type:text" json:"no_content"`
	Comment    		string 	`gorm:"type:text" json:"comment"`
	IsUsable   		string 	`gorm:"type:varchar(255)" json:"is_usable"`
	CreateUser 		uint64 	`gorm:"type:int" json:"create_user"`
	CreateDate 		string 	`gorm:"type:datetime" json:"create_date"`
	UpdateUser 		uint64 	`gorm:"type:int" json:"update_user"`
	UpdateDate 		string 	`gorm:"type:datetime" json:"update_date"`
	// ContentID    	uint64 	`gorm:"type:int" json:"content_id"`
}
