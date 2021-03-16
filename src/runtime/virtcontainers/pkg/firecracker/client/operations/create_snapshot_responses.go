// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// CreateSnapshotReader is a Reader for the CreateSnapshot structure.
type CreateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCreateSnapshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSnapshotNoContent creates a CreateSnapshotNoContent with default headers values
func NewCreateSnapshotNoContent() *CreateSnapshotNoContent {
	return &CreateSnapshotNoContent{}
}

/*CreateSnapshotNoContent handles this case with default header values.

Snapshot created
*/
type CreateSnapshotNoContent struct {
}

func (o *CreateSnapshotNoContent) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotNoContent ", 204)
}

func (o *CreateSnapshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnapshotBadRequest creates a CreateSnapshotBadRequest with default headers values
func NewCreateSnapshotBadRequest() *CreateSnapshotBadRequest {
	return &CreateSnapshotBadRequest{}
}

/*CreateSnapshotBadRequest handles this case with default header values.

Snapshot cannot be created due to bad input
*/
type CreateSnapshotBadRequest struct {
	Payload *models.Error
}

func (o *CreateSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotDefault creates a CreateSnapshotDefault with default headers values
func NewCreateSnapshotDefault(code int) *CreateSnapshotDefault {
	return &CreateSnapshotDefault{
		_statusCode: code,
	}
}

/*CreateSnapshotDefault handles this case with default header values.

Internal server error
*/
type CreateSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create snapshot default response
func (o *CreateSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *CreateSnapshotDefault) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
