package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const (
	QueryPath      = "/query"
	PlaygroundPath = "/__playground"
)

func New(srv *handler.Server) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group(QueryPath)
	{
		r.POST("", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Status healthy")
	})

	e.GET(PlaygroundPath, playgroundHandler())

	return e
}

// Defining the Playground handler
func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", QueryPath)
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
