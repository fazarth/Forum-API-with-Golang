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

//ContentDetailController is a ...
type ContentDetailController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type contentdetailController struct {
	contentdetailService service.ContentDetailService
	jwtService           service.JWTService
}

//NewContentDetailController create a new instances of BoookController
func NewContentDetailController(contentdetailServ service.ContentDetailService, jwtServ service.JWTService) ContentDetailController {
	return &contentdetailController{
		contentdetailService: contentdetailServ,
		jwtService:           jwtServ,
	}
}

func (c *contentdetailController) All(context *gin.Context) {
	var contentdetails []model.ContentDetail = c.contentdetailService.All()
	res := helper.BuildResponse(true, "OK", contentdetails)
	context.JSON(http.StatusOK, res)
}

func (c *contentdetailController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var contentdetail model.ContentDetail = c.contentdetailService.FindByID(id)
	if (contentdetail == model.ContentDetail{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", contentdetail)
		context.JSON(http.StatusOK, res)
	}
}

func (c *contentdetailController) Insert(context *gin.Context) {
	var contentdetailCreateDTO model.ContentDetail
	errDTO := context.ShouldBind(&contentdetailCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			contentdetailCreateDTO.CreateUser = convertedUserID
		}
		result := c.contentdetailService.Insert(contentdetailCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *contentdetailController) Update(context *gin.Context) {
	var contentdetailUpdateDTO model.ContentDetail
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	errDTO := context.ShouldBind(&contentdetailUpdateDTO)
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
	contentdetailUpdateDTO.ContentDetailID = id
	contentdetailUpdateDTO.UpdateUser = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.contentdetailService.Update(contentdetailUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *contentdetailController) Delete(context *gin.Context) {
	var contentdetail model.ContentDetail
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	contentdetail.ContentDetailID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.contentdetailService.IsAllowedToEdit(userID, contentdetail.ContentDetailID) {
		c.contentdetailService.Delete(contentdetail)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *contentdetailController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
