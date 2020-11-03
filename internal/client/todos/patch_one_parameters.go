// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPatchOneParams creates a new PatchOneParams object
// with the default values initialized.
func NewPatchOneParams() *PatchOneParams {
	var ()
	return &PatchOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOneParamsWithTimeout creates a new PatchOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchOneParamsWithTimeout(timeout time.Duration) *PatchOneParams {
	var ()
	return &PatchOneParams{

		timeout: timeout,
	}
}

// NewPatchOneParamsWithContext creates a new PatchOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchOneParamsWithContext(ctx context.Context) *PatchOneParams {
	var ()
	return &PatchOneParams{

		Context: ctx,
	}
}

// NewPatchOneParamsWithHTTPClient creates a new PatchOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchOneParamsWithHTTPClient(client *http.Client) *PatchOneParams {
	var ()
	return &PatchOneParams{
		HTTPClient: client,
	}
}

/*PatchOneParams contains all the parameters to send to the API endpoint
for the patch one operation typically these are written to a http.Request
*/
type PatchOneParams struct {

	/*Body*/
	Body PatchOneBody
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch one params
func (o *PatchOneParams) WithTimeout(timeout time.Duration) *PatchOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch one params
func (o *PatchOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch one params
func (o *PatchOneParams) WithContext(ctx context.Context) *PatchOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch one params
func (o *PatchOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch one params
func (o *PatchOneParams) WithHTTPClient(client *http.Client) *PatchOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch one params
func (o *PatchOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch one params
func (o *PatchOneParams) WithBody(body PatchOneBody) *PatchOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch one params
func (o *PatchOneParams) SetBody(body PatchOneBody) {
	o.Body = body
}

// WithID adds the id to the patch one params
func (o *PatchOneParams) WithID(id int64) *PatchOneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch one params
func (o *PatchOneParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}