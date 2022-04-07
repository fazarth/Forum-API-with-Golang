package main

import (
	"buddyku/config"
	"buddyku/controller"
	"buddyku/middleware"
	"buddyku/repository"
	"buddyku/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	db                      *gorm.DB                           = config.SetupDatabaseConnection()
	userRepository          repository.UserRepository          = repository.NewUserRepository(db)
	contentRepository       repository.ContentRepository       = repository.NewContentRepository(db)
	contentdetailRepository repository.ContentDetailRepository = repository.NewContentDetailRepository(db)
	contentmediaRepository  repository.ContentMediaRepository  = repository.NewContentMediaRepository(db)
	jwtService              service.JWTService                 = service.NewJWTService()
	userService             service.UserService                = service.NewUserService(userRepository)
	contentService          service.ContentService             = service.NewContentService(contentRepository)
	contentdetailService    service.ContentDetailService       = service.NewContentDetailService(contentdetailRepository)
	contentmediaService     service.ContentMediaService        = service.NewContentMediaService(contentmediaRepository)
	authService             service.AuthService                = service.NewAuthService(userRepository)
	authController          controller.AuthController          = controller.NewAuthController(authService, jwtService)
	userController          controller.UserController          = controller.NewUserController(userService, jwtService)
	contentController       controller.ContentController       = controller.NewContentController(contentService, jwtService)
	contentmediaController  controller.ContentMediaController  = controller.NewContentMediaController(contentmediaService, jwtService)
	contentdetailController controller.ContentDetailController = controller.NewContentDetailController(contentdetailService, jwtService)
)

//Kode untuk Swagger
// @title Example Employee API
// @version 1.0
// @description This is a Example Employee API
// @termsOfService http://swagger.io/terms/
// contact.name API Support
// contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Use(middleware.SetupCorsMiddleware())
	r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	contentRoutes := r.Group("api/contents", middleware.AuthorizeJWT(jwtService))
	{
		contentRoutes.GET("/", contentController.All)
		contentRoutes.POST("/", contentController.Insert)
		contentRoutes.GET("/:id", contentController.FindByID)
		contentRoutes.PUT("/:id", contentController.Update)
		contentRoutes.DELETE("/:id", contentController.Delete)
	}

	contentdetailRoutes := r.Group("api/contentdetails", middleware.AuthorizeJWT(jwtService))
	{
		contentdetailRoutes.GET("/", contentdetailController.All)
		contentdetailRoutes.POST("/", contentdetailController.Insert)
		contentdetailRoutes.GET("/:id", contentdetailController.FindByID)
		contentdetailRoutes.PUT("/:id", contentdetailController.Update)
		contentdetailRoutes.DELETE("/:id", contentdetailController.Delete)
	}

	contentmediaRoutes := r.Group("api/contentmedia", middleware.AuthorizeJWT(jwtService))
	{
		contentmediaRoutes.GET("/", contentmediaController.All)
		contentmediaRoutes.POST("/", contentmediaController.Insert)
		contentmediaRoutes.GET("/:id", contentmediaController.FindByID)
		contentmediaRoutes.PUT("/:id", contentmediaController.Update)
		contentmediaRoutes.DELETE("/:id", contentmediaController.Delete)
		contentmediaRoutes.POST("/uploadFile", contentmediaController.UploadFile)
	}

	r.StaticFS("/upload", http.Dir("upload"))

	r.Run(":3000")
}
