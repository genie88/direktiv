// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewInstanceLogsParams creates a new InstanceLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstanceLogsParams() *InstanceLogsParams {
	return &InstanceLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceLogsParamsWithTimeout creates a new InstanceLogsParams object
// with the ability to set a timeout on a request.
func NewInstanceLogsParamsWithTimeout(timeout time.Duration) *InstanceLogsParams {
	return &InstanceLogsParams{
		timeout: timeout,
	}
}

// NewInstanceLogsParamsWithContext creates a new InstanceLogsParams object
// with the ability to set a context for a request.
func NewInstanceLogsParamsWithContext(ctx context.Context) *InstanceLogsParams {
	return &InstanceLogsParams{
		Context: ctx,
	}
}

// NewInstanceLogsParamsWithHTTPClient creates a new InstanceLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstanceLogsParamsWithHTTPClient(client *http.Client) *InstanceLogsParams {
	return &InstanceLogsParams{
		HTTPClient: client,
	}
}

/* InstanceLogsParams contains all the parameters to send to the API endpoint
   for the instance logs operation.

   Typically these are written to a http.Request.
*/
type InstanceLogsParams struct {

	/* Instance.

	   target instance id
	*/
	Instance string

	/* Namespace.

	   target namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instance logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceLogsParams) WithDefaults() *InstanceLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instance logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instance logs params
func (o *InstanceLogsParams) WithTimeout(timeout time.Duration) *InstanceLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance logs params
func (o *InstanceLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance logs params
func (o *InstanceLogsParams) WithContext(ctx context.Context) *InstanceLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance logs params
func (o *InstanceLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance logs params
func (o *InstanceLogsParams) WithHTTPClient(client *http.Client) *InstanceLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance logs params
func (o *InstanceLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstance adds the instance to the instance logs params
func (o *InstanceLogsParams) WithInstance(instance string) *InstanceLogsParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the instance logs params
func (o *InstanceLogsParams) SetInstance(instance string) {
	o.Instance = instance
}

// WithNamespace adds the namespace to the instance logs params
func (o *InstanceLogsParams) WithNamespace(namespace string) *InstanceLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the instance logs params
func (o *InstanceLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instance
	if err := r.SetPathParam("instance", o.Instance); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
