// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostResourceByAppCodeHandlerFunc turns a function with the right signature into a post resource by app code handler
type PostResourceByAppCodeHandlerFunc func(PostResourceByAppCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostResourceByAppCodeHandlerFunc) Handle(params PostResourceByAppCodeParams) middleware.Responder {
	return fn(params)
}

// PostResourceByAppCodeHandler interface for that can handle valid post resource by app code params
type PostResourceByAppCodeHandler interface {
	Handle(PostResourceByAppCodeParams) middleware.Responder
}

// NewPostResourceByAppCode creates a new http.Handler for the post resource by app code operation
func NewPostResourceByAppCode(ctx *middleware.Context, handler PostResourceByAppCodeHandler) *PostResourceByAppCode {
	return &PostResourceByAppCode{Context: ctx, Handler: handler}
}

/* PostResourceByAppCode swagger:route POST /Resource Service postResourceByAppCode

Create a new resource for a given application.

*/
type PostResourceByAppCode struct {
	Context *middleware.Context
	Handler PostResourceByAppCodeHandler
}

func (o *PostResourceByAppCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostResourceByAppCodeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
