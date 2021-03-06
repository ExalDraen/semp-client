// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsernameMsgVpnAccessLevelException username msg vpn access level exception
// swagger:model UsernameMsgVpnAccessLevelException
type UsernameMsgVpnAccessLevelException struct {

	// vpn-scope access-level to assign to CLI users. The default value is `"none"`. The allowed values and their meaning are:
	//
	// <pre>
	// "none" - User has no access to a Message VPN.
	// "read-only" - User has read-only access to a Message VPN.
	// "read-write" - User has read-write access to most Message VPN settings.
	// </pre>
	//
	// Enum: [none read-only read-write]
	AccessLevel string `json:"accessLevel,omitempty"`

	// The name of the Message VPN for which an access level exception may be configured.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// Username.
	UserName string `json:"userName,omitempty"`
}

// Validate validates this username msg vpn access level exception
func (m *UsernameMsgVpnAccessLevelException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var usernameMsgVpnAccessLevelExceptionTypeAccessLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","read-only","read-write"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usernameMsgVpnAccessLevelExceptionTypeAccessLevelPropEnum = append(usernameMsgVpnAccessLevelExceptionTypeAccessLevelPropEnum, v)
	}
}

const (

	// UsernameMsgVpnAccessLevelExceptionAccessLevelNone captures enum value "none"
	UsernameMsgVpnAccessLevelExceptionAccessLevelNone string = "none"

	// UsernameMsgVpnAccessLevelExceptionAccessLevelReadOnly captures enum value "read-only"
	UsernameMsgVpnAccessLevelExceptionAccessLevelReadOnly string = "read-only"

	// UsernameMsgVpnAccessLevelExceptionAccessLevelReadWrite captures enum value "read-write"
	UsernameMsgVpnAccessLevelExceptionAccessLevelReadWrite string = "read-write"
)

// prop value enum
func (m *UsernameMsgVpnAccessLevelException) validateAccessLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, usernameMsgVpnAccessLevelExceptionTypeAccessLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UsernameMsgVpnAccessLevelException) validateAccessLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessLevelEnum("accessLevel", "body", m.AccessLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsernameMsgVpnAccessLevelException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsernameMsgVpnAccessLevelException) UnmarshalBinary(b []byte) error {
	var res UsernameMsgVpnAccessLevelException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
