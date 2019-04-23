package team

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/api"
	"github.com/thavel/goban/pkg/database"
)

// List all teams
func List(ctx *fasthttp.RequestCtx) {
	query := database.DB().Model(&models.Team{})
	query = api.Filtered(query, ctx)

	var entities []models.Team
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

// Create a new team
func Create(ctx *fasthttp.RequestCtx) {
	entity, err := models.NewTeam(ctx.PostBody())
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

// Get a specific team
func Get(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("tid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	var entity models.Team
	if res := database.DB().Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("%d doesn't exist", uuid))
		return
	}

	api.Response(ctx, 200, entity)
}

// Update a team
func Update(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("tid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Team
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

// Delete a team
func Delete(ctx *fasthttp.RequestCtx) {
	uuid, err := strconv.ParseUint(ctx.UserValue("tid").(string), 10, 64)
	if err != nil {
		api.Response(ctx, 400, "invalid id")
		return
	}

	db := database.DB()
	var entity models.Team
	if res := db.Find(&entity, uuid); res.Error != nil {
		api.Response(ctx, 404, fmt.Sprintf("team %d doesn't exist", uuid))
		return
	}
	if res := db.Delete(&entity); res.Error != nil || res.RowsAffected == 0 {
		api.Response(ctx, 500, "something went wrong")
		return
	}

	api.Response(ctx, 204, nil)
}
