package routes

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/pkg/api"
	ac "github.com/thavel/goban/pkg/auth"
	"github.com/thavel/goban/routes/auth"
	"github.com/thavel/goban/routes/user"
)

// Handler gives API endpoints.
func Handler() func(*fasthttp.RequestCtx) {
	router := fasthttprouter.New()

	// API users
	router.GET("/users", user.List)
	router.POST("/users", user.Create)
	router.GET("/users/:uid", user.Get)
	router.PATCH("/users/:uid", user.Update)
	router.DELETE("/users/:uid", user.Delete)

	// API auth
	router.POST("/auth/token", auth.Auth)

	return ac.RBAC(api.CORS(router.Handler))
}
