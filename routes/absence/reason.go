package absence

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/api"
	"github.com/thavel/goban/pkg/database"
)

func ListReasons(ctx *fasthttp.RequestCtx) {
	query := database.DB().Model(&models.Reason{})
	query = api.Filtered(query, ctx)

	var entities []models.Reason
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

func CreateReason(ctx *fasthttp.RequestCtx) {
	entity, err := models.NewReason(ctx.PostBody())
	if err != nil {
		api.Response(ctx, 400, fmt.Sprintf("invalid format: %s", err.Error()))
		return
	}

	db := database.DB()
	if res := db.Create(&entity); res.Error != nil {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 201, entity)
}

func GetReason(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("rid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	var entity models.Reason
	if res := database.DB().Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("%d doesn't exist", uuid))
		return
	}

	api.Response(ctx, 200, entity)
}

func UpdateReason(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("rid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Reason
	if res := db.Find(&entity, uuid); res.Error != nil {
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

func DeleteReason(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("rid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Reason
	if res := db.Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("reason %d doesn't exist", uuid))
		return
	}
	if res := db.Delete(&entity); res.Error != nil || res.RowsAffected == 0 {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 204, nil)
}
