// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Limits Limits is a struct that contains the response of a limit query.
//
// swagger:model Limits
type Limits struct {

	// absolute
	Absolute *Absolute `json:"absolute,omitempty"`
}

// Validate validates this limits
func (m *Limits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsolute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Limits) validateAbsolute(formats strfmt.Registry) error {

	if swag.IsZero(m.Absolute) { // not required
		return nil
	}

	if m.Absolute != nil {
		if err := m.Absolute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("absolute")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Limits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Limits) UnmarshalBinary(b []byte) error {
	var res Limits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
