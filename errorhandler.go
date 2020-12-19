package main

import "github.com/valyala/fasthttp"

// Handles every request other than GET method also bad requests
func errorHandler(ctx *fasthttp.RequestCtx, err error) {
	// Clear all existing header
	ctx.Response.Reset()

	if ctx.IsOptions() {
		ctx.Response.Header.Set("Server", "kuda")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", origins)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Host,Accept,Accept-Encoding,Connection,User-Agent")
		ctx.Response.SetStatusCode(204)
	} else if !ctx.IsGet() {
		ctx.Response.SetStatusCode(405)
	} else {
		ctx.Response.SetStatusCode(400)
	}
}
