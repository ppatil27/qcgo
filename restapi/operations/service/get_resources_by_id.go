// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetResourcesByIDHandlerFunc turns a function with the right signature into a get resources by Id handler
type GetResourcesByIDHandlerFunc func(GetResourcesByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourcesByIDHandlerFunc) Handle(params GetResourcesByIDParams) middleware.Responder {
	return fn(params)
}

// GetResourcesByIDHandler interface for that can handle valid get resources by Id params
type GetResourcesByIDHandler interface {
	Handle(GetResourcesByIDParams) middleware.Responder
}

// NewGetResourcesByID creates a new http.Handler for the get resources by Id operation
func NewGetResourcesByID(ctx *middleware.Context, handler GetResourcesByIDHandler) *GetResourcesByID {
	return &GetResourcesByID{Context: ctx, Handler: handler}
}

/* GetResourcesByID swagger:route GET /{resourceId}/ Service getResourcesById

Get  resources

*/
type GetResourcesByID struct {
	Context *middleware.Context
	Handler GetResourcesByIDHandler
}

func (o *GetResourcesByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetResourcesByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
