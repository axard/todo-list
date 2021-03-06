// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/axard/todo-list/internal/restmodels"
)

// PatchOneOKCode is the HTTP code returned for type PatchOneOK
const PatchOneOKCode int = 200

/*PatchOneOK Ок

swagger:response patchOneOK
*/
type PatchOneOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Item `json:"body,omitempty"`
}

// NewPatchOneOK creates PatchOneOK with default headers values
func NewPatchOneOK() *PatchOneOK {

	return &PatchOneOK{}
}

// WithPayload adds the payload to the patch one o k response
func (o *PatchOneOK) WithPayload(payload *restmodels.Item) *PatchOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one o k response
func (o *PatchOneOK) SetPayload(payload *restmodels.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchOneDefault Ошибка

swagger:response patchOneDefault
*/
type PatchOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPatchOneDefault creates PatchOneDefault with default headers values
func NewPatchOneDefault(code int) *PatchOneDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch one default response
func (o *PatchOneDefault) WithStatusCode(code int) *PatchOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch one default response
func (o *PatchOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch one default response
func (o *PatchOneDefault) WithPayload(payload *restmodels.Error) *PatchOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one default response
func (o *PatchOneDefault) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
