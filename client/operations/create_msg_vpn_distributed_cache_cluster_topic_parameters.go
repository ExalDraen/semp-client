// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/ExalDraen/semp-client/models"
)

// NewCreateMsgVpnDistributedCacheClusterTopicParams creates a new CreateMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized.
func NewCreateMsgVpnDistributedCacheClusterTopicParams() *CreateMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterTopicParamsWithTimeout creates a new CreateMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnDistributedCacheClusterTopicParamsWithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterTopicParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterTopicParamsWithContext creates a new CreateMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnDistributedCacheClusterTopicParamsWithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterTopicParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient creates a new CreateMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterTopicParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnDistributedCacheClusterTopicParams contains all the parameters to send to the API endpoint
for the create msg vpn distributed cache cluster topic operation typically these are written to a http.Request
*/
type CreateMsgVpnDistributedCacheClusterTopicParams struct {

	/*Body
	  The Topic object's attributes.

	*/
	Body *models.MsgVpnDistributedCacheClusterTopic
	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*ClusterName
	  The name of the Cache Cluster.

	*/
	ClusterName string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithBody(body *models.MsgVpnDistributedCacheClusterTopic) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetBody(body *models.MsgVpnDistributedCacheClusterTopic) {
	o.Body = body
}

// WithCacheName adds the cacheName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithCacheName(cacheName string) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithClusterName(clusterName string) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WithSelect(selectVar []string) *CreateMsgVpnDistributedCacheClusterTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn distributed cache cluster topic params
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnDistributedCacheClusterTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
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
