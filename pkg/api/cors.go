package api

import (
	"github.com/valyala/fasthttp"
)

// CORS middleware function.
func CORS(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		origin := string(ctx.Request.Header.Peek("Origin"))
		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,PATCH,HEAD,PUT,POST,DELETE,OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,Authorization")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "X-Paginated-Items")

		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(200)
			return
		}
		handler(ctx)
	})
}
