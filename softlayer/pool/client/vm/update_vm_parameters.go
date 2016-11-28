package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/models"
)

// NewUpdateVMParams creates a new UpdateVMParams object
// with the default values initialized.
func NewUpdateVMParams() *UpdateVMParams {
	var ()
	return &UpdateVMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMParamsWithTimeout creates a new UpdateVMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVMParamsWithTimeout(timeout time.Duration) *UpdateVMParams {
	var ()
	return &UpdateVMParams{

		timeout: timeout,
	}
}

// NewUpdateVMParamsWithContext creates a new UpdateVMParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVMParamsWithContext(ctx context.Context) *UpdateVMParams {
	var ()
	return &UpdateVMParams{

		Context: ctx,
	}
}

/*UpdateVMParams contains all the parameters to send to the API endpoint
for the update Vm operation typically these are written to a http.Request
*/
type UpdateVMParams struct {

	/*Body
	  Vm object that needs to be added to the pool

	*/
	Body *models.VM

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the update Vm params
func (o *UpdateVMParams) WithTimeout(timeout time.Duration) *UpdateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm params
func (o *UpdateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm params
func (o *UpdateVMParams) WithContext(ctx context.Context) *UpdateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm params
func (o *UpdateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the update Vm params
func (o *UpdateVMParams) WithBody(body *models.VM) *UpdateVMParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Vm params
func (o *UpdateVMParams) SetBody(body *models.VM) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VM)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
