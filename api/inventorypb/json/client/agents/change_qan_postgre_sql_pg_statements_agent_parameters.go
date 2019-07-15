// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewChangeQANPostgreSQLPgStatementsAgentParams creates a new ChangeQANPostgreSQLPgStatementsAgentParams object
// with the default values initialized.
func NewChangeQANPostgreSQLPgStatementsAgentParams() *ChangeQANPostgreSQLPgStatementsAgentParams {
	var ()
	return &ChangeQANPostgreSQLPgStatementsAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeQANPostgreSQLPgStatementsAgentParamsWithTimeout creates a new ChangeQANPostgreSQLPgStatementsAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeQANPostgreSQLPgStatementsAgentParamsWithTimeout(timeout time.Duration) *ChangeQANPostgreSQLPgStatementsAgentParams {
	var ()
	return &ChangeQANPostgreSQLPgStatementsAgentParams{

		timeout: timeout,
	}
}

// NewChangeQANPostgreSQLPgStatementsAgentParamsWithContext creates a new ChangeQANPostgreSQLPgStatementsAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeQANPostgreSQLPgStatementsAgentParamsWithContext(ctx context.Context) *ChangeQANPostgreSQLPgStatementsAgentParams {
	var ()
	return &ChangeQANPostgreSQLPgStatementsAgentParams{

		Context: ctx,
	}
}

// NewChangeQANPostgreSQLPgStatementsAgentParamsWithHTTPClient creates a new ChangeQANPostgreSQLPgStatementsAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeQANPostgreSQLPgStatementsAgentParamsWithHTTPClient(client *http.Client) *ChangeQANPostgreSQLPgStatementsAgentParams {
	var ()
	return &ChangeQANPostgreSQLPgStatementsAgentParams{
		HTTPClient: client,
	}
}

/*ChangeQANPostgreSQLPgStatementsAgentParams contains all the parameters to send to the API endpoint
for the change QAN postgre SQL pg statements agent operation typically these are written to a http.Request
*/
type ChangeQANPostgreSQLPgStatementsAgentParams struct {

	/*Body*/
	Body ChangeQANPostgreSQLPgStatementsAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) WithTimeout(timeout time.Duration) *ChangeQANPostgreSQLPgStatementsAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) WithContext(ctx context.Context) *ChangeQANPostgreSQLPgStatementsAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) WithHTTPClient(client *http.Client) *ChangeQANPostgreSQLPgStatementsAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) WithBody(body ChangeQANPostgreSQLPgStatementsAgentBody) *ChangeQANPostgreSQLPgStatementsAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change QAN postgre SQL pg statements agent params
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) SetBody(body ChangeQANPostgreSQLPgStatementsAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeQANPostgreSQLPgStatementsAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}