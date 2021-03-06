// Code generated by go-swagger; DO NOT EDIT.

package access_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccessKeyPoliciesHandlerFunc turns a function with the right signature into a get access key policies handler
type GetAccessKeyPoliciesHandlerFunc func(GetAccessKeyPoliciesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccessKeyPoliciesHandlerFunc) Handle(params GetAccessKeyPoliciesParams) middleware.Responder {
	return fn(params)
}

// GetAccessKeyPoliciesHandler interface for that can handle valid get access key policies params
type GetAccessKeyPoliciesHandler interface {
	Handle(GetAccessKeyPoliciesParams) middleware.Responder
}

// NewGetAccessKeyPolicies creates a new http.Handler for the get access key policies operation
func NewGetAccessKeyPolicies(ctx *middleware.Context, handler GetAccessKeyPoliciesHandler) *GetAccessKeyPolicies {
	return &GetAccessKeyPolicies{Context: ctx, Handler: handler}
}

/*GetAccessKeyPolicies swagger:route GET /accesskeys/{key_id}/policies AccessKeys getAccessKeyPolicies

get Policy

获取指定 access_key_id 匹配的策略


*/
type GetAccessKeyPolicies struct {
	Context *middleware.Context
	Handler GetAccessKeyPoliciesHandler
}

func (o *GetAccessKeyPolicies) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccessKeyPoliciesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
