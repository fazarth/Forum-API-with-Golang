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

//ContentController is a ...
type ContentController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type contentController struct {
	contentService service.ContentService
	jwtService     service.JWTService
}

//NewContentController create a new instances of BoookController
func NewContentController(contentServ service.ContentService, jwtServ service.JWTService) ContentController {
	return &contentController{
		contentService: contentServ,
		jwtService:     jwtServ,
	}
}

// Select All Content
// @Description Select All Content
// @Summary Select All Content
// @Tags Content
// @Param Content body model.Content true "Select All Content Data"
// @Produce json
// @Success 200 {object} model.Content
// @Router /auth/register [get]

func (c *contentController) All(context *gin.Context) {
	var contents []model.Content = c.contentService.All()
	res := helper.BuildResponse(true, "OK", contents)
	context.JSON(http.StatusOK, res)
}

// Select Content By ID
// @Description Select Content By ID
// @Summary Select Content By ID
// @Tags Content
// @Param Content body model.Content true "Select Content By ID Data"
// @Produce json
// @Success 200 {object} model.Content
// @Router /auth/register [get]

func (c *contentController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var content model.Content = c.contentService.FindByID(id)
	if (content == model.Content{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", content)
		context.JSON(http.StatusOK, res)
	}
}

// Insert Content
// @Description Insert Content
// @Summary Insert Content
// @Tags Content
// @Param Content body model.Content true "Insert Content Data"
// @Produce json
// @Success 200 {object} model.Content
// @Router /auth/register [post]

func (c *contentController) Insert(context *gin.Context) {
	var contentCreateDTO model.Content
	errDTO := context.ShouldBind(&contentCreateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Do not leave the field blank, please fill it with 0 or -", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			contentCreateDTO.CreateUser = convertedUserID
			contentCreateDTO.UpdateUser = convertedUserID
		}
		result := c.contentService.Insert(contentCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// Update Content
// @Description Update Content
// @Summary Update Content
// @Tags Content
// @Param Content body model.Content true "Update Content Data"
// @Produce json
// @Success 200 {object} model.Content
// @Router /auth/register [put]

func (c *contentController) Update(context *gin.Context) {
	var contentUpdateDTO model.Content
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No Param ID Was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	errDTO := context.ShouldBind(&contentUpdateDTO)
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
	contentUpdateDTO.ContentID = id
	contentUpdateDTO.UpdateUser = idUser
	if errID != nil {
		response := helper.BuildErrorResponse("User ID Not Found", "User ID Not Found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.contentService.Update(contentUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// Delete Content
// @Description Delete Content
// @Summary Delete Content
// @Tags Content
// @Param Content body model.Content true "Delete Content Data"
// @Produce json
// @Success 200 {object} model.Content
// @Router /auth/register [delete]

func (c *contentController) Delete(context *gin.Context) {
	var content model.Content
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	content.ContentID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.contentService.IsAllowedToEdit(userID, content.ContentID) {
		c.contentService.Delete(content)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *contentController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *contentController) UploadFile(context *gin.Context) {

	form, err := context.MultipartForm()
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("Get GORM err: %s", err.Error()))
		fmt.Println(err)
		return
	}
	files := form.File["files"]

	for _, file := range files {
		path := "./upload/" + file.Filename
		if err := context.SaveUploadedFile(file, path); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("Upload File err: %s", err.Error()))
			return
		}
	}

	// Response
	context.String(http.StatusOK, fmt.Sprintf("File %d uploaded successfully", len(files)))
}
