// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/axard/todo-list/internal/restmodels"
)

// ReadTodosReader is a Reader for the ReadTodos structure.
type ReadTodosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTodosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTodosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadTodosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTodosOK creates a ReadTodosOK with default headers values
func NewReadTodosOK() *ReadTodosOK {
	return &ReadTodosOK{}
}

/*ReadTodosOK handles this case with default header values.

Список дел
*/
type ReadTodosOK struct {
	Payload *restmodels.Itemlist
}

func (o *ReadTodosOK) Error() string {
	return fmt.Sprintf("[GET /][%d] readTodosOK  %+v", 200, o.Payload)
}

func (o *ReadTodosOK) GetPayload() *restmodels.Itemlist {
	return o.Payload
}

func (o *ReadTodosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Itemlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTodosDefault creates a ReadTodosDefault with default headers values
func NewReadTodosDefault(code int) *ReadTodosDefault {
	return &ReadTodosDefault{
		_statusCode: code,
	}
}

/*ReadTodosDefault handles this case with default header values.

Ошибка
*/
type ReadTodosDefault struct {
	_statusCode int

	Payload *restmodels.Error
}

// Code gets the status code for the read todos default response
func (o *ReadTodosDefault) Code() int {
	return o._statusCode
}

func (o *ReadTodosDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] readTodos default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTodosDefault) GetPayload() *restmodels.Error {
	return o.Payload
}

func (o *ReadTodosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
