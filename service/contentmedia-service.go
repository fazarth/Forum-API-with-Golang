package service

import (
	"fmt"
	"log"

	"buddyku/model"
	"buddyku/repository"

	"github.com/mashingan/smapping"
)

//ContentMediaService is a ....
type ContentMediaService interface {
	Insert(b model.ContentMedia) model.ContentMedia
	Update(b model.ContentMedia) model.ContentMedia
	Delete(b model.ContentMedia)
	All() []model.ContentMedia
	FindByID(contentmediaID uint64) model.ContentMedia
	IsAllowedToEdit(userID string, contentmediaID uint64) bool
	ByTicketID(sTicketID string) []model.ContentMedia
	FilterCAPA(contentID string) []model.ContentMedia
	FilterY() []model.ContentMedia
}

type contentmediaService struct {
	contentmediaRepository repository.ContentMediaRepository
}

//NewContentMediaService .....
func NewContentMediaService(contentmediaRepo repository.ContentMediaRepository) ContentMediaService {
	return &contentmediaService{
		contentmediaRepository: contentmediaRepo,
	}
}

func (service *contentmediaService) Insert(b model.ContentMedia) model.ContentMedia {
	contentmedia := model.ContentMedia{}
	err := smapping.FillStruct(&contentmedia, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contentmediaRepository.InsertContentMedia(contentmedia)
	return res
}

func (service *contentmediaService) Update(b model.ContentMedia) model.ContentMedia {
	res := service.contentmediaRepository.UpdateContentMedia(b)
	return res
}

func (service *contentmediaService) Delete(b model.ContentMedia) {
	service.contentmediaRepository.DeleteContentMedia(b)
}

func (service *contentmediaService) All() []model.ContentMedia {
	return service.contentmediaRepository.AllContentMedia()
}

func (service *contentmediaService) FindByID(contentmediaID uint64) model.ContentMedia {
	return service.contentmediaRepository.FindContentMediaByID(contentmediaID)
}

func (service *contentmediaService) IsAllowedToEdit(userID string, contentmediaID uint64) bool {
	b := service.contentmediaRepository.FindContentMediaByID(contentmediaID)
	id := fmt.Sprintf("%v", b.CreateUser)
	return userID == id
}

func (service *contentmediaService) ByTicketID(sTicketID string) []model.ContentMedia {
	return service.contentmediaRepository.ByTicketID(sTicketID)
}

func (service *contentmediaService) FilterCAPA(contentID string) []model.ContentMedia {
	return service.contentmediaRepository.FilterCAPA(contentID)
}

func (service *contentmediaService) FilterY() []model.ContentMedia {
	return service.contentmediaRepository.FilterY()
}
