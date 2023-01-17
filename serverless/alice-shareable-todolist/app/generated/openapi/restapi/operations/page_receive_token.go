// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PageReceiveTokenHandlerFunc turns a function with the right signature into a page receive token handler
type PageReceiveTokenHandlerFunc func(PageReceiveTokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PageReceiveTokenHandlerFunc) Handle(params PageReceiveTokenParams) middleware.Responder {
	return fn(params)
}

// PageReceiveTokenHandler interface for that can handle valid page receive token params
type PageReceiveTokenHandler interface {
	Handle(PageReceiveTokenParams) middleware.Responder
}

// NewPageReceiveToken creates a new http.Handler for the page receive token operation
func NewPageReceiveToken(ctx *middleware.Context, handler PageReceiveTokenHandler) *PageReceiveToken {
	return &PageReceiveToken{Context: ctx, Handler: handler}
}

/* PageReceiveToken swagger:route GET /receive-token pageReceiveToken

PageReceiveToken page receive token API

*/
type PageReceiveToken struct {
	Context *middleware.Context
	Handler PageReceiveTokenHandler
}

func (o *PageReceiveToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPageReceiveTokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
