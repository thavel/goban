package absence

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/api"
	"github.com/thavel/goban/pkg/database"
	"github.com/thavel/goban/routes/user"
)

func filtered(query *gorm.DB, ctx *fasthttp.RequestCtx) *gorm.DB {
	args := ctx.QueryArgs()

	// Handles 'from'/'to' key
	from := dbTime(args.Peek("from"))
	to := dbTime(args.Peek("to"))
	if from != "" && to != "" {
		query = query.Where(
			"(`from` >= ? OR `to` >= ?) OR (`to` <= ? OR `from` <= ?)",
			from, from, to, to,
		)
	} else if from != "" {
		query = query.Where("`from` >= ? OR `to` >= ?", from, from)
	} else if to != "" {
		query = query.Where("`to` <= ? OR `from` <= ?", to, to)
	}

	// Handles 'team' key
	team := args.Peek("team")
	if team != nil {
		tid, err := strconv.ParseUint(string(team), 10, 64)
		if err == nil {
			query = query.Joins(
				"JOIN users ON users.id = absences.user_id AND users.team_id = ?",
				tid,
			)
		}
	}

	return query
}

func ListAllAbsences(ctx *fasthttp.RequestCtx) {
	query := database.DB().Model(&models.Absence{})
	query = filtered(query, ctx)

	var entities []models.Absence
	res := api.Paged(query, ctx).Find(&entities)
	if res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	var count int
	res = query.Count(&count)
	if res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	ctx.Response.Header.Set("X-Paginated-Items", strconv.Itoa(count))
	api.Response(ctx, 200, entities)
}

func ListAbsences(ctx *fasthttp.RequestCtx) {
	query := database.DB().Model(&models.Absence{})
	query = user.FromUser(query, ctx)
	query = filtered(query, ctx)

	var entities []models.Absence
	res := api.Paged(query, ctx).Find(&entities)
	if res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	var count int
	res = query.Count(&count)
	if res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	ctx.Response.Header.Set("X-Paginated-Items", strconv.Itoa(count))
	api.Response(ctx, 200, entities)
}

func CreateAbsence(ctx *fasthttp.RequestCtx) {
	entity, err := models.NewAbsence(ctx.PostBody())
	if err != nil {
		api.Response(ctx, 400, fmt.Sprintf("invalid format: %s", err.Error()))
		return
	}

	uid := user.GetUser(ctx)
	if uid == 0 {
		api.Response(ctx, 400, "invalid user id")
		return
	}
	entity.UserID = uid

	db := database.DB()
	if res := db.Create(&entity); res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 201, entity)
}

func GetAbsence(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("aid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	var entity models.Absence
	query := user.FromUser(database.DB(), ctx)
	if res := query.Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("%d doesn't exist", uuid))
		return
	}

	api.Response(ctx, 200, entity)
}

func UpdateAbsence(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("aid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Absence
	if res := user.FromUser(db, ctx).Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("%d doesn't exist", uuid))
		return
	}
	if err := json.Unmarshal(ctx.PostBody(), &entity); err != nil {
		api.Response(ctx, 400, fmt.Sprintf("invalid updates: %s", err))
		return
	}
	if err := entity.Validate(); err != nil {
		api.Response(ctx, 400, fmt.Sprintf("invalid updates: %s", err))
		return
	}
	if res := db.Save(&entity); res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 200, entity)
}

func DeleteAbsence(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("aid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Absence
	if res := user.FromUser(db, ctx).Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("absence %d doesn't exist", uuid))
		return
	}
	if res := db.Delete(&entity); res.Error != nil || res.RowsAffected == 0 {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 204, nil)
}
