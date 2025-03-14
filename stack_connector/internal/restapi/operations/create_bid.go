// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateBidHandlerFunc turns a function with the right signature into a create bid handler
type CreateBidHandlerFunc func(CreateBidParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBidHandlerFunc) Handle(params CreateBidParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateBidHandler interface for that can handle valid create bid params
type CreateBidHandler interface {
	Handle(CreateBidParams, interface{}) middleware.Responder
}

// NewCreateBid creates a new http.Handler for the create bid operation
func NewCreateBid(ctx *middleware.Context, handler CreateBidHandler) *CreateBid {
	return &CreateBid{Context: ctx, Handler: handler}
}

/*
	CreateBid swagger:route POST /bid createBid

Create bid
*/
type CreateBid struct {
	Context *middleware.Context
	Handler CreateBidHandler
}

func (o *CreateBid) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateBidParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateBidOKBody create bid o k body
//
// swagger:model CreateBidOKBody
type CreateBidOKBody struct {

	// bid id
	BidID int64 `json:"bid_id,omitempty"`
}

// Validate validates this create bid o k body
func (o *CreateBidOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create bid o k body based on context it is used
func (o *CreateBidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBidOKBody) UnmarshalBinary(b []byte) error {
	var res CreateBidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
