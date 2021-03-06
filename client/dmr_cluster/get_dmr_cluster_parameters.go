// Code generated by go-swagger; DO NOT EDIT.

package dmr_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDmrClusterParams creates a new GetDmrClusterParams object
// with the default values initialized.
func NewGetDmrClusterParams() *GetDmrClusterParams {
	var ()
	return &GetDmrClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDmrClusterParamsWithTimeout creates a new GetDmrClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDmrClusterParamsWithTimeout(timeout time.Duration) *GetDmrClusterParams {
	var ()
	return &GetDmrClusterParams{

		timeout: timeout,
	}
}

// NewGetDmrClusterParamsWithContext creates a new GetDmrClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDmrClusterParamsWithContext(ctx context.Context) *GetDmrClusterParams {
	var ()
	return &GetDmrClusterParams{

		Context: ctx,
	}
}

// NewGetDmrClusterParamsWithHTTPClient creates a new GetDmrClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDmrClusterParamsWithHTTPClient(client *http.Client) *GetDmrClusterParams {
	var ()
	return &GetDmrClusterParams{
		HTTPClient: client,
	}
}

/*GetDmrClusterParams contains all the parameters to send to the API endpoint
for the get dmr cluster operation typically these are written to a http.Request
*/
type GetDmrClusterParams struct {

	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dmr cluster params
func (o *GetDmrClusterParams) WithTimeout(timeout time.Duration) *GetDmrClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dmr cluster params
func (o *GetDmrClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dmr cluster params
func (o *GetDmrClusterParams) WithContext(ctx context.Context) *GetDmrClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dmr cluster params
func (o *GetDmrClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dmr cluster params
func (o *GetDmrClusterParams) WithHTTPClient(client *http.Client) *GetDmrClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dmr cluster params
func (o *GetDmrClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDmrClusterName adds the dmrClusterName to the get dmr cluster params
func (o *GetDmrClusterParams) WithDmrClusterName(dmrClusterName string) *GetDmrClusterParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the get dmr cluster params
func (o *GetDmrClusterParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithSelect adds the selectVar to the get dmr cluster params
func (o *GetDmrClusterParams) WithSelect(selectVar []string) *GetDmrClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get dmr cluster params
func (o *GetDmrClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetDmrClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dmrClusterName
	if err := r.SetPathParam("dmrClusterName", o.DmrClusterName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
