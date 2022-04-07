package repository

import (
	"buddyku/model"

	"gorm.io/gorm"
)

//ContentMediaRepository is a ....
type ContentMediaRepository interface {
	InsertContentMedia(b model.ContentMedia) model.ContentMedia
	UpdateContentMedia(b model.ContentMedia) model.ContentMedia
	DeleteContentMedia(b model.ContentMedia)
	AllContentMedia() []model.ContentMedia
	FindContentMediaByID(Media_Id uint64) model.ContentMedia
	ByTicketID(contentID string) []model.ContentMedia
	FilterCAPA(contentID string) []model.ContentMedia
	FilterY() []model.ContentMedia
}

type contentmediaConnection struct {
	connection *gorm.DB
}

//NewContentMediaRepository creates an instance ContentMediaRepository
func NewContentMediaRepository(dbConn *gorm.DB) ContentMediaRepository {
	return &contentmediaConnection{
		connection: dbConn,
	}
}

func (db *contentmediaConnection) InsertContentMedia(b model.ContentMedia) model.ContentMedia {
	db.connection.Save(&b)
	db.connection.Preload("ContentMedia").Find(&b)
	return b
}

func (db *contentmediaConnection) UpdateContentMedia(b model.ContentMedia) model.ContentMedia {
	var contentmedia model.ContentMedia
	db.connection.Find(&contentmedia).Where("media_id = ?", b.Media_Id)
	contentmedia.Comment = b.Comment
	db.connection.Updates(&b)
	return b
}

func (db *contentmediaConnection) DeleteContentMedia(b model.ContentMedia) {
	db.connection.Delete(&b)
}

func (db *contentmediaConnection) FindContentMediaByID(Media_Id uint64) model.ContentMedia {
	var contentmedia model.ContentMedia
	db.connection.Preload("ContentMedia").Find(&contentmedia, Media_Id)
	// db.connection.Raw("SELECT content_media.media_id, content_media.media_name, content_media.media_file, content_media.media_size, content_media.comment, content_media.is_usable,content_media.create_user, content_media.create_date, content_media.update_user, content_media.update_date, contents.content_id, contents.content_type, contents.no_content, contents.type_ncr, contents.case, contents.description, contents.send_to, contents.forward_default, contents.verification, contents.send_ncr, contents.verification_date, contents.corrective_action, contents.preventive_action, contents.capa_user, contents.capa_date, contents.validation_date, contents.status, contents.confirm, contents.verification_user FROM content_media JOIN contents ON content_media.content_id = contents.content_id WHERE content_media.media_id = ? ", Media_Id).Scan(&contentmedia)
	return contentmedia
}

func (db *contentmediaConnection) AllContentMedia() []model.ContentMedia {
	var contentmedia []model.ContentMedia
	db.connection.Preload("ContentMedia").Find(&contentmedia)
	// db.connection.Raw("SELECT content_media.media_id, content_media.media_name, content_media.media_file, content_media.media_size, content_media.comment, content_media.is_usable,content_media.create_user, content_media.create_date, content_media.update_user, content_media.update_date, contents.content_id, contents.content_type, contents.no_content, contents.type_ncr, contents.case, contents.description, contents.send_to, contents.forward_default, contents.verification, contents.send_ncr, contents.verification_date, contents.corrective_action, contents.preventive_action, contents.capa_user, contents.capa_date, contents.validation_date, contents.status, contents.confirm, contents.verification_user FROM content_media JOIN contents ON content_media.content_id = contents.content_id").Scan(&contentmedia)
	return contentmedia
}

func (db *contentmediaConnection) ByTicketID(contentID string) []model.ContentMedia {
	var contentmedia []model.ContentMedia
	db.connection.Raw("SELECT content_media.media_id, content_media.media_name, content_media.media_file, content_media.media_size, content_media.comment, content_media.is_usable,content_media.create_user, content_media.create_date, content_media.update_user, content_media.update_date, contents.content_id, contents.content_type, contents.no_content, contents.type_ncr, contents.case, contents.description, contents.send_to, contents.forward_default, contents.verification, contents.send_ncr, contents.verification_date, contents.corrective_action, contents.preventive_action, contents.capa_user, contents.capa_date, contents.validation_date, contents.status, contents.confirm, contents.verification_user FROM content_media JOIN contents ON content_media.content_id = contents.content_id WHERE content_media.content_id = ? AND content_media.is_usable = 'Y' ", contentID).Scan(&contentmedia)
	return contentmedia
}

func (db *contentmediaConnection) FilterCAPA(contentID string) []model.ContentMedia {
	var contentmedia []model.ContentMedia
	db.connection.Raw("SELECT content_media.media_id, content_media.media_name, content_media.media_file, content_media.media_size, content_media.comment, content_media.is_usable,content_media.create_user, content_media.create_date, content_media.update_user, content_media.update_date, contents.content_id, contents.content_type, contents.no_content, contents.type_ncr, contents.case, contents.description, contents.send_to, contents.forward_default, contents.verification, contents.send_ncr, contents.verification_date, contents.corrective_action, contents.preventive_action, contents.capa_user, contents.capa_date, contents.validation_date, contents.status, contents.confirm, contents.verification_user FROM content_media JOIN contents ON content_media.content_id = contents.content_id WHERE content_media.content_id = ? AND content_media.is_usable = 'C' ", contentID).Scan(&contentmedia)
	return contentmedia
}

func (db *contentmediaConnection) FilterY() []model.ContentMedia {
	var contentmedia []model.ContentMedia
	db.connection.Raw("SELECT content_media.media_id, content_media.media_name, content_media.media_file, content_media.media_size, content_media.comment, content_media.is_usable,content_media.create_user, content_media.create_date, content_media.update_user, content_media.update_date, contents.content_id, contents.content_type, contents.no_content, contents.type_ncr, contents.case, contents.description, contents.send_to, contents.forward_default, contents.verification, contents.send_ncr, contents.verification_date, contents.corrective_action, contents.preventive_action, contents.capa_user, contents.capa_date, contents.validation_date, contents.status, contents.confirm, contents.verification_user FROM content_media JOIN contents ON content_media.content_id = contents.content_id WHERE content_media.is_usable = 'Y' ").Scan(&contentmedia)
	return contentmedia
}
