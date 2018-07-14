package api

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
)

// Filtered query using query params
func Filtered(query *gorm.DB, ctx *fasthttp.RequestCtx) *gorm.DB {
	model := query.Value
	args := ctx.QueryArgs()
	tref := reflect.TypeOf(model).Elem()

	for i := 0; i < tref.NumField(); i++ {
		field := tref.Field(i)
		name := gorm.ToDBName(field.Name)

		// Ignore private fields
		initChar := string([]byte{field.Name[0]})
		if name == "model" || initChar != strings.ToUpper(initChar) {
			continue
		}
		// Ignore private JSON fields
		jname := field.Tag.Get("json")
		switch jname {
		case "-":
			continue
		case "":
			jname = name
		}

		// Get the query param
		value := args.Peek(jname)
		if value != nil && len(value) > 0 {
			// Add where clause to the existing query
			raw := fmt.Sprintf("%s = ?", name)
			query = query.Where(raw, value)
		}
	}

	return query
}

// Paged query using query params
func Paged(query *gorm.DB, ctx *fasthttp.RequestCtx) *gorm.DB {
	args := ctx.QueryArgs()
	page, err := args.GetUint("page")
	if err != nil {
		page = 0
	}
	limit, err := args.GetUint("size")
	if err != nil {
		limit = 10
	}
	offset := page * limit
	query = query.Offset(offset).Limit(limit)

	return query
}
