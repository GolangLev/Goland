package handler

import (
	"github.com/GolangLev/Goland/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes /*Инициализация маршрутов работы приложения*/
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.CreateList)
			list.GET("/", h.GetAllLists)
			list.GET("/:id", h.GetListsById)
			list.PUT("/:id", h.UpdateList)
			list.DELETE("/:id", h.DeleteList)

			items := list.Group(":id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetItemsById)
				items.GET("/:item_id", h.GetItemsById)
				items.PUT("/:item_id", h.UpdateItems)
				items.DELETE("/:item_id", h.DeleteItems)
			}
		}
	}

	return router
}
