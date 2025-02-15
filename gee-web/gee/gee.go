package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

// Engine is the core structure of the framework that holds the router.
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute registers a handler for a given HTTP method and URL pattern.
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route: %s - %q", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run starts the HTTP server on the specified address.
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine) // Starts the server with Engine handling requests
}

// ServeHTTP handles incoming HTTP requests, satisfying the http.Handler interface.
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)
}
