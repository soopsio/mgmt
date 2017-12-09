// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTaskHandlerFunc turns a function with the right signature into a get task handler
type GetTaskHandlerFunc func(GetTaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskHandlerFunc) Handle(params GetTaskParams) middleware.Responder {
	return fn(params)
}

// GetTaskHandler interface for that can handle valid get task params
type GetTaskHandler interface {
	Handle(GetTaskParams) middleware.Responder
}

// NewGetTask creates a new http.Handler for the get task operation
func NewGetTask(ctx *middleware.Context, handler GetTaskHandler) *GetTask {
	return &GetTask{Context: ctx, Handler: handler}
}

/*GetTask swagger:route GET /tasks/{task_id} Tasks getTask

Get Task

获取任务状态


*/
type GetTask struct {
	Context *middleware.Context
	Handler GetTaskHandler
}

func (o *GetTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTaskParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
