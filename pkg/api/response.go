package api

import (
	"encoding/json"
	"net/url"
	"path"
	"strings"

	"github.com/valyala/fasthttp"
)

// Response from an API endpoint.
func Response(ctx *fasthttp.RequestCtx, code int, data interface{}) {
	ctx.SetStatusCode(code)
	if data == nil {
		return
	}

	ctx.SetContentType("application/json")
	switch d := data.(type) {
	case error:
		data = map[string]string{
			"message": d.Error(),
		}
	case string:
		data = map[string]string{
			"message": string(d),
		}
	}
	payload, err := json.Marshal(data)
	if err != nil {
		Response(ctx, 500, "something went wrong formatting response payload")
	}
	ctx.SetBody(payload)
}

// Redirect handler to a specific path.
func Redirect(to string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		p := string(ctx.Request.URI().Path())
		if u := ctx.Request.Header.Peek("X-Original-URL"); u != nil {
			pu, err := url.Parse(string(u))
			if err != nil {
				return
			}
			p = pu.Path
		} else if u := ctx.Request.Header.Peek("X-Original-Uri"); u != nil {
			p = string(u)
		}
		target := path.Join(p, to)
		if strings.HasSuffix(to, "/") {
			target += "/"
		}

		ctx.Redirect(target, 301)
	}
}
