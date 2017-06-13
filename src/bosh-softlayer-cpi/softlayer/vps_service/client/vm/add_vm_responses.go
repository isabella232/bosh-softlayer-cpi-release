package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"bosh-softlayer-cpi/softlayer/vps_service/models"
)

// AddVMReader is a Reader for the AddVM structure.
type AddVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewAddVMOK creates a AddVMOK with default headers values
func NewAddVMOK() *AddVMOK {
	return &AddVMOK{}
}

// WithPayload adds the payload to the add Vm o k response
func (o *AddVMOK) WithPayload(payload string) *AddVMOK {
	o.Payload = payload
	return o
}

/*AddVMOK handles this case with default header values.

add a new vm into the pool
*/
type AddVMOK struct {
	Payload string
}

func (o *AddVMOK) Error() string {
	return fmt.Sprintf("[POST /vms][%d] addVmOK  %+v", 200, o.Payload)
}

func (o *AddVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddVMDefault creates a AddVMDefault with default headers values
func NewAddVMDefault(code int) *AddVMDefault {
	return &AddVMDefault{
		_statusCode: code,
	}
}

/*AddVMDefault handles this case with default header values.

unexpected error
*/
type AddVMDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add Vm default response
func (o *AddVMDefault) Code() int {
	return o._statusCode
}

func (o *AddVMDefault) Error() string {
	return fmt.Sprintf("[POST /vms][%d] addVm default  %+v", o._statusCode, o.Payload)
}

func (o *AddVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
