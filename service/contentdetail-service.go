package service

import (
	"fmt"
	"log"

	"buddyku/model"
	"buddyku/repository"

	"github.com/mashingan/smapping"
)

//ContentDetailService is a ....
type ContentDetailService interface {
	Insert(b model.ContentDetail) model.ContentDetail
	Update(b model.ContentDetail) model.ContentDetail
	Delete(b model.ContentDetail)
	All() []model.ContentDetail
	FindByID(contentdetailID uint64) model.ContentDetail
	IsAllowedToEdit(userID string, contentdetailID uint64) bool
}

type contentdetailService struct {
	contentdetailRepository repository.ContentDetailRepository
}

//NewContentDetailService .....
func NewContentDetailService(contentdetailRepo repository.ContentDetailRepository) ContentDetailService {
	return &contentdetailService{
		contentdetailRepository: contentdetailRepo,
	}
}

func (service *contentdetailService) Insert(b model.ContentDetail) model.ContentDetail {
	contentdetail := model.ContentDetail{}
	err := smapping.FillStruct(&contentdetail, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contentdetailRepository.InsertContentDetail(contentdetail)
	return res
}

func (service *contentdetailService) Update(b model.ContentDetail) model.ContentDetail {
	res := service.contentdetailRepository.UpdateContentDetail(b)
	return res
}

func (service *contentdetailService) Delete(b model.ContentDetail) {
	service.contentdetailRepository.DeleteContentDetail(b)
}

func (service *contentdetailService) All() []model.ContentDetail {
	return service.contentdetailRepository.AllContentDetail()
}

func (service *contentdetailService) FindByID(contentdetailID uint64) model.ContentDetail {
	return service.contentdetailRepository.FindContentDetailByID(contentdetailID)
}

func (service *contentdetailService) IsAllowedToEdit(userID string, contentdetailID uint64) bool {
	b := service.contentdetailRepository.FindContentDetailByID(contentdetailID)
	id := fmt.Sprintf("%v", b.CreateUser)
	return userID == id
}
