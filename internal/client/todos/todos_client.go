// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOne(params *CreateOneParams) (*CreateOneCreated, error)

	DeleteOne(params *DeleteOneParams) (*DeleteOneNoContent, error)

	ReadTodos(params *ReadTodosParams) (*ReadTodosOK, error)

	UpdateOne(params *UpdateOneParams) (*UpdateOneOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateOne create one API
*/
func (a *Client) CreateOne(params *CreateOneParams) (*CreateOneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOne",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOneCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteOne delete one API
*/
func (a *Client) DeleteOne(params *DeleteOneParams) (*DeleteOneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOne",
		Method:             "DELETE",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOneNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReadTodos read todos API
*/
func (a *Client) ReadTodos(params *ReadTodosParams) (*ReadTodosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadTodosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readTodos",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadTodosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadTodosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadTodosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOne update one API
*/
func (a *Client) UpdateOne(params *UpdateOneParams) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOne",
		Method:             "PUT",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.goswagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}