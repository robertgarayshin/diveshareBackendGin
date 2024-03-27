package handler

import (
	"diveshareBackendGin/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"net/http"
	"os"

	_ "diveshareBackendGin/docs"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	wd, _ := os.Getwd()
	router.StaticFS("/images/", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/catalog/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/about/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/blog/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/car/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/career/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/contact/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/description/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/faq/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/safety/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/share/images", http.Dir(wd+"/internal/static/images"))
	router.StaticFS("/upgrade/images", http.Dir(wd+"/internal/static/images"))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	user := router.Group("/user")
	{
		user.GET(":id", h.GetUserById)
		user.GET("", h.GetAllUsers)
		user.PUT(":id", h.EditProfile)
		user.DELETE(":id", h.DeleteProfile)
	}
	car := router.Group("/car")
	{
		car.POST("/new", h.NewCar)
		car.GET("/", h.GetAllCars)
		car.GET("/:id", h.GetCarById)
		car.PUT("/:id", h.EditCar)
		car.DELETE("/:id", h.DeleteCar)
	}
	rent := router.Group("/rent")
	{
		rent.POST("/new", h.NewRent)
		rent.GET("/", h.GetAllRents)
		rent.GET("/:id", h.GetRentById)
		rent.PUT("/:id", h.EditRent)
		rent.DELETE("/:id", h.DeleteRent)
	}
	review := router.Group("/review")
	{
		review.POST("/new", h.NewReview)
		review.GET("/", h.GetAllReviews)
		review.GET("/:id", h.GetReviewById)
		review.PUT("/:id", h.EditReview)
		review.DELETE("/:id", h.DeleteReview)
	}
	router.GET("/catalog", h.staticCatalog)
	router.GET("/auto", h.staticCar)
	router.GET("/about", h.staticAbout)
	router.GET("/blog", h.staticBlog)
	router.GET("/career", h.staticCareer)
	router.GET("/contact", h.staticContact)
	router.GET("/description", h.staticDescription)
	router.GET("/faq", h.staticFaq)
	router.GET("/safety", h.staticSafety)
	router.GET("/share", h.staticShare)
	router.GET("/upgrade", h.staticUpgrade)
	router.GET("/", h.staticMain)

	return router
}
