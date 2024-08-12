package config

import (
	"goethe/auth"
	"goethe/handlers"

	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	e.Static("/public", "public")
	e.File("/favicon.ico", "public/favicon.ico")
	e.File("/robots.txt", "robots.txt")

	e.RouteNotFound("/*", handlers.Route404)

	e.GET("/", handlers.Home)

	e.GET("/bookshelf", handlers.BookshelfBase)
	e.POST("/bookshelf/add-book", auth.WithJWT(handlers.AddBook, handlers.Route404))
	e.DELETE("/bookshelf/remove-book", auth.WithJWT(handlers.RemoveBook, handlers.Route404))
	e.GET("/bookshelf/book", handlers.HandleBook)

	e.GET("/blog", handlers.BlogBase)
	e.GET("/blog/post", handlers.BlogPost)
}
