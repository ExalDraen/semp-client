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
)

// NewGetMsgVpnQueueSubscriptionParams creates a new GetMsgVpnQueueSubscriptionParams object
// with the default values initialized.
func NewGetMsgVpnQueueSubscriptionParams() *GetMsgVpnQueueSubscriptionParams {
	var ()
	return &GetMsgVpnQueueSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnQueueSubscriptionParamsWithTimeout creates a new GetMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnQueueSubscriptionParamsWithTimeout(timeout time.Duration) *GetMsgVpnQueueSubscriptionParams {
	var ()
	return &GetMsgVpnQueueSubscriptionParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnQueueSubscriptionParamsWithContext creates a new GetMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnQueueSubscriptionParamsWithContext(ctx context.Context) *GetMsgVpnQueueSubscriptionParams {
	var ()
	return &GetMsgVpnQueueSubscriptionParams{

		Context: ctx,
	}
}

// NewGetMsgVpnQueueSubscriptionParamsWithHTTPClient creates a new GetMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnQueueSubscriptionParamsWithHTTPClient(client *http.Client) *GetMsgVpnQueueSubscriptionParams {
	var ()
	return &GetMsgVpnQueueSubscriptionParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnQueueSubscriptionParams contains all the parameters to send to the API endpoint
for the get msg vpn queue subscription operation typically these are written to a http.Request
*/
type GetMsgVpnQueueSubscriptionParams struct {

	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*QueueName
	  The name of the Queue.

	*/
	QueueName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string
	/*SubscriptionTopic
	  The topic of the Subscription.

	*/
	SubscriptionTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithTimeout(timeout time.Duration) *GetMsgVpnQueueSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithContext(ctx context.Context) *GetMsgVpnQueueSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithHTTPClient(client *http.Client) *GetMsgVpnQueueSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnQueueSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueName adds the queueName to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithQueueName(queueName string) *GetMsgVpnQueueSubscriptionParams {
	o.SetQueueName(queueName)
	return o
}

// SetQueueName adds the queueName to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetQueueName(queueName string) {
	o.QueueName = queueName
}

// WithSelect adds the selectVar to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithSelect(selectVar []string) *GetMsgVpnQueueSubscriptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithSubscriptionTopic adds the subscriptionTopic to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) WithSubscriptionTopic(subscriptionTopic string) *GetMsgVpnQueueSubscriptionParams {
	o.SetSubscriptionTopic(subscriptionTopic)
	return o
}

// SetSubscriptionTopic adds the subscriptionTopic to the get msg vpn queue subscription params
func (o *GetMsgVpnQueueSubscriptionParams) SetSubscriptionTopic(subscriptionTopic string) {
	o.SubscriptionTopic = subscriptionTopic
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnQueueSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param queueName
	if err := r.SetPathParam("queueName", o.QueueName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	// path param subscriptionTopic
	if err := r.SetPathParam("subscriptionTopic", o.SubscriptionTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
