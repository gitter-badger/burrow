package request

import (
	"github.com/carrot/go-base-api/controllers"
	"github.com/carrot/go-base-api/middleware"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

func BuildEcho() *echo.Echo {
	// ----------
	// Framework
	// ----------

	e := echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(echo_middleware.Logger())
	e.Use(middleware.Recover())

	// ------------
	// Controllers
	// ------------

	topicsController := new(controllers.TopicsController)

	// ----------
	// Endpoints
	// ----------

	e.Get("/topics", topicsController.Index)
	e.Get("/topics/:id", topicsController.Show)
	e.Post("/topics", topicsController.Create)
	e.Put("/topics/:id", topicsController.Update)
	e.Delete("/topics/:id", topicsController.Delete)

	return e
}
