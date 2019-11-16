package main

import (
    "flag"
    "fmt"
    "log"
    "templates"
    
    "github.com/fasthttp/router"
    "github.com/valyala/fasthttp"
)

var addr = flag.String( "addr", "127.0.0.1:10000",
    "TCP address to listen to for incoming connections" )
    
func Index( ctx *fasthttp.RequestCtx ) {
	p := &templates.MainPage{}
	templates.WritePageTemplate(ctx, p)
    ctx.SetContentType("text/html; charset=utf-8")
}

func Hello( ctx *fasthttp.RequestCtx ) {
	fmt.Fprintf( ctx, "Hello, %s!\n", ctx.UserValue( "name" ) )
}

func QueryArgs( ctx *fasthttp.RequestCtx ) {
	name := ctx.QueryArgs().Peek( "name" )
	fmt.Fprintf( ctx, "Pong! %s\n", string( name ) )
}
    
func main() {
    flag.Parse()
    
    r := router.New()
	r.GET( "/", Index )
	r.GET( "/hello/:name", Hello )
    r.GET( "/ping", QueryArgs )

    log.Fatal( fasthttp.ListenAndServe( *addr, r.Handler ) )
}