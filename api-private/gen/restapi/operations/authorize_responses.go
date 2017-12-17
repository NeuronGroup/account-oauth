// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronOauth/oauth/api-private/gen/models"
)

// AuthorizeOKCode is the HTTP code returned for type AuthorizeOK
const AuthorizeOKCode int = 200

/*AuthorizeOK ok

swagger:response authorizeOK
*/
type AuthorizeOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthorizationCode `json:"body,omitempty"`
}

// NewAuthorizeOK creates AuthorizeOK with default headers values
func NewAuthorizeOK() *AuthorizeOK {
	return &AuthorizeOK{}
}

// WithPayload adds the payload to the authorize o k response
func (o *AuthorizeOK) WithPayload(payload *models.AuthorizationCode) *AuthorizeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize o k response
func (o *AuthorizeOK) SetPayload(payload *models.AuthorizationCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
