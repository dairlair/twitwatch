// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SignupHandlerFunc turns a function with the right signature into a signup handler
type SignupHandlerFunc func(SignupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SignupHandlerFunc) Handle(params SignupParams) middleware.Responder {
	return fn(params)
}

// SignupHandler interface for that can handle valid signup params
type SignupHandler interface {
	Handle(SignupParams) middleware.Responder
}

// NewSignup creates a new http.Handler for the signup operation
func NewSignup(ctx *middleware.Context, handler SignupHandler) *Signup {
	return &Signup{Context: ctx, Handler: handler}
}

/*Signup swagger:route POST /signup signup

Signup signup API

*/
type Signup struct {
	Context *middleware.Context
	Handler SignupHandler
}

func (o *Signup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSignupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
