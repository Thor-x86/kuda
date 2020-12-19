package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

// Handle each request
func handler(ctx *fasthttp.RequestCtx) {
	// Making sure all files already loaded
	if pathMap == nil {
		panic("Not loaded yet!")
	}

	// For CORS acknowledge
	if ctx.IsOptions() {
		ctx.SetStatusCode(204)
		return
	}

	// Start incoming request timestamp
	requestTime := ctx.Time()

	path := string(ctx.Path())
	responseCode := 200
	isCompressed := false

	// Remove trailing slash
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		ctx.Redirect(strings.TrimSuffix(path, "/"), 308)
		return
	}

	// Process response based on specific condition
	path = strings.TrimPrefix(path, "/")
	if file, isExist := (*pathMap)[path]; isExist {
		// Condition: found static file, e.g. www.mydomain.com/img/something.png
		ctx.Response.SetBodyRaw(file.data)
		isCompressed = file.isCompressed
		ctx.SetContentType(file.mime)
	} else if file, isExist := (*pathMap)[path+"/index.html"]; isExist {
		// Condition: has index.html in the subpath
		//            e.g. www.mydomain.com/path/subpath => www.mydomain.com/path/subpath/index.html
		ctx.Response.SetBodyRaw(file.data)
		isCompressed = file.isCompressed
		ctx.SetContentType(file.mime)
	} else if file, isExist := (*pathMap)["index.html"]; isExist {
		// Condition: found nothing in path map, so index.html will handle the route
		ctx.Response.SetBodyRaw(file.data)
		isCompressed = file.isCompressed
		ctx.SetContentType(file.mime)
	} else {
		// Condition: developer forgot to add index.html into public root directory
		responseCode = 404
		ctx.Response.SetBodyString("<html><body><h1>There is no root index.html</h1></body></html>")
		ctx.SetContentType("text/html")
	}

	// Tell the browser whether it's compressed or not
	if isCompressed {
		ctx.Response.Header.Set("Content-Encoding", "gzip")
	} else {
		ctx.Response.Header.Set("Content-Encoding", "")
	}

	// Set headers
	ctx.Response.SetStatusCode(responseCode)
	ctx.Response.Header.Set("Server", "kuda")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", origins)
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Host,Accept,Accept-Encoding,Connection,User-Agent")

	// Get elapsed request-response time and client IPv4
	elapsed := time.Now().Sub(requestTime)
	clientIP := ctx.RemoteIP().String()

	// Format the elapsed time
	elapsedValue := float32(elapsed.Nanoseconds())
	elapsedUnit := "ns"
	if elapsedValue >= 1000 {
		elapsedValue /= 1000
		elapsedUnit = "μs"
	}
	if elapsedValue >= 1000 {
		elapsedValue /= 1000
		elapsedUnit = "ms"
	}
	if elapsedValue >= 1000 {
		elapsedValue /= 1000
		elapsedUnit = "s"
	}

	// Print report to log
	fmt.Printf("[KUDA] %d | %.3f%s\t| %s\t| /%s\n", responseCode, elapsedValue, elapsedUnit, clientIP, path)
}
