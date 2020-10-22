// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadTodosHandlerFunc turns a function with the right signature into a read todos handler
type ReadTodosHandlerFunc func(ReadTodosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadTodosHandlerFunc) Handle(params ReadTodosParams) middleware.Responder {
	return fn(params)
}

// ReadTodosHandler interface for that can handle valid read todos params
type ReadTodosHandler interface {
	Handle(ReadTodosParams) middleware.Responder
}

// NewReadTodos creates a new http.Handler for the read todos operation
func NewReadTodos(ctx *middleware.Context, handler ReadTodosHandler) *ReadTodos {
	return &ReadTodos{Context: ctx, Handler: handler}
}

/*ReadTodos swagger:route GET / todos readTodos

ReadTodos read todos API

*/
type ReadTodos struct {
	Context *middleware.Context
	Handler ReadTodosHandler
}

func (o *ReadTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadTodosParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
