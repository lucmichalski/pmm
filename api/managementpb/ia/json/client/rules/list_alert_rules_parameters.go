// Code generated by go-swagger; DO NOT EDIT.

package rules

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
)

// NewListAlertRulesParams creates a new ListAlertRulesParams object
// with the default values initialized.
func NewListAlertRulesParams() *ListAlertRulesParams {
	var ()
	return &ListAlertRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAlertRulesParamsWithTimeout creates a new ListAlertRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAlertRulesParamsWithTimeout(timeout time.Duration) *ListAlertRulesParams {
	var ()
	return &ListAlertRulesParams{

		timeout: timeout,
	}
}

// NewListAlertRulesParamsWithContext creates a new ListAlertRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAlertRulesParamsWithContext(ctx context.Context) *ListAlertRulesParams {
	var ()
	return &ListAlertRulesParams{

		Context: ctx,
	}
}

// NewListAlertRulesParamsWithHTTPClient creates a new ListAlertRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAlertRulesParamsWithHTTPClient(client *http.Client) *ListAlertRulesParams {
	var ()
	return &ListAlertRulesParams{
		HTTPClient: client,
	}
}

/*ListAlertRulesParams contains all the parameters to send to the API endpoint
for the list alert rules operation typically these are written to a http.Request
*/
type ListAlertRulesParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list alert rules params
func (o *ListAlertRulesParams) WithTimeout(timeout time.Duration) *ListAlertRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alert rules params
func (o *ListAlertRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alert rules params
func (o *ListAlertRulesParams) WithContext(ctx context.Context) *ListAlertRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alert rules params
func (o *ListAlertRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alert rules params
func (o *ListAlertRulesParams) WithHTTPClient(client *http.Client) *ListAlertRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alert rules params
func (o *ListAlertRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list alert rules params
func (o *ListAlertRulesParams) WithBody(body interface{}) *ListAlertRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list alert rules params
func (o *ListAlertRulesParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlertRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
