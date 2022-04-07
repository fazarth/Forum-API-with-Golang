package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"buddyku/helper"
	"buddyku/model"
	"buddyku/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//ContentMediaController is a ...
type ContentMediaController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	UploadFile(context *gin.Context)
}

type contentmediaController struct {
	contentmediaService service.ContentMediaService
	jwtService          service.JWTService
}

//NewContentMediaController create a new instances of BoookController
func NewContentMediaController(contentmediaServ service.ContentMediaService, jwtServ service.JWTService) ContentMediaController {
	return &contentmediaController{
		contentmediaService: contentmediaServ,
		jwtService:          jwtServ,
	}
}

func (c *contentmediaController) All(context *gin.Context) {
	var contentmedias []model.ContentMedia = c.contentmediaService.All()
	res := helper.BuildResponse(true, "OK", contentmedias)
	context.JSON(http.StatusOK, res)
}

func (c *contentmediaController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var contentmedia model.ContentMedia = c.contentmediaService.FindByID(id)
	if (contentmedia == model.ContentMedia{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", contentmedia)
		context.JSON(http.StatusOK, res)
	}
}

func (c *contentmediaController) Insert(context *gin.Context) {
	var contentmediaCreateDTO model.ContentMedia
	errDTO := context.ShouldBind(&contentmediaCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			contentmediaCreateDTO.CreateUser = convertedUserID
		}
		result := c.contentmediaService.Insert(contentmediaCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *contentmediaController) UploadFile(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("Get GORM err: %s", err.Error()))
		fmt.Println(err)
		return
	}

	files := form.File["files"]

	for _, file := range files {
		path := "./contentmedia/" + file.Filename
		if err := context.SaveUploadedFile(file, path); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	// Response
	context.String(http.StatusOK, fmt.Sprintf("File %d uploaded successfully", len(files)))
}

func (c *contentmediaController) Update(context *gin.Context) {
	var contentmediaUpdateDTO model.ContentMedia
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	errDTO := context.ShouldBind(&contentmediaUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	idUser, errID := strconv.ParseUint(userID, 10, 64)
	contentmediaUpdateDTO.Media_Id = id
	contentmediaUpdateDTO.UpdateUser = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.contentmediaService.Update(contentmediaUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *contentmediaController) Delete(context *gin.Context) {
	var contentmedia model.ContentMedia
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	contentmedia.Media_Id = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.contentmediaService.IsAllowedToEdit(userID, contentmedia.Media_Id) {
		c.contentmediaService.Delete(contentmedia)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *contentmediaController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *contentmediaController) ByTicketID(context *gin.Context) {
	contentID := context.Query("content_id")
	// contentID, _ := strconv.ParseUint(context.Param("content_id"), 0, 0)
	var contentmedias []model.ContentMedia = c.contentmediaService.ByTicketID(contentID)
	res := helper.BuildResponse(true, "OK", contentmedias)
	context.JSON(http.StatusOK, res)
}

func (c *contentmediaController) FilterCAPA(context *gin.Context) {
	contentID := context.Query("content_id")
	var contentmedias []model.ContentMedia = c.contentmediaService.FilterCAPA(contentID)
	res := helper.BuildResponse(true, "OK", contentmedias)
	context.JSON(http.StatusOK, res)
}

func (c *contentmediaController) FilterY(context *gin.Context) {
	var contentmedias []model.ContentMedia = c.contentmediaService.FilterY()
	res := helper.BuildResponse(true, "OK", contentmedias)
	context.JSON(http.StatusOK, res)
}
