package gee

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by the framework.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine is the core structure of the framework that holds the router.
type Engine struct {
	router map[string]HandlerFunc
}

// New creates a new Engine instance with an initialized router.
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute registers a handler for a given HTTP method and URL pattern.
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // Unique key combining HTTP method and URL pattern
	log.Printf("Route: %s - %q", method, pattern)
	engine.router[key] = handler
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
	key := request.Method + "-" + request.URL.Path // Generates the key to look up the handler
	if handler, ok := engine.router[key]; ok {
		handler(writer, request) // Executes the matched handler
	} else {
		log.Printf("404 not found: %s", request.URL.Path) // Logs if no handler matches the request
	}
}
