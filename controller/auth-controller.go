package controller

import (
	"net/http"
	"strconv"

	"buddyku/helper"
	"buddyku/model"
	"buddyku/service"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

type CredentialsLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login User
// @Description Login User
// @Summary Login User
// @Tags User
// @Param user body CredentialsLogin true "Input username & password"
// @Produce json
// @Success 200 {object} global.User
// @Router /auth/login [post]

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO CredentialsLogin
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.UserID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your Username or Password", "Invalid Username or Password", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register User
// @Description Register User
// @Summary Register User
// @Tags User
// @Param user body model.User true "Register User Data"
// @Produce json
// @Success 200 {object} model.User
// @Router /auth/register [post]

func (c *authController) Register(ctx *gin.Context) {
	var register model.User
	errDTO := ctx.ShouldBind(&register)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateUserName(register.UserName) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Username", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(register)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.UserID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
