// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/soopsio/mgmt/api/v1/models"
)

// DelPolicyCreatedCode is the HTTP code returned for type DelPolicyCreated
const DelPolicyCreatedCode int = 201

/*DelPolicyCreated 无返回值

swagger:response delPolicyCreated
*/
type DelPolicyCreated struct {
}

// NewDelPolicyCreated creates DelPolicyCreated with default headers values
func NewDelPolicyCreated() *DelPolicyCreated {
	return &DelPolicyCreated{}
}

// WriteResponse to the client
func (o *DelPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*DelPolicyDefault Unexpected error

swagger:response delPolicyDefault
*/
type DelPolicyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDelPolicyDefault creates DelPolicyDefault with default headers values
func NewDelPolicyDefault(code int) *DelPolicyDefault {
	if code <= 0 {
		code = 500
	}

	return &DelPolicyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the del policy default response
func (o *DelPolicyDefault) WithStatusCode(code int) *DelPolicyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the del policy default response
func (o *DelPolicyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the del policy default response
func (o *DelPolicyDefault) WithPayload(payload *models.Error) *DelPolicyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the del policy default response
func (o *DelPolicyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelPolicyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
