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

// NewGetMsgVpnClientProfilesParams creates a new GetMsgVpnClientProfilesParams object
// with the default values initialized.
func NewGetMsgVpnClientProfilesParams() *GetMsgVpnClientProfilesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnClientProfilesParams{
		Count: &countDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnClientProfilesParamsWithTimeout creates a new GetMsgVpnClientProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnClientProfilesParamsWithTimeout(timeout time.Duration) *GetMsgVpnClientProfilesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnClientProfilesParams{
		Count: &countDefault,

		timeout: timeout,
	}
}

// NewGetMsgVpnClientProfilesParamsWithContext creates a new GetMsgVpnClientProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnClientProfilesParamsWithContext(ctx context.Context) *GetMsgVpnClientProfilesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnClientProfilesParams{
		Count: &countDefault,

		Context: ctx,
	}
}

// NewGetMsgVpnClientProfilesParamsWithHTTPClient creates a new GetMsgVpnClientProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnClientProfilesParamsWithHTTPClient(client *http.Client) *GetMsgVpnClientProfilesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnClientProfilesParams{
		Count:      &countDefault,
		HTTPClient: client,
	}
}

/*GetMsgVpnClientProfilesParams contains all the parameters to send to the API endpoint
for the get msg vpn client profiles operation typically these are written to a http.Request
*/
type GetMsgVpnClientProfilesParams struct {

	/*Count
	  Limit the count of objects in the response. See [Count](#count "Description of the syntax of the `count` parameter").

	*/
	Count *int64
	/*Cursor
	  The cursor, or position, for the next page of objects. See [Cursor](#cursor "Description of the syntax of the `cursor` parameter").

	*/
	Cursor *string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string
	/*Where
	  Include in the response only objects where certain conditions are true. See [Where](#where "Description of the syntax of the `where` parameter").

	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithTimeout(timeout time.Duration) *GetMsgVpnClientProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithContext(ctx context.Context) *GetMsgVpnClientProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithHTTPClient(client *http.Client) *GetMsgVpnClientProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithCount(count *int64) *GetMsgVpnClientProfilesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithCursor(cursor *string) *GetMsgVpnClientProfilesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnClientProfilesParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithSelect(selectVar []string) *GetMsgVpnClientProfilesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) WithWhere(where []string) *GetMsgVpnClientProfilesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn client profiles params
func (o *GetMsgVpnClientProfilesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnClientProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
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

	valuesWhere := o.Where

	joinedWhere := swag.JoinByFormat(valuesWhere, "csv")
	// query array param where
	if err := r.SetQueryParam("where", joinedWhere...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}