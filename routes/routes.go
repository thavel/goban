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

type handler struct {
	method  string
	path    string
	handler fasthttp.RequestHandler
}

// Handler gives API endpoints.
func Handler() func(*fasthttp.RequestCtx) {
	router := fasthttprouter.New()

	handlers := []handler{
		// API auth
		handler{"POST", "/auth/token", auth.Auth},

		// API users
		handler{"GET", "/users", user.List},
		handler{"POST", "/users", user.Create},
		handler{"GET", "/users/:uid", user.Get},
		handler{"PATCH", "/users/:uid", user.Update},
		handler{"DELETE", "/users/:uid", user.Delete},

		// API teams
		handler{"GET", "/teams", team.List},
		handler{"POST", "/teams", team.Create},
		handler{"GET", "/teams/:tid", team.Get},
		handler{"PATCH", "/teams/:tid", team.Update},
		handler{"DELETE", "/teams/:tid", team.Delete},

		// API absences
		handler{"GET", "/reasons", absence.ListReasons},
		handler{"POST", "/reasons", absence.CreateReason},
		handler{"GET", "/reasons/:rid", absence.GetReason},
		handler{"PATCH", "/reasons/:rid", absence.UpdateReason},
		handler{"DELETE", "/reasons/:rid", absence.DeleteReason},
		handler{"GET", "/absences", absence.ListAllAbsences},
		handler{"GET", "/users/:uid/absences", absence.ListAbsences},
		handler{"POST", "/users/:uid/absences", absence.CreateAbsence},
		handler{"GET", "/users/:uid/absences/:aid", absence.GetAbsence},
		handler{"PATCH", "/users/:uid/absences/:aid", absence.UpdateAbsence},
		handler{"DELETE", "/users/:uid/absences/:aid", absence.DeleteAbsence},
	}
	for _, h := range handlers {
		router.Handle(h.method, h.path, ac.RBAC(h.handler))
	}

	// UI
	router.GET("/ui/*filepath", fasthttp.FSHandler("ui/", 1))
	router.GET("/", api.Redirect("/ui/"))

	return api.CORS(router.Handler)
}
