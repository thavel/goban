package auth

import (
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/api"
	"github.com/thavel/goban/pkg/database"
	"github.com/thavel/goban/pkg/jwt"
	"github.com/thavel/goban/pkg/logger"
)

func Token(ctx *fasthttp.RequestCtx) (*jwt.Data, error) {
	// Get auth header
	header := ctx.Request.Header.Peek("Authorization")
	if header == nil {
		return &jwt.Data{Role: Anonymous}, nil
	}
	result := bearer.FindStringSubmatch(string(header))
	if len(result) != 2 {
		return nil, errors.New("invalid authorization header format")
	}
	token := result[1]

	// Parse token
	data, err := jwt.Parse(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if data.Expired() {
		return nil, errors.New("expirated token")
	}
	return data, nil
}

// RBAC middleware function.
func RBAC(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		var data *jwt.Data
		var err error
		method := string(ctx.Request.Header.Method())
		path := string(ctx.Request.URI().Path())

		// Defer access log
		defer func() {
			c := ctx.Response.StatusCode()
			verb := "granted"
			if c == 401 || c == 403 {
				verb = "rejected"
			}
			r := Anonymous
			if data != nil {
				r = data.Role
			}
			logger.Debugf("%s %s, role %s %s", method, path, r, verb)
		}()

		// Invalid header/token
		data, err = Token(ctx)
		if err != nil {
			api.Response(ctx, 401, err)
			return
		}

		// Check token
		if data.Role != Anonymous {
			db := database.DB()
			memkey := strconv.FormatUint(uint64(data.User), 10)
			if _, err := memcache.get(memkey); err != nil {
				var user models.User
				if res := db.Find(&user, data.User); res.Error != nil {
					api.Response(ctx, 401, "unknown user")
					return
				}
				if user.Role == nil || *user.Role != data.Role {
					api.Response(ctx, 401, "user's role has changed")
					return
				}
				memcache.set(memkey, nil)
			}
		}

		// Authorization enforcement
		ok, err := enforcer.EnforceSafe(data.Role, path, method)
		if err != nil {
			api.Response(ctx, 500, nil)
			return
		}
		if !ok {
			api.Response(ctx, 403, "insufficient permissions")
			return
		}

		// Process to request handler
		handler(ctx)
	})
}
