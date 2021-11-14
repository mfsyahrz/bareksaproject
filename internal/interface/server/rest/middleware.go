package rest

import (
	"math/rand"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oklog/ulid"

	"github.com/mfsyahrz/bareksaproject/internal/interface/ioc"
	"github.com/mfsyahrz/bareksaproject/internal/interface/validator"
)

func generateThreadId() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	uniqueID := ulid.MustNew(ulid.Timestamp(t), entropy)
	return uniqueID.String()
}

func SetupMiddleware(server *echo.Echo, container *ioc.IOC) {

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", "Pv", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			container.Log.Info("incoming request")
			return h(c)
		}
	})

	server.Validator = validator.NewValidator()
}
