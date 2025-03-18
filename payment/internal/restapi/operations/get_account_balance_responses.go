// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/crypto-market/payment/internal/models"
)

// GetAccountBalanceOKCode is the HTTP code returned for type GetAccountBalanceOK
const GetAccountBalanceOKCode int = 200

/*
GetAccountBalanceOK Success operation

swagger:response getAccountBalanceOK
*/
type GetAccountBalanceOK struct {

	/*
	  In: Body
	*/
	Payload []*GetAccountBalanceOKBodyItems0 `json:"body,omitempty"`
}

// NewGetAccountBalanceOK creates GetAccountBalanceOK with default headers values
func NewGetAccountBalanceOK() *GetAccountBalanceOK {

	return &GetAccountBalanceOK{}
}

// WithPayload adds the payload to the get account balance o k response
func (o *GetAccountBalanceOK) WithPayload(payload []*GetAccountBalanceOKBodyItems0) *GetAccountBalanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account balance o k response
func (o *GetAccountBalanceOK) SetPayload(payload []*GetAccountBalanceOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountBalanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetAccountBalanceOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountBalanceUnauthorizedCode is the HTTP code returned for type GetAccountBalanceUnauthorized
const GetAccountBalanceUnauthorizedCode int = 401

/*
GetAccountBalanceUnauthorized Unauthorized

swagger:response getAccountBalanceUnauthorized
*/
type GetAccountBalanceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAccountBalanceUnauthorized creates GetAccountBalanceUnauthorized with default headers values
func NewGetAccountBalanceUnauthorized() *GetAccountBalanceUnauthorized {

	return &GetAccountBalanceUnauthorized{}
}

// WithPayload adds the payload to the get account balance unauthorized response
func (o *GetAccountBalanceUnauthorized) WithPayload(payload *models.Error) *GetAccountBalanceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account balance unauthorized response
func (o *GetAccountBalanceUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountBalanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
