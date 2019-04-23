package user

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/thavel/goban/pkg/auth"
	"github.com/valyala/fasthttp"
)

// FromUser filters DB queries by user.
func FromUser(query *gorm.DB, ctx *fasthttp.RequestCtx) *gorm.DB {
	return query.Where("user_id = ?", GetUser(ctx))
}

// GetUser from the URL path.
func GetUser(ctx *fasthttp.RequestCtx) uint {
	id := ctx.UserValue("uid").(string)
	var uid uint
	switch id {
	case "":
		uid = 0
	case "me":
		data, err := auth.Token(ctx)
		if err != nil {
			uid = 0
			break
		}
		uid = data.User
	default:
		parsed, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			uid = 0
			break
		}
		uid = uint(parsed)
	}
	return uid
}
