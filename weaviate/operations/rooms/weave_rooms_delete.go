package rooms


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveRoomsDeleteHandlerFunc turns a function with the right signature into a weave rooms delete handler
type WeaveRoomsDeleteHandlerFunc func(WeaveRoomsDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveRoomsDeleteHandlerFunc) Handle(params WeaveRoomsDeleteParams) middleware.Responder {
	return fn(params)
}

// WeaveRoomsDeleteHandler interface for that can handle valid weave rooms delete params
type WeaveRoomsDeleteHandler interface {
	Handle(WeaveRoomsDeleteParams) middleware.Responder
}

// NewWeaveRoomsDelete creates a new http.Handler for the weave rooms delete operation
func NewWeaveRoomsDelete(ctx *middleware.Context, handler WeaveRoomsDeleteHandler) *WeaveRoomsDelete {
	return &WeaveRoomsDelete{Context: ctx, Handler: handler}
}

/*WeaveRoomsDelete swagger:route DELETE /places/{placeId}/rooms/{roomId} rooms weaveRoomsDelete

Deletes a room.

*/
type WeaveRoomsDelete struct {
	Context *middleware.Context
	Handler WeaveRoomsDeleteHandler
}

func (o *WeaveRoomsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveRoomsDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
