package repository

import (
	"buddyku/model"

	"gorm.io/gorm"
)

//ContentRepository is a ....
type ContentRepository interface {
	InsertContent(b model.Content) model.Content
	UpdateContent(b model.Content) model.Content
	DeleteContent(b model.Content)
	AllContent() []model.Content
	FindContentByID(contentID uint64) model.Content
}

type contentConnection struct {
	connection *gorm.DB
}

//NewContentRepository creates an instance ContentRepository
func NewContentRepository(dbConn *gorm.DB) ContentRepository {
	return &contentConnection{
		connection: dbConn,
	}
}

func (db *contentConnection) InsertContent(b model.Content) model.Content {
	db.connection.Save(&b)
	db.connection.Preload("Contents").Find(&b)
	return b
}

func (db *contentConnection) UpdateContent(b model.Content) model.Content {
	db.connection.Save(&b)
	db.connection.Preload("Contents").Find(&b)
	return b
}

func (db *contentConnection) DeleteContent(b model.Content) {
	db.connection.Delete(&b)
}

func (db *contentConnection) FindContentByID(contentID uint64) model.Content {
	var content model.Content
	db.connection.Where("content_id = ?", contentID).Find(&content)
	return content
}

func (db *contentConnection) AllContent() []model.Content {
	var content []model.Content
	db.connection.Preload("content_details").Find(&content)
	// db.connection.Joins("JOIN content_details ON contents.content_detail_id = content_details.content_detail_id").Find(&content)
	// db.connection.Joins("content_details").Find(&content)
	// db.connection.Model(&content).Joins("JOIN content_details ON contents.content_detail_id = content_details.content_detail_id").Scan(&content)
	return content
}
