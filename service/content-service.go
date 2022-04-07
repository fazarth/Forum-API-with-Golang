package service

import (
	"fmt"
	"log"

	"buddyku/model"
	"buddyku/repository"

	"github.com/mashingan/smapping"
)

//ContentService is a ....
type ContentService interface {
	Insert(b model.Content) model.Content
	Update(b model.Content) model.Content
	Delete(b model.Content)
	All() []model.Content
	FindByID(contentID uint64) model.Content
	IsAllowedToEdit(userID string, contentID uint64) bool
}

type contentService struct {
	contentRepository repository.ContentRepository
}

//NewContentService .....
func NewContentService(contentRepo repository.ContentRepository) ContentService {
	return &contentService{
		contentRepository: contentRepo,
	}
}

func (service *contentService) Insert(b model.Content) model.Content {
	content := model.Content{}
	err := smapping.FillStruct(&content, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contentRepository.InsertContent(content)
	return res
}

func (service *contentService) Update(b model.Content) model.Content {
	res := service.contentRepository.UpdateContent(b)
	return res
}

func (service *contentService) Delete(b model.Content) {
	service.contentRepository.DeleteContent(b)
}

func (service *contentService) All() []model.Content {
	return service.contentRepository.AllContent()
}

func (service *contentService) FindByID(contentID uint64) model.Content {
	return service.contentRepository.FindContentByID(contentID)
}

func (service *contentService) IsAllowedToEdit(userID string, contentID uint64) bool {
	b := service.contentRepository.FindContentByID(contentID)
	id := fmt.Sprintf("%v", b.CreateUser)
	return userID == id
}
