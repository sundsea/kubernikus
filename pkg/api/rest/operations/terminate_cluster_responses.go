// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// TerminateClusterAcceptedCode is the HTTP code returned for type TerminateClusterAccepted
const TerminateClusterAcceptedCode int = 202

/*TerminateClusterAccepted OK

swagger:response terminateClusterAccepted
*/
type TerminateClusterAccepted struct {
}

// NewTerminateClusterAccepted creates TerminateClusterAccepted with default headers values
func NewTerminateClusterAccepted() *TerminateClusterAccepted {
	return &TerminateClusterAccepted{}
}

// WriteResponse to the client
func (o *TerminateClusterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

/*TerminateClusterDefault Error

swagger:response terminateClusterDefault
*/
type TerminateClusterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTerminateClusterDefault creates TerminateClusterDefault with default headers values
func NewTerminateClusterDefault(code int) *TerminateClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &TerminateClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terminate cluster default response
func (o *TerminateClusterDefault) WithStatusCode(code int) *TerminateClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terminate cluster default response
func (o *TerminateClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terminate cluster default response
func (o *TerminateClusterDefault) WithPayload(payload *models.Error) *TerminateClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terminate cluster default response
func (o *TerminateClusterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerminateClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}