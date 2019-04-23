package auth

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/api"
	"github.com/thavel/goban/pkg/crypto"
	"github.com/thavel/goban/pkg/database"
	"github.com/thavel/goban/pkg/jwt"
)

func Auth(ctx *fasthttp.RequestCtx) {
	// Unmarshal payload
	body := ctx.PostBody()
	var creds Credentials
	err := json.Unmarshal(body, &creds)
	if err != nil || !creds.Valid() {
		api.Response(ctx, 400, "invalid credentials")
		return
	}

	// Get and check user
	db := database.DB()
	var user models.User
	if creds.User != nil {
		res := db.Where("email = ?", *creds.User).Find(&user)
		if res.Error != nil {
			api.Response(ctx, 400, "invalid credentials")
			return
		}
		if !crypto.IsEqual(user.Password, *creds.Password) {
			api.Response(ctx, 400, "invalid credentials")
			return
		}
	} else {
		data, err := jwt.Parse(*creds.Token)
		if err != nil {
			api.Response(ctx, 400, "invalid credentials")
			return
		}
		if res := db.Find(&user, data.User); res.Error != nil {
			api.Response(ctx, 400, "invalid credentials")
			return
		}
	}

	// Check role
	if user.Role == nil {
		api.Response(ctx, 400, "user has no permission")
		return
	}
	var role models.Role
	if res := db.Where("name = ?", *user.Role).Find(&role); res.Error != nil {
		api.Response(ctx, 500, fmt.Sprintf("role %s doesn't exist", *user.Role))
		return
	}

	// Provide JWT token
	token, err := jwt.Build(jwt.Data{User: user.ID, Role: *user.Role})
	if err != nil {
		api.Response(ctx, 500, nil)
		return
	}
	api.Response(ctx, 200, map[string]string{"token": *token})
}
