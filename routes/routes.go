package routes

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/pkg/api"
	ac "github.com/thavel/goban/pkg/auth"
	"github.com/thavel/goban/routes/absence"
	"github.com/thavel/goban/routes/auth"
	"github.com/thavel/goban/routes/team"
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

	// API teams
	router.GET("/teams", team.List)
	router.POST("/teams", team.Create)
	router.GET("/teams/:tid", team.Get)
	router.PATCH("/teams/:tid", team.Update)
	router.DELETE("/teams/:tid", team.Delete)

	// API absences
	router.GET("/reasons", absence.ListReasons)
	router.POST("/reasons", absence.CreateReason)
	router.GET("/reasons/:rid", absence.GetReason)
	router.PATCH("/reasons/:rid", absence.UpdateReason)
	router.DELETE("/reasons/:rid", absence.DeleteReason)
	router.GET("/absences", absence.ListAllAbsences)
	router.GET("/users/:uid/absences", absence.ListAbsences)
	router.POST("/users/:uid/absences", absence.CreateAbsence)
	router.GET("/users/:uid/absences/:aid", absence.GetAbsence)
	router.PATCH("/users/:uid/absences/:aid", absence.UpdateAbsence)
	router.DELETE("/users/:uid/absences/:aid", absence.DeleteAbsence)

	// API auth
	router.POST("/auth/token", auth.Auth)

	return api.CORS(ac.RBAC(router.Handler))
}
