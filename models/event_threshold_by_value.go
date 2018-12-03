// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EventThresholdByValue event threshold by value
// swagger:model EventThresholdByValue
type EventThresholdByValue struct {

	// Threshold clear value.
	ClearValue int64 `json:"clearValue,omitempty"`

	// Threshold set value.
	SetValue int64 `json:"setValue,omitempty"`
}

// Validate validates this event threshold by value
func (m *EventThresholdByValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventThresholdByValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventThresholdByValue) UnmarshalBinary(b []byte) error {
	var res EventThresholdByValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}