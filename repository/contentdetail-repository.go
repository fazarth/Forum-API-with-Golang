package repository

import (
	"buddyku/model"

	"gorm.io/gorm"
)

//ContentDetailRepository is a ....
type ContentDetailRepository interface {
	InsertContentDetail(b model.ContentDetail) model.ContentDetail
	UpdateContentDetail(b model.ContentDetail) model.ContentDetail
	DeleteContentDetail(b model.ContentDetail)
	AllContentDetail() []model.ContentDetail
	FindContentDetailByID(contentdetailID uint64) model.ContentDetail
}

type contentdetailConnection struct {
	connection *gorm.DB
}

//NewContentDetailRepository creates an instance ContentDetailRepository
func NewContentDetailRepository(dbConn *gorm.DB) ContentDetailRepository {
	return &contentdetailConnection{
		connection: dbConn,
	}
}

func (db *contentdetailConnection) InsertContentDetail(b model.ContentDetail) model.ContentDetail {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *contentdetailConnection) UpdateContentDetail(b model.ContentDetail) model.ContentDetail {
	var contentdetail model.ContentDetail
	db.connection.Find(&contentdetail).Where("content_detail_id = ?", b.ContentDetailID)
	contentdetail.ContentDetailID = b.ContentDetailID
	db.connection.Updates(&b)
	return b
}

func (db *contentdetailConnection) DeleteContentDetail(b model.ContentDetail) {
	db.connection.Delete(&b)
}

func (db *contentdetailConnection) FindContentDetailByID(contentdetailID uint64) model.ContentDetail {
	var contentdetail model.ContentDetail
	db.connection.Where("content_detail_id = ?", contentdetailID).Joins("contents").Find(&contentdetail)


	return contentdetail
}

func (db *contentdetailConnection) AllContentDetail() []model.ContentDetail {
	var contentdetails []model.ContentDetail
	db.connection.Joins("contents").Find(&contentdetails)
	return contentdetails
}
