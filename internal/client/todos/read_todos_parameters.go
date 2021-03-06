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

// NewReadTodosParams creates a new ReadTodosParams object
// with the default values initialized.
func NewReadTodosParams() *ReadTodosParams {
	var (
		limitDefault = int32(20)
	)
	return &ReadTodosParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReadTodosParamsWithTimeout creates a new ReadTodosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadTodosParamsWithTimeout(timeout time.Duration) *ReadTodosParams {
	var (
		limitDefault = int32(20)
	)
	return &ReadTodosParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewReadTodosParamsWithContext creates a new ReadTodosParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadTodosParamsWithContext(ctx context.Context) *ReadTodosParams {
	var (
		limitDefault = int32(20)
	)
	return &ReadTodosParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewReadTodosParamsWithHTTPClient creates a new ReadTodosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadTodosParamsWithHTTPClient(client *http.Client) *ReadTodosParams {
	var (
		limitDefault = int32(20)
	)
	return &ReadTodosParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*ReadTodosParams contains all the parameters to send to the API endpoint
for the read todos operation typically these are written to a http.Request
*/
type ReadTodosParams struct {

	/*Limit*/
	Limit *int32
	/*Since*/
	Since *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read todos params
func (o *ReadTodosParams) WithTimeout(timeout time.Duration) *ReadTodosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read todos params
func (o *ReadTodosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read todos params
func (o *ReadTodosParams) WithContext(ctx context.Context) *ReadTodosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read todos params
func (o *ReadTodosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read todos params
func (o *ReadTodosParams) WithHTTPClient(client *http.Client) *ReadTodosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read todos params
func (o *ReadTodosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the read todos params
func (o *ReadTodosParams) WithLimit(limit *int32) *ReadTodosParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the read todos params
func (o *ReadTodosParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the read todos params
func (o *ReadTodosParams) WithSince(since *int64) *ReadTodosParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the read todos params
func (o *ReadTodosParams) SetSince(since *int64) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *ReadTodosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
