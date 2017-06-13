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

// OrderVMByFilterReader is a Reader for the OrderVMByFilter structure.
type OrderVMByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderVMByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderVMByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewOrderVMByFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewOrderVMByFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewOrderVMByFilterOK creates a OrderVMByFilterOK with default headers values
func NewOrderVMByFilterOK() *OrderVMByFilterOK {
	return &OrderVMByFilterOK{}
}

// WithPayload adds the payload to the order Vm by filter o k response
func (o *OrderVMByFilterOK) WithPayload(payload *models.VMResponse) *OrderVMByFilterOK {
	o.Payload = payload
	return o
}

/*OrderVMByFilterOK handles this case with default header values.

successful operation
*/
type OrderVMByFilterOK struct {
	Payload *models.VMResponse
}

func (o *OrderVMByFilterOK) Error() string {
	return fmt.Sprintf("[POST /vms/order][%d] orderVmByFilterOK  %+v", 200, o.Payload)
}

func (o *OrderVMByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderVMByFilterNotFound creates a OrderVMByFilterNotFound with default headers values
func NewOrderVMByFilterNotFound() *OrderVMByFilterNotFound {
	return &OrderVMByFilterNotFound{}
}

/*OrderVMByFilterNotFound handles this case with default header values.

vm not found
*/
type OrderVMByFilterNotFound struct {
}

func (o *OrderVMByFilterNotFound) Error() string {
	return fmt.Sprintf("[POST /vms/order][%d] orderVmByFilterNotFound ", 404)
}

func (o *OrderVMByFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrderVMByFilterDefault creates a OrderVMByFilterDefault with default headers values
func NewOrderVMByFilterDefault(code int) *OrderVMByFilterDefault {
	return &OrderVMByFilterDefault{
		_statusCode: code,
	}
}

/*OrderVMByFilterDefault handles this case with default header values.

unexpected error
*/
type OrderVMByFilterDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the order Vm by filter default response
func (o *OrderVMByFilterDefault) Code() int {
	return o._statusCode
}

func (o *OrderVMByFilterDefault) Error() string {
	return fmt.Sprintf("[POST /vms/order][%d] orderVmByFilter default  %+v", o._statusCode, o.Payload)
}

func (o *OrderVMByFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
