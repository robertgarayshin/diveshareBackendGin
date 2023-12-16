package handler

import (
	"diveshareBackendGin/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

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

	return router
}
