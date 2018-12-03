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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams creates a new DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams object
// with the default values initialized.
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams() *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithTimeout creates a new DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithContext creates a new DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithContext(ctx context.Context) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithHTTPClient creates a new DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
for the delete msg vpn rest delivery point rest consumer Tls trusted common name operation typically these are written to a http.Request
*/
type DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*RestConsumerName
	  The restConsumerName of the REST Consumer.

	*/
	RestConsumerName string
	/*RestDeliveryPointName
	  The restDeliveryPointName of the REST Delivery Point.

	*/
	RestDeliveryPointName string
	/*TLSTrustedCommonName
	  The tlsTrustedCommonName of the Trusted Common Name.

	*/
	TLSTrustedCommonName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithContext(ctx context.Context) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRestConsumerName adds the restConsumerName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithRestConsumerName(restConsumerName string) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetRestConsumerName(restConsumerName)
	return o
}

// SetRestConsumerName adds the restConsumerName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetRestConsumerName(restConsumerName string) {
	o.RestConsumerName = restConsumerName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithRestDeliveryPointName(restDeliveryPointName string) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithTLSTrustedCommonName adds the tLSTrustedCommonName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WithTLSTrustedCommonName(tLSTrustedCommonName string) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams {
	o.SetTLSTrustedCommonName(tLSTrustedCommonName)
	return o
}

// SetTLSTrustedCommonName adds the tlsTrustedCommonName to the delete msg vpn rest delivery point rest consumer Tls trusted common name params
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) SetTLSTrustedCommonName(tLSTrustedCommonName string) {
	o.TLSTrustedCommonName = tLSTrustedCommonName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param restConsumerName
	if err := r.SetPathParam("restConsumerName", o.RestConsumerName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
		return err
	}

	// path param tlsTrustedCommonName
	if err := r.SetPathParam("tlsTrustedCommonName", o.TLSTrustedCommonName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}