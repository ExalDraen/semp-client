// Code generated by go-swagger; DO NOT EDIT.

package msg_vpn

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

// NewCreateMsgVpnSequencedTopicParams creates a new CreateMsgVpnSequencedTopicParams object
// with the default values initialized.
func NewCreateMsgVpnSequencedTopicParams() *CreateMsgVpnSequencedTopicParams {
	var ()
	return &CreateMsgVpnSequencedTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnSequencedTopicParamsWithTimeout creates a new CreateMsgVpnSequencedTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnSequencedTopicParamsWithTimeout(timeout time.Duration) *CreateMsgVpnSequencedTopicParams {
	var ()
	return &CreateMsgVpnSequencedTopicParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnSequencedTopicParamsWithContext creates a new CreateMsgVpnSequencedTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnSequencedTopicParamsWithContext(ctx context.Context) *CreateMsgVpnSequencedTopicParams {
	var ()
	return &CreateMsgVpnSequencedTopicParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnSequencedTopicParamsWithHTTPClient creates a new CreateMsgVpnSequencedTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnSequencedTopicParamsWithHTTPClient(client *http.Client) *CreateMsgVpnSequencedTopicParams {
	var ()
	return &CreateMsgVpnSequencedTopicParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnSequencedTopicParams contains all the parameters to send to the API endpoint
for the create msg vpn sequenced topic operation typically these are written to a http.Request
*/
type CreateMsgVpnSequencedTopicParams struct {

	/*Body
	  The Sequenced Topic object's attributes.

	*/
	Body *models.MsgVpnSequencedTopic
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithTimeout(timeout time.Duration) *CreateMsgVpnSequencedTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithContext(ctx context.Context) *CreateMsgVpnSequencedTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithHTTPClient(client *http.Client) *CreateMsgVpnSequencedTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithBody(body *models.MsgVpnSequencedTopic) *CreateMsgVpnSequencedTopicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetBody(body *models.MsgVpnSequencedTopic) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnSequencedTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) WithSelect(selectVar []string) *CreateMsgVpnSequencedTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn sequenced topic params
func (o *CreateMsgVpnSequencedTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnSequencedTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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