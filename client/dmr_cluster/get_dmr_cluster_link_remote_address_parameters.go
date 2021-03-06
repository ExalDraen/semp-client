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

// NewGetDmrClusterLinkRemoteAddressParams creates a new GetDmrClusterLinkRemoteAddressParams object
// with the default values initialized.
func NewGetDmrClusterLinkRemoteAddressParams() *GetDmrClusterLinkRemoteAddressParams {
	var ()
	return &GetDmrClusterLinkRemoteAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDmrClusterLinkRemoteAddressParamsWithTimeout creates a new GetDmrClusterLinkRemoteAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDmrClusterLinkRemoteAddressParamsWithTimeout(timeout time.Duration) *GetDmrClusterLinkRemoteAddressParams {
	var ()
	return &GetDmrClusterLinkRemoteAddressParams{

		timeout: timeout,
	}
}

// NewGetDmrClusterLinkRemoteAddressParamsWithContext creates a new GetDmrClusterLinkRemoteAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDmrClusterLinkRemoteAddressParamsWithContext(ctx context.Context) *GetDmrClusterLinkRemoteAddressParams {
	var ()
	return &GetDmrClusterLinkRemoteAddressParams{

		Context: ctx,
	}
}

// NewGetDmrClusterLinkRemoteAddressParamsWithHTTPClient creates a new GetDmrClusterLinkRemoteAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDmrClusterLinkRemoteAddressParamsWithHTTPClient(client *http.Client) *GetDmrClusterLinkRemoteAddressParams {
	var ()
	return &GetDmrClusterLinkRemoteAddressParams{
		HTTPClient: client,
	}
}

/*GetDmrClusterLinkRemoteAddressParams contains all the parameters to send to the API endpoint
for the get dmr cluster link remote address operation typically these are written to a http.Request
*/
type GetDmrClusterLinkRemoteAddressParams struct {

	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*RemoteAddress
	  The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed).

	*/
	RemoteAddress string
	/*RemoteNodeName
	  The name of the node at the remote end of the Link.

	*/
	RemoteNodeName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithTimeout(timeout time.Duration) *GetDmrClusterLinkRemoteAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithContext(ctx context.Context) *GetDmrClusterLinkRemoteAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithHTTPClient(client *http.Client) *GetDmrClusterLinkRemoteAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDmrClusterName adds the dmrClusterName to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithDmrClusterName(dmrClusterName string) *GetDmrClusterLinkRemoteAddressParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithRemoteAddress adds the remoteAddress to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithRemoteAddress(remoteAddress string) *GetDmrClusterLinkRemoteAddressParams {
	o.SetRemoteAddress(remoteAddress)
	return o
}

// SetRemoteAddress adds the remoteAddress to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetRemoteAddress(remoteAddress string) {
	o.RemoteAddress = remoteAddress
}

// WithRemoteNodeName adds the remoteNodeName to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithRemoteNodeName(remoteNodeName string) *GetDmrClusterLinkRemoteAddressParams {
	o.SetRemoteNodeName(remoteNodeName)
	return o
}

// SetRemoteNodeName adds the remoteNodeName to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetRemoteNodeName(remoteNodeName string) {
	o.RemoteNodeName = remoteNodeName
}

// WithSelect adds the selectVar to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) WithSelect(selectVar []string) *GetDmrClusterLinkRemoteAddressParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get dmr cluster link remote address params
func (o *GetDmrClusterLinkRemoteAddressParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetDmrClusterLinkRemoteAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dmrClusterName
	if err := r.SetPathParam("dmrClusterName", o.DmrClusterName); err != nil {
		return err
	}

	// path param remoteAddress
	if err := r.SetPathParam("remoteAddress", o.RemoteAddress); err != nil {
		return err
	}

	// path param remoteNodeName
	if err := r.SetPathParam("remoteNodeName", o.RemoteNodeName); err != nil {
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
