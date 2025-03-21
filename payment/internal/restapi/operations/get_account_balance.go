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

// GetAccountBalanceHandlerFunc turns a function with the right signature into a get account balance handler
type GetAccountBalanceHandlerFunc func(GetAccountBalanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountBalanceHandlerFunc) Handle(params GetAccountBalanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAccountBalanceHandler interface for that can handle valid get account balance params
type GetAccountBalanceHandler interface {
	Handle(GetAccountBalanceParams, interface{}) middleware.Responder
}

// NewGetAccountBalance creates a new http.Handler for the get account balance operation
func NewGetAccountBalance(ctx *middleware.Context, handler GetAccountBalanceHandler) *GetAccountBalance {
	return &GetAccountBalance{Context: ctx, Handler: handler}
}

/*
	GetAccountBalance swagger:route GET /account/balance getAccountBalance

# Get user's balance

Returns balance of all cryptocurrencies
*/
type GetAccountBalance struct {
	Context *middleware.Context
	Handler GetAccountBalanceHandler
}

func (o *GetAccountBalance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountBalanceParams()
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

// GetAccountBalanceOKBodyItems0 get account balance o k body items0
//
// swagger:model GetAccountBalanceOKBodyItems0
type GetAccountBalanceOKBodyItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

// Validate validates this get account balance o k body items0
func (o *GetAccountBalanceOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get account balance o k body items0 based on context it is used
func (o *GetAccountBalanceOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountBalanceOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountBalanceOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetAccountBalanceOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
