package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teguhbayu/go-restapi/controllers"
)

func (r routes) bookRoute(v1 *gin.RouterGroup) {
	bookRouter := v1.Group("/book")

	bookRouter.POST("/", controllers.CreateBook)
	bookRouter.GET("/", controllers.GetBooks)
	bookRouter.GET("/:id", controllers.GetBookById)
	bookRouter.PUT("/:id", controllers.UpdateBook)
	bookRouter.DELETE("/:id", controllers.DeleteBook)
}
