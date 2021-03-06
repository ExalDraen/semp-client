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

// NewUpdateMsgVpnReplicatedTopicParams creates a new UpdateMsgVpnReplicatedTopicParams object
// with the default values initialized.
func NewUpdateMsgVpnReplicatedTopicParams() *UpdateMsgVpnReplicatedTopicParams {
	var ()
	return &UpdateMsgVpnReplicatedTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnReplicatedTopicParamsWithTimeout creates a new UpdateMsgVpnReplicatedTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnReplicatedTopicParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnReplicatedTopicParams {
	var ()
	return &UpdateMsgVpnReplicatedTopicParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnReplicatedTopicParamsWithContext creates a new UpdateMsgVpnReplicatedTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnReplicatedTopicParamsWithContext(ctx context.Context) *UpdateMsgVpnReplicatedTopicParams {
	var ()
	return &UpdateMsgVpnReplicatedTopicParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnReplicatedTopicParamsWithHTTPClient creates a new UpdateMsgVpnReplicatedTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnReplicatedTopicParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnReplicatedTopicParams {
	var ()
	return &UpdateMsgVpnReplicatedTopicParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnReplicatedTopicParams contains all the parameters to send to the API endpoint
for the update msg vpn replicated topic operation typically these are written to a http.Request
*/
type UpdateMsgVpnReplicatedTopicParams struct {

	/*Body
	  The Replicated Topic object's attributes.

	*/
	Body *models.MsgVpnReplicatedTopic
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*ReplicatedTopic
	  The topic for applying replication. Published messages matching this topic will be replicated to the standby site.

	*/
	ReplicatedTopic string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnReplicatedTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithContext(ctx context.Context) *UpdateMsgVpnReplicatedTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnReplicatedTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithBody(body *models.MsgVpnReplicatedTopic) *UpdateMsgVpnReplicatedTopicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetBody(body *models.MsgVpnReplicatedTopic) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnReplicatedTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithReplicatedTopic adds the replicatedTopic to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithReplicatedTopic(replicatedTopic string) *UpdateMsgVpnReplicatedTopicParams {
	o.SetReplicatedTopic(replicatedTopic)
	return o
}

// SetReplicatedTopic adds the replicatedTopic to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetReplicatedTopic(replicatedTopic string) {
	o.ReplicatedTopic = replicatedTopic
}

// WithSelect adds the selectVar to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) WithSelect(selectVar []string) *UpdateMsgVpnReplicatedTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn replicated topic params
func (o *UpdateMsgVpnReplicatedTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnReplicatedTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param replicatedTopic
	if err := r.SetPathParam("replicatedTopic", o.ReplicatedTopic); err != nil {
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
