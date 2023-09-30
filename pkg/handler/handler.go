package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/zhashkevych/todo-app/docs"
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

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/id", h.getListById)
			lists.PUT("/id", h.updateList)
			lists.DELETE("/id", h.deleteList)

			items := lists.Group("id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/id", h.getItemById)
			items.PUT("/id", h.updateItem)
			items.DELETE("/id", h.deleteItem)
		}
	}

	apiv1 := router.Group("api/v1")
	book := apiv1.Group("/book")
	{
		book.POST("/", h.CreateBook) // create book
		book.GET("/list", h.GetBookList) // get book list
		book.GET("/", h.GetBookById) // get book by id
		book.PUT("/", h.UpdateBook) // update book
		book.DELETE("/", h.DeleteBook) // delete book
	}
	genr := apiv1.Group("/genr")
		{
			genr.POST("/", h.creteGenre)    // create genre
			genr.GET("/", h.GetGenre)    // get genre
			genr.GET("/id", h.GetGenreById) // get genre by id
			genr.PUT("/", h.UpdateGenre)    // update genre
			genr.DELETE("/", h.DeleteGenre) // delete genre
		}


	return router
}
