package rest

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetupRouter(server *echo.Echo, handler *Handler) {

	// - health check
	server.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "services up and running... ")
	})

	newsRoute := server.Group("/bareksa/news")
	{
		newsRoute.GET("", handler.NewsHandler.FindAll)
		newsRoute.GET("/:id", handler.NewsHandler.FindOne)
		newsRoute.POST("", handler.NewsHandler.Create)
		newsRoute.PUT("/:id", handler.NewsHandler.UpdateByID)
		newsRoute.DELETE("/:id", handler.NewsHandler.DeleteByID)
	}

	tagRoute := server.Group("/bareksa/tag")
	{
		tagRoute.GET("", handler.TagHandler.FindAll)
		tagRoute.GET("/:id", handler.TagHandler.FindOne)
		tagRoute.POST("", handler.TagHandler.Create)
		tagRoute.PUT("/:id", handler.TagHandler.UpdateByID)
		tagRoute.DELETE("/:id", handler.TagHandler.DeleteByID)
	}

	topicRoute := server.Group("/bareksa/topic")
	{
		topicRoute.GET("", handler.TopicHandler.FindAll)
	}

}
