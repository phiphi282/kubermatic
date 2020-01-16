// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UpdateWindow update window
// swagger:model UpdateWindow
type UpdateWindow struct {

	// length
	Length string `json:"length,omitempty"`

	// start
	Start string `json:"start,omitempty"`
}

// Validate validates this update window
func (m *UpdateWindow) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateWindow) UnmarshalBinary(b []byte) error {
	var res UpdateWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}