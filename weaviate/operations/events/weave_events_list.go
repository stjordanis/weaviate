package events


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveEventsListHandlerFunc turns a function with the right signature into a weave events list handler
type WeaveEventsListHandlerFunc func(WeaveEventsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveEventsListHandlerFunc) Handle(params WeaveEventsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveEventsListHandler interface for that can handle valid weave events list params
type WeaveEventsListHandler interface {
	Handle(WeaveEventsListParams, interface{}) middleware.Responder
}

// NewWeaveEventsList creates a new http.Handler for the weave events list operation
func NewWeaveEventsList(ctx *middleware.Context, handler WeaveEventsListHandler) *WeaveEventsList {
	return &WeaveEventsList{Context: ctx, Handler: handler}
}

/*WeaveEventsList swagger:route GET /events events weaveEventsList

Lists events.

*/
type WeaveEventsList struct {
	Context *middleware.Context
	Handler WeaveEventsListHandler
}

func (o *WeaveEventsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveEventsListParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
